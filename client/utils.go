package main

import (
	"bufio"
	crypto_rand "crypto/rand"
	"fmt"
	"io"
	"os"
)

func genNonce() *[24]byte {
	var nonce [24]byte
	if _, err := io.ReadFull(crypto_rand.Reader, nonce[:]); err != nil {
		panic(err)
	}

	return &nonce
}

func printHelp() {
	fmt.Println(`
		"--port": the port of the server, defaults to :3000
		"--address": the address of the server, defaults to localhost
		"--stream": "the source of numbers stream, defaults to stdin
		"--config": "path of the config file that contains the public key and private key
		`)
}

func doesFileExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func openFileForReader(path string) *bufio.Scanner {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("out file does not exist")
		panic(err)
	}

	reader := bufio.NewScanner(f)

	return reader
}

func openFileForWriter(path string) *os.File {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("out file does not exist")
		panic(err)
	}

	f.Sync()

	return f
}

func createThreeChannels() (*chan struct{}, *chan struct{}, *chan struct{}) {
	ch1 := make(chan struct{}, 1)
	ch2 := make(chan struct{}, 1)
	ch3 := make(chan struct{}, 1)

	return &ch1, &ch2, &ch3
}
