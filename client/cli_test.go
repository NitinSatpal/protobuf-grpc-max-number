package main

import (
	"bufio"
	"bytes"
	crypto_rand "crypto/rand"
	"strings"
	"testing"
	"time"

	"golang.org/x/crypto/nacl/box"
)

func setupClient() *Client {
	stream := connectToServer("localhost", ":3000")

	publicKey, privateKey, _ := box.GenerateKey(crypto_rand.Reader)

	client := NewClient(stream, publicKey, privateKey)

	return client
}

func TestNonceIsRandom(t *testing.T) {
	n1 := genNonce()
	n2 := genNonce()

	for i, v := range n1 {
		if n2[i] != v {
			return
		}
	}

	t.Errorf("generating Nonce is random")
}

func TestIsCorrectMaxNumbers(t *testing.T) {
	expectedServerResponse := "1\n5\n6\n20\n"

	client := setupClient()

	var buffWriter bytes.Buffer
	writer := bufio.NewWriter(&buffWriter)

	buffReader := strings.NewReader("1\n5\n3\n6\n2\n20\n")
	reader := bufio.NewScanner(buffReader)

	quitCh, fun := run(client, reader, writer)

	go func() {
		// wait until the server responds
		time.Sleep(1 * time.Second)
		quitCh <- struct{}{}
	}()

	fun()

	writer.Flush()

	buffWriterString := buffWriter.String()
	if buffWriterString != expectedServerResponse {
		t.Fatal("Response from the server does not equal the expected result.")
	}
}
