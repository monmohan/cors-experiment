package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

type User struct {
	UserName  string
	FirstName string
	LastName  string
	Country   string
}

var userData = map[string]User{
	"john": User{"jdoe", "John", "Doe", "France"},
}
var port = flag.Int("port", 12346, "port to listen on, default is 12346")

func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(userData[r.URL.Path[len("/users/"):]])
	io.WriteString(w, string(b))

}

func main() {
	flag.Parse()
	http.HandleFunc("/users/", userHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("apiserver.cors.com:%d", *port), nil))
}
