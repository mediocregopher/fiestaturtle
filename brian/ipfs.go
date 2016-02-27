package main

import (
	"bytes"
	"encoding/json"
	"io"
	"time"

	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreunix"
	uio "github.com/ipfs/go-ipfs/unixfs/io"
	"golang.org/x/net/context"
)

func get(id BlockID, res interface{}) error {
	var err error
	n.with(func(nd *core.IpfsNode) {
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		var r *uio.DagReader
		if r, err = coreunix.Cat(ctx, nd, string(id)); err != nil {
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
	n.with(func(nd *core.IpfsNode) {
		braw, err = coreunix.Add(nd, r)
	})
	return BlockID(braw), err
}

func getNodeID() NodeID {
	var nid NodeID
	n.with(func(nd *core.IpfsNode) {
		nid = NodeID(nd.Identity.Pretty())
	})
	return nid
}
