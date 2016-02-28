package main

import (
	"log"
	"net/http"

	"github.com/mediocregopher/fiestaturtle/richard/richard"
)

func main() {
	http.Handle("/rpc", richard.RPC())
	log.Printf("listening on :5678")
	log.Fatal(http.ListenAndServe(":5678", nil))
}
