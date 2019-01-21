package main

import (
	"bufio"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func parseConfigFile(path string) (publicKey *[32]byte, privateKey *[32]byte) {
	publicKey = &[32]byte{}
	privateKey = &[32]byte{}

	var config struct {
		PublicKey  string `json:"publicKey"`
		PrivateKey string `json:"privateKey"`
	}

	buff, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Config file %s is not found\n", path)
		panic(err)
	}

	err = json.Unmarshal(buff, &config)
	if err != nil {
		panic(err)
	}

	pub, _ := hex.DecodeString(config.PublicKey)
	priv, _ := hex.DecodeString(config.PrivateKey)

	copy(publicKey[:], pub)
	copy(privateKey[:], priv)

	return
}

func saveKeys(publicKey *[32]byte, privateKey *[32]byte) error {
	if doesFileExist("config.json") {
		fmt.Println("Are you sure you want to override the existing keys: (Y/N)")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		// trim '\n'
		text = text[:len(text)-1]

		if text == "N" || text == "n" {
			return errors.New("config.json alreay exists")
		}
	}

	var config struct {
		PublicKey  string `json:"publicKey"`
		PrivateKey string `json:"privateKey"`
	}

	config.PublicKey = hex.EncodeToString(publicKey[:])
	config.PrivateKey = hex.EncodeToString(privateKey[:])

	buff, err := json.Marshal(config)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("config.json", buff, 0644)
	if err != nil {
		panic(err)
	}

	return nil
}
