package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port = flag.Int("port", 12345, "port to run the server on, default is 12345")

func fileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Requested URL %v\n", r.URL.Path)
	http.ServeFile(w, r, r.URL.Path[1:])
}

func main() {
	flag.Parse()
	http.HandleFunc("/", fileHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("pageserver.cors.com:%d", *port), nil))
}
