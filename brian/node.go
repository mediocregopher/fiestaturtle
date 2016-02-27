package main

import (
	"log"

	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/repo/fsrepo"
	"golang.org/x/net/context"
)

type node struct {
	ch chan *core.IpfsNode
}

func newNode(dir string) (node, error) {
	r, err := fsrepo.Open(dir)
	if err != nil {
		return node{}, err
	}

	cfg := &core.BuildCfg{
		Repo: r,
		//Online: true,
	}

	log.Printf("initializing node object")
	nd, err := core.NewNode(context.Background(), cfg)
	if err != nil {
		return node{}, err
	}
	log.Printf("done initializing")

	n := node{
		ch: make(chan *core.IpfsNode, 1),
	}
	n.ch <- nd
	return n, nil
}

func (n node) with(fn func(*core.IpfsNode)) {
	nd := <-n.ch
	fn(nd)
	n.ch <- nd
}
