/******************************************************************************
 *
 *  Description :
 *
 *  Setup & initialization.
 *
 *****************************************************************************/

package main

import (
	"flag"
	"log"

	// Database backends
	_ "max/server/db/boltdb"

	"max/server/store"

	"google.golang.org/grpc"
)

var globals struct {
	sessionStore *SessionStore
	grpcServer   *grpc.Server
}

var portPtr = flag.String("port", ":3000", "the port the server will use, defaults to `:3000`")
var dbPath = flag.String("dbPath", "my.db", "the path of the DB")

func main() {
	flag.Parse()

	err := store.Open(*dbPath)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
		return
	}

	defer func() {
		store.Close()
		log.Println("Closed database connection(s)")
		log.Println("All done, good bye")
	}()

	globals.sessionStore = NewSessionStore()

	if globals.grpcServer, err = serveGrpc(*portPtr); err != nil {
		log.Fatal(err)
	}

	select {}
}
