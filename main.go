package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var CachedBeacons Beacons = select_all()

func main() {

	fmt.Println("Initializing")

	router := NewRouter()

	var port int
	flag.IntVar(&port, "p", 8080, "specify port to use.  defaults to 8080.")
	flag.Parse()

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
