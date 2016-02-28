package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"path"
	"time"

	"github.com/ipfs/go-ipfs/core/coreunix"
	uio "github.com/ipfs/go-ipfs/unixfs/io"
	"golang.org/x/net/context"
)

func newCtx() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	return ctx
}

func nsFilePath() string {
	return path.Join(n.rdir, "ns")
}

func get(id BlockID, res interface{}) error {
	var err error
	n.with(func() {
		var r *uio.DagReader
		if r, err = coreunix.Cat(newCtx(), n.nd, string(id)); err != nil {
			return
		}
		defer r.Close()

		if err = json.NewDecoder(r).Decode(res); err != nil {
			return
		}

		if ut, ok := res.(uploader); ok {
			ut.uploaded(id)
		}
	})
	return err
}

func put(res interface{}) (BlockID, error) {
	if ut, ok := res.(uploader); ok {
		if err := ut.sign(); err != nil {
			return "", err
		}
	}

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(res); err != nil {
		return "", err
	}

	id, err := putRaw(buf)
	if err != nil {
		return "", err
	}

	if ut, ok := res.(uploader); ok {
		ut.uploaded(id)
	}

	return id, nil
}

func putRaw(r io.Reader) (BlockID, error) {
	var braw string
	var err error
	n.with(func() {
		braw, err = coreunix.Add(n.nd, r)
	})
	return BlockID(braw), err
}

func putNS(res interface{}) (BlockID, error) {
	id, err := put(res)
	if err != nil {
		return "", err
	}

	return id, ioutil.WriteFile(nsFilePath(), []byte(id), 0644)
}

func getNodeID() NodeID {
	var nid NodeID
	n.with(func() {
		nid = NodeID(n.nd.Identity.Pretty())
	})
	return nid
}

func getNS(res interface{}) (BlockID, error) {
	nsb, err := ioutil.ReadFile(nsFilePath())
	if err != nil {
		return "", err
	}

	id := BlockID(nsb)
	err = get(id, res)
	return id, err
}
