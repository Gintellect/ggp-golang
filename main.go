package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", messageHandler)
	port := "9147"
	log.Printf("Launching player on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func messageHandler(rw http.ResponseWriter, req *http.Request) {
	msg, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Printf("%s\n", err)
	}

	log.Printf("Received: %s\n", msg)
}
