package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type Reader interface {
	// ReadString(delim byte) (string, error)
	Scan() bool
	Err() error
	Text() string
}

// reader from file or stdin to send to client and then the client sends that to server
// reader means where to read the data from
func reader(r Reader, input *chan interface{}, quit *chan struct{}) {
	for r.Scan() {
		text := r.Text()

		// if len(text) == 1 {
		// 	continue
		// }

		text = strings.TrimSpace(text)

		if text == "c" {
			(*input) <- "c"
		} else if text == "r" {
			(*input) <- "r"
		} else {
			i, err := strconv.Atoi(text)
			if err != nil {
				color.Red("Enter an Integer.")
			} else {
				(*input) <- i
			}
		}

	}

	if r.Err() != nil {
		panic(r.Err())
	}

	<-(*quit)
	fmt.Println("Reader closed")
}
