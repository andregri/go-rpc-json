package main

import (
	jsonparse "encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

// Arguments passed by the rpc client
type Args struct {
	Id string // Book id
}

// Struct for reply passed by rpc server
type Book struct {
	Id     string `"json:string,omitempty"`
	Name   string `"json:string,omitempty"`
	Author string `"json:authot,omitempty"`
}

// struct for the service to register to the rpc server
type JSONServer struct{}

// remote function to give book details
func (t *JSONServer) GiveBookDetail(r *http.Request, args *Args, reply *Book) error {
	var books []Book

	// Read JSON file and load data
	raw, readerr := ioutil.ReadFile("./books.json")
	if readerr != nil {
		log.Fatal("read error:", readerr)
	}

	// Unmarshal raw data into books array
	marsherr := jsonparse.Unmarshal(raw, &books)
	if marsherr != nil {
		log.Fatal("unmarshal error:", marsherr)
	}

	// Iterate over each book to find the given book
	for _, book := range books {
		if book.Id == args.Id {
			*reply = book
			break
		}
	}

	return nil
}

func main() {
	// Create a new rpc server
	s := rpc.NewServer()

	// Register the type of data requested as JSON
	s.RegisterCodec(json.NewCodec(), "application/json")

	// Create a new json server
	jsonserver := new(JSONServer)

	// Register the json server to the rpc server
	s.RegisterService(jsonserver, "")

	// Create a mux router for HTTP routes
	r := mux.NewRouter()

	// Handle /rpc route with the rpc server
	r.Handle("/rpc", s)

	log.Fatal(http.ListenAndServe(":1234", r))
}
