# Prerequisites
go1.11.4

# How it works?
The project is about creating client server architecture using protocol buffers with grpc.
The system allows client to get the maximum numbers from the stream of numbers it can send to the server. The client sends hello to server with public key. The server responds with its own public key as the confirmation of connection establishment.
Thereafter, the can exhange the stream of numbers with client send encrypted numbers with it's private key and server returning the maximum number till now for that client.


# How to run?

## server

go run *.go

## client

Run `go run client.go config.go main.go reader.go utils.go writer.go --help` to show help

There are two mode for running the client
- the first is to read numbers from stdin, to run the client in this mode run: 
    `go run client.go config.go main.go reader.go utils.go writer.go`
- the other mode is to take numbers from a file, to run the client in this mode run:
    `go run client.go config.go main.go reader.go utils.go writer.go --input <path of the file>`

the first time you run the client keys will be created automatically, and will be saved in a file called `config.json`.

to run the client with a specific `config.json` run:
    `go run client.go config.go main.go reader.go utils.go writer.go --config <path of the config.json>`

if you did not specify a config file and a config.json already exists in the current directory, you will get a warning that if you hit yes you will orverride the current config, you need to run `go run client.go config.go main.go reader.go utils.go writer.go --config <path of the config.json>` so the config is not overriddin.


if you want the responses of the server to be written to a file, run:
    `go run client.go config.go main.go reader.go utils.go writer.go --config config.json --output <name of the file>`
