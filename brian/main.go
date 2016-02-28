package main

import (
	logging "gx/ipfs/QmQvJiADDe7JR4m968MwXobTCCzUqQkP87aRHe29MEBGHV/go-logging"
	logger "gx/ipfs/Qmazh5oNUVsDZTs2g59rq8aYQqwpss8tcUWQzor5sCCEuH/go-log"
	"net/http"
)

var n node
var log = logger.Logger("brian")

func init() {
	logger.SetAllLoggers(logging.INFO)
}

func main() {
	log.Infof("starting")
	var err error
	n, err = newNode(true)
	if err != nil {
		log.Fatalf("could not setup node: %s", err)
	}

	http.Handle("/rpc", RPC())
	log.Info("listening on :4567")
	log.Fatal(http.ListenAndServe("127.0.0.1:4567", nil))
}
