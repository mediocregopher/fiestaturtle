package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/ipfs/go-ipfs/core/coreunix"
	uio "github.com/ipfs/go-ipfs/unixfs/io"
	"github.com/mediocregopher/fiestaturtle/fiestatypes"
	"golang.org/x/net/context"
)

func newCtx() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	return ctx
}

func nsFilePath() string {
	return path.Join(n.rdir, "ns")
}

func cat(id fiestatypes.BlockID) (io.ReadCloser, error) {
	var err error
	var r *uio.DagReader
	n.with(func() {
		if r, err = coreunix.Cat(newCtx(), n.nd, string(id)); err != nil {
			return
		}
	})
	return r, err
}

func get(id fiestatypes.BlockID, res interface{}) error {
	rc, err := cat(id)
	if err != nil {
		return err
	}
	defer rc.Close()

	if err = json.NewDecoder(rc).Decode(res); err != nil {
		return err
	}

	if ut, ok := res.(fiestatypes.Uploader); ok {
		ut.SetID(id)
	}
	return nil
}

func put(res interface{}) (fiestatypes.BlockID, error) {
	if ut, ok := res.(fiestatypes.Uploader); ok {
		if err := signUploader(ut); err != nil {
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

	if ut, ok := res.(fiestatypes.Uploader); ok {
		ut.SetID(id)
	}

	return id, nil
}

func putRaw(r io.Reader) (fiestatypes.BlockID, error) {
	var braw string
	var err error
	n.with(func() {
		braw, err = coreunix.Add(n.nd, r)
	})
	return fiestatypes.BlockID(braw), err
}

func putNS(res interface{}) (fiestatypes.BlockID, error) {
	id, err := put(res)
	if err != nil {
		return "", err
	}

	return id, ioutil.WriteFile(nsFilePath(), []byte(id), 0644)
}

func getNodeID() fiestatypes.NodeID {
	var nid fiestatypes.NodeID
	n.with(func() {
		nid = fiestatypes.NodeID(n.nd.Identity.Pretty())
	})
	return nid
}

func getNS(res interface{}) (fiestatypes.BlockID, error) {
	nsb, err := ioutil.ReadFile(nsFilePath())
	if err != nil {
		return "", err
	}

	id := fiestatypes.BlockID(nsb)
	err = get(id, res)
	return id, err
}

func getNSUser() (fiestatypes.User, error) {
	var u fiestatypes.User
	u.Uploaded = &fiestatypes.Uploaded{}
	_, err := getNS(&u)
	if err != nil && !os.IsNotExist(err) {
		return u, err
	}
	if err := decryptUser(&u); err != nil {
		return u, err
	}
	return u, nil
}

func putNSUser(u fiestatypes.User) (fiestatypes.User, error) {
	if err := encryptUser(&u); err != nil {
		return u, err
	}

	// Make super sure that we never store the user with UserPrivate set
	u.UserPrivate = fiestatypes.UserPrivate{}

	_, err := putNS(&u)
	return u, err
}
