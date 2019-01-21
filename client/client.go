package main

import (
	"errors"
	"fmt"
	"io"
	"max/pb"
	"strconv"
	"sync"

	"github.com/fatih/color"
	"golang.org/x/crypto/nacl/box"
)

type Client struct {
	stream pb.Service_MessageClient
	myID   string
	// writes to stream field
	sendCh   chan *pb.ClientMessage
	quitCh   *chan struct{}
	isClosed bool

	// this where the client receive integer to send to the server
	inputCh *chan interface{}
	// this where the client will send the responses of the server
	outputCh *chan int

	myPrivateKey    *[32]byte
	myPublicKey     *[32]byte
	serverPublicKey *[32]byte

	// to distinguish between the messages sent
	nextMessageID int
	// lock for nextMessageID
	lock sync.Mutex
}

// NewClient create a new Client
func NewClient(stream pb.Service_MessageClient, publicKey *[32]byte, privateKey *[32]byte) *Client {
	return &Client{
		stream:          stream,
		sendCh:          make(chan *pb.ClientMessage, 256),
		myPrivateKey:    privateKey,
		myPublicKey:     publicKey,
		serverPublicKey: &[32]byte{},
		nextMessageID:   0,
		lock:            sync.Mutex{},
	}
}

func (client *Client) SetChannels(in *chan interface{}, out *chan int, quit *chan struct{}) {
	client.inputCh = in
	client.outputCh = out
	client.quitCh = quit
}

// Run sends a Hello message to the server with the client public ID and waits for the server to send back its public ID
func (client *Client) Run() {
	client.Hello()

	// recieve one message that contains the servers public key
	m, err := client.stream.Recv()
	if err != nil {
		panic(err)
	}

	hi := m.GetHi()
	copy(client.serverPublicKey[:], hi.PublicId)

	client.run()
}

func (client *Client) run() {
	go client.read()
	go client.write()

	go client.runInputCh()
	go client.runQuitCh()
}

func (client *Client) runInputCh() {
	if client.inputCh == nil {
		return
	}

	for client.isClosed != true {
		i := <-*client.inputCh

		switch v := i.(type) {
		case int:
			client.Number(v)
		case string:
			if v == "c" {
				client.CurrentMaxNumber()
			} else if v == "r" {
				client.ResetCurrentMaxNumber()
			}
		}
	}

}

func (client *Client) runQuitCh() {
	<-(*client.quitCh)
	client.isClosed = true
	client.stream.CloseSend()

	fmt.Println("Client closed")
}

func (client *Client) read() {
	for client.isClosed != true {
		in, err := client.stream.Recv()
		if err == io.EOF {
			// read done.
			break
		}
		if err != nil {
			break
		}

		i, err := client.DecryptToNumber(in.GetNumber().Number)
		if err != nil {
			continue
		}

		(*client.outputCh) <- i
		// color.Yellow("Current Max Number: %d", i)
	}
}

func (client *Client) write() {
	for client.isClosed != true {
		m := <-client.sendCh
		client.stream.Send(m)
	}
}

// ### The below function are like the api to use a client

// Hello sends Hello message to the server
func (client *Client) Hello() {
	client.stream.Send(&pb.ClientMessage{
		Message: &pb.ClientMessage_Hello{
			Hello: &pb.ClientHello{
				Id:       client.getNextMessageID(),
				PublicId: client.myPublicKey[:],
			},
		},
	})
}

// Number send a number to the server
func (client *Client) Number(number int) {
	client.sendCh <- &pb.ClientMessage{
		Message: &pb.ClientMessage_Number{
			Number: &pb.ClientNumber{
				Id:     client.getNextMessageID(),
				Number: client.EncryptFromNumber(number),
			},
		},
	}
}

// CurrentMaxNumber gets the current max number from the server
func (client *Client) CurrentMaxNumber() {
	client.sendCh <- &pb.ClientMessage{
		Message: &pb.ClientMessage_CurrentMaxNumber{
			CurrentMaxNumber: &pb.ClientCurrentMaxNumber{
				Id: client.getNextMessageID(),
			},
		},
	}
}

func (client *Client) ResetCurrentMaxNumber() {
	color.Green("Resetting current max number...")

	client.sendCh <- &pb.ClientMessage{
		Message: &pb.ClientMessage_ResetCurrentMaxNumber{
			ResetCurrentMaxNumber: &pb.ClientResetCurrentMaxNumber{
				Id: client.getNextMessageID(),
			},
		},
	}
}

// ### functions below are helper functions

// Encrypt takes a slice and encrypts it
func (client *Client) Encrypt(msg []byte) []byte {
	nonce := genNonce()

	encrypted := box.Seal(nonce[:], msg, nonce, client.serverPublicKey, client.myPrivateKey)

	return encrypted
}

// EncryptFromString takes a string
func (client *Client) EncryptFromString(msg string) []byte {
	return client.Encrypt([]byte(msg))
}

// EncryptFromNumber takes an int
func (client *Client) EncryptFromNumber(n int) []byte {
	return client.EncryptFromString(fmt.Sprintf("%d", n))
}

// Decrypt takes a slice of bytes and returns it unencrypted
func (client *Client) Decrypt(encrypted []byte) ([]byte, error) {
	var decryptNonce [24]byte
	copy(decryptNonce[:], encrypted[:24])

	decrypted, ok := box.Open(nil, encrypted[24:], &decryptNonce, client.serverPublicKey, client.myPrivateKey)
	if !ok {
		return nil, errors.New("error decrypting")
	}

	return decrypted, nil
}

// DecryptToNumber takes a slice of bytes an decrypts it into an int
func (client *Client) DecryptToNumber(encrypted []byte) (int, error) {
	d, err := client.Decrypt(encrypted)
	if err != nil {
		return 0, err
	}

	i, err := strconv.Atoi(string(d))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func (client *Client) getNextMessageID() int64 {
	client.lock.Lock()
	defer client.lock.Unlock()

	client.nextMessageID++

	return int64(client.nextMessageID)
}
