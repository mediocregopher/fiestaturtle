package main

import (
	"github.com/levenlabs/golib/rpcutil"
	"github.com/mediocregopher/fiestaturtle/fiestatypes"
)

var noRichardOps = false

func forEachRichard(fn func(r Richard) error) error {
	if noRichardOps {
		return nil
	}

	u, err := getNSUser()
	if err != nil {
		return err
	}
	for _, r := range u.UserPrivate.Richards {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

func addRichardSong(s fiestatypes.Song) error {
	args := richard.AddSongArgs{
		Song: s,
	}
	return forEachRichard(func(r Richard) error {
		return rpcutil.JSONRPC2Call("http://"+r.Addr+"/rpc", &struct{}{}, "Richard.AddSong", &args)
	})
}

func addRichardPlaylist(s fiestatypes.Playlist) error {
	args := richard.AddPlaylistArgs{
		Playlist: s,
	}
	return forEachRichard(func(r Richard) error {
		return rpcutil.JSONRPC2Call("http://"+r.Addr+"/rpc", &struct{}{}, "Richard.AddPlaylist", &args)
	})
}
