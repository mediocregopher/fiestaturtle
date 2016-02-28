package main

import (
	"encoding/json"
	logging "gx/ipfs/Qmazh5oNUVsDZTs2g59rq8aYQqwpss8tcUWQzor5sCCEuH/go-log"
)

var n node
var log = logging.Logger("brian")

func main() {
	log.Infof("starting")
	var err error
	n, err = newNode(false)
	if err != nil {
		log.Fatalf("could not setup node: %s", err)
	}

	log.Infof("id: %q", n.nd.Identity.Pretty())

	log.Infof("putNS")
	msg := json.RawMessage(`{"foo":"bar"}`)
	if _, err := putNS(&msg); err != nil {
		log.Fatalf("err putNS: %s", err)
	}
	log.Infof("done with putNS")

	log.Infof("getNS")
	var res json.RawMessage
	id, err := getNS(&res)
	if err != nil {
		log.Fatalf("getNS: %s", err)
	}
	log.Infof("id:%q res:%s", id, string(res))
}
