package main

import (
	logging "gx/ipfs/QmQvJiADDe7JR4m968MwXobTCCzUqQkP87aRHe29MEBGHV/go-logging"
	logger "gx/ipfs/Qmazh5oNUVsDZTs2g59rq8aYQqwpss8tcUWQzor5sCCEuH/go-log"
	"io"
	"net/http"
	"path"

	"github.com/mediocregopher/fiestaturtle/fiestatypes"
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
	http.HandleFunc("/ipfs/", catFile)
	log.Info("listening on :4567")
	log.Fatal(http.ListenAndServe("127.0.0.1:4567", nil))
}

func catFile(w http.ResponseWriter, r *http.Request) {
	idr := path.Base(r.URL.Path)
	rc, err := cat(fiestatypes.BlockID(idr))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rc.Close()

	io.Copy(w, rc)
}
