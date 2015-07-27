package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	log.Println("start")

	port := flag.String("p", ":3000", "address the server listen on")
	war := flag.String("war", "./public", "directory of war files")
	flag.Parse()

	mount(*war)

	log.Fatal(http.ListenAndServe(*port, nil))
}
