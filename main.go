package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

func main() {

	router := NewRouter()

	var port int
	flag.IntVar(&port, "p", 8080, "specify port to use.  defaults to 8080.")
	flag.Parse()

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
