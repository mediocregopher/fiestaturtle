package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/rpc", RPC())
	log.Printf("listening on :5678")
	log.Fatal(http.ListenAndServe(":5678", nil))
}
