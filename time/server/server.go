package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

// struct for the input arguments passed from the client
type Args struct{}

// number to register the rpc method
type TimeServer int64

func (t *TimeServer) GiveServerTime(args *Args, reply *int64) error {
	// Fill the reply pointer to send data back
	*reply = time.Now().Unix()
	return nil
}

func main() {
	// Create a new RPC server
	timeserver := new(TimeServer)

	// Register the RPC server
	rpc.Register(timeserver)
	rpc.HandleHTTP()

	// Listen for requests on port 1234
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listen error", err)
	}

	http.Serve(listener, nil)
}
