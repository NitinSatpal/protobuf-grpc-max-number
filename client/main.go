package main

import (
	"bufio"
	"context"
	crypto_rand "crypto/rand"
	"fmt"
	"io"
	"max/pb"
	"os"
	"os/signal"
	"syscall"
	"time"

	"flag"

	"github.com/fatih/color"
	"golang.org/x/crypto/nacl/box"

	"google.golang.org/grpc"
)

var helpPtr = flag.Bool("help", false, "")
var portPtr = flag.String("port", ":3000", "the port of the server, defaults to `:3000`")
var serverAddressPtr = flag.String("address", "localhost", "the address of the server, defaults to `localhost`")
var inputPtr = flag.String("input", "stdin", "the source of numbers stream, defaults to `stdin`")
var outputPtr = flag.String("output", "stdout", "where the responds of the server will be written")
var configPtr = flag.String("config", "", "path of the config file that contains the public key and private key")

func main() {
	flag.Parse()

	if *helpPtr {
		printHelp()
		return
	}

	stream := connectToServer(*serverAddressPtr, *portPtr)

	var client *Client

	// whether to read the keys from the file you supplied in configPtr or generate new keys
	if *configPtr != "" {
		pub, priv := parseConfigFile(*configPtr)
		client = NewClient(stream, pub, priv)
	} else {
		publicKey, privateKey, err := box.GenerateKey(crypto_rand.Reader)
		if err != nil {
			panic("Error generating Keys")
		}

		err = saveKeys(publicKey, privateKey)
		if err != nil {
			return
		}

		client = NewClient(stream, publicKey, privateKey)
	}

	// setup writer
	var w io.Writer

	if *outputPtr == "stdout" {
		w = os.Stdout
	} else {
		w = openFileForWriter(*outputPtr)
	}

	// setup reader
	var r Reader
	if *inputPtr == "stdin" {
		r = bufio.NewScanner(os.Stdin)
		color.HiBlue("Enter `c` to print current max number, Enter `r` to reset current max number")
		color.Magenta("Enter Integers...")
	} else {
		r = openFileForReader(*inputPtr)
	}

	// fun is where `quitCh` is read, and it's not a Goroutine, so it hangs until `quitCh`
	// is written into, which happens in `go func(){...}` below
	quitCh, fun := run(client, r, w)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		// this is for so when you hit Ctrl+c, `^C` gets printed on a line
		fmt.Println()
		quitCh <- struct{}{}
	}()

	fun()

	// wait a little so goroutines exit
	time.Sleep(100 * time.Millisecond)
	fmt.Println("exited.")
}

// func run(client *Client, r Reader, w io.Writer) chan struct{} {
func run(client *Client, r Reader, w io.Writer) (chan struct{}, func()) {
	quit := make(chan struct{}, 1)

	clientQuitCh, readerQuitCh, writerQuitCh := createThreeChannels()

	// input means input data, that needs to be sent to server through client
	input := make(chan interface{}, 256)
	// output means output data, which is responses from the server
	output := make(chan int, 256)

	client.SetChannels(&input, &output, clientQuitCh)

	client.Run()
	// reader reads from `r` and feeds that into tha client
	go reader(r, &input, readerQuitCh)
	//writer writes the data sent from client to `w`.
	go writer(w, &output, writerQuitCh)

	// fun basicly reads `quit` and then prodcast to `Client`, `Reader` and `Writer`
	fun := func() {
		<-quit
		(*clientQuitCh) <- struct{}{}
		(*readerQuitCh) <- struct{}{}
		(*writerQuitCh) <- struct{}{}
	}

	return quit, fun

	// return quit
}

func connectToServer(serverAddress string, port string) pb.Service_MessageClient {
	conn, err := grpc.Dial(serverAddress+port, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	conn2 := pb.NewServiceClient(conn)

	stream, err := conn2.Message(context.Background())
	if err != nil {
		panic(err)
	}

	return stream
}
