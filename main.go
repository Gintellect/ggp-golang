package main

import (
	"fmt"
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

func messageHandler(w http.ResponseWriter, r *http.Request) {
	msg, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Printf("%s\n", err)
	}

	log.Printf("Received: %s\n", msg)
	reply := "( (name DeepBeige) (status available) )"
	fmt.Fprintf(w, "%s", reply)
	log.Printf("Replied : %s\n", reply)
}
