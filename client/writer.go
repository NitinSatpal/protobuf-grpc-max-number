// this file is concerned with writing the responses of the server
// responses of the server may be written either to `stdout` or a file
package main

import (
	"fmt"
	"io"
)

// writer write to a file or stdout, gets the data from the ch which the client writes into
// writer means where to write the data
func writer(w io.Writer, ch *chan int, quit *chan struct{}) {
loop:
	for {
		select {
		case i := <-(*ch):
			fmt.Fprintln(w, i)
		case <-(*quit):
			break loop
		}

	}

	fmt.Println("Writer closed")
}
