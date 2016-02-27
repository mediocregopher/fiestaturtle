package main

import "log"

var n node

func main() {
	log.Printf("starting")
	var err error
	n, err = newNode("~/.ipfs")
	if err != nil {
		log.Fatalf("could not setup node: %s", err)
	}
}
