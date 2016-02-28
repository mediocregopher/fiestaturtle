package main

import (
	"sync"

	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/repo/fsrepo"
	"golang.org/x/net/context"
)

type node struct {
	sync.Mutex
	nd   *core.IpfsNode
	rdir string
}

func newNode(live bool) (node, error) {
	log.Infof("finding repo")
	dir, err := fsrepo.BestKnownPath()
	if err != nil {
		return node{}, err
	}

	cfg := &core.BuildCfg{
		Online: live,
	}

	if live {
		log.Infof("opening repo at %q", dir)
		r, err := fsrepo.Open(dir)
		if err != nil {
			return node{}, err
		}
		cfg.Repo = r
	}

	log.Infof("initializing node object")
	nd, err := core.NewNode(context.Background(), cfg)
	if err != nil {
		return node{}, err
	}
	log.Infof("done initializing")

	if !live {
		log.Infof("setting up offline routing")
		if err := nd.SetupOfflineRouting(); err != nil {
			return node{}, err
		}
		log.Infof("done setting up offline routing")
	}

	n := node{
		nd:   nd,
		rdir: dir,
	}
	return n, nil
}

func (n node) with(fn func()) {
	n.Lock()
	fn()
	n.Unlock()
}
