package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mediocregopher/fiestaturtle/fiestatypes"
	"github.com/mediocregopher/radix.v2/pool"
)

var p, _ = pool.New("tcp", "127.0.0.1:6379", 10)

const autocompleteKey = "autocomplete"

func itemKeyRaw(typ string, id fiestatypes.BlockID) string {
	return fmt.Sprintf("item:%s:%s", typ, id)
}

func keyType(key string) string {
	p := strings.SplitN(key, ":", 3)
	if len(p) != 3 {
		return ""
	}
	return p[1]
}

type Items struct {
	Songs     []fiestatypes.Song     `json:"song"`
	Playlists []fiestatypes.Playlist `json:"playlist"`
	Users     []fiestatypes.User     `json:"user"`
}

func getItems(keys []string) (Items, error) {
	bb, err := p.Cmd("MGET", keys).ListBytes()
	if err != nil {
		return Items{}, err
	}

	var it Items
	for i := range bb {
		if bb[i] == nil {
			continue
		}

		switch keyType(keys[i]) {
		case "song":
			var s fiestatypes.Song
			if err := json.Unmarshal(bb[i], &s); err != nil {
				return it, err
			}
			it.Songs = append(it.Songs, s)

		case "playlist":
			var p fiestatypes.Playlist
			if err := json.Unmarshal(bb[i], &p); err != nil {
				return it, err
			}
			it.Playlists = append(it.Playlists, p)

		case "user":
			var u fiestatypes.User
			if err := json.Unmarshal(bb[i], &u); err != nil {
				return it, err
			}
			it.Users = append(it.Users, u)
		}
	}

	return it, nil
}

func itemType(i interface{}) string {
	var typ string
	switch i.(type) {
	case fiestatypes.Song, *fiestatypes.Song:
		typ = "song"
	case fiestatypes.Playlist, *fiestatypes.Playlist:
		typ = "playlist"
	case fiestatypes.User, *fiestatypes.User:
		typ = "user"
	}
	return typ
}

func itemKey(i interface{}) string {
	up := i.(fiestatypes.Uploader).Get()
	if up.ID == "" {
		panic("invalid ID")
	}

	return itemKeyRaw(itemType(i), up.ID)
}

func setItem(i interface{}) error {
	b, err := json.Marshal(i)
	if err != nil {
		return err
	}

	return p.Cmd("SET", itemKey(i), b).Err
}
