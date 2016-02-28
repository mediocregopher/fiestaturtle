package richard

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"unicode"

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
	typ := itemType(i)
	if typ == "" {
		return ""
	}

	return itemKeyRaw(typ, up.ID)
}

func setItem(i interface{}) error {
	k := itemKey(i)
	if k == "" {
		return nil
	}

	b, err := json.Marshal(i)
	if err != nil {
		return err
	}

	return p.Cmd("SET", k, b).Err
}

func tokenize(s string) []string {
	ret := make([]string, 0, 1)
	next := ""
	for _, c := range s {
		if unicode.IsLetter(c) || unicode.IsNumber(c) {
			c = unicode.ToLower(c)
			next += string(c)
		} else if next == "" {
			continue
		} else {
			ret = append(ret, next)
			next = ""
		}
	}
	if next != "" {
		ret = append(ret, next)
	}
	return ret
}

func indexItem(i interface{}) error {
	k := itemKey(i)
	if k == "" {
		return nil
	}

	var tokens []string

	switch ii := i.(type) {
	case fiestatypes.Song:
		sm := ii.SongMeta
		tokens = append(tokens, tokenize(sm.Title)...)
		for _, a := range sm.Artists {
			tokens = append(tokens, tokenize(a)...)
		}
		tokens = append(tokens, tokenize(sm.Album)...)

	case fiestatypes.Playlist:
		tokens = append(tokens, tokenize(ii.Name)...)

	case fiestatypes.User:
		tokens = append(tokens, tokenize(ii.Name)...)
	}
	if len(tokens) == 0 {
		return nil
	}

	args := make([]interface{}, 0, len(tokens)*2)
	for _, tok := range tokens {
		args = append(args, 0, fmt.Sprintf("%s|%s", tok, k))
	}
	if len(args) == 0 {
		return nil
	}

	return p.Cmd("ZADD", autocompleteKey, args).Err
}

type sorterEl struct {
	s string
	i int
}

type sorter []sorterEl

func (s sorter) Len() int {
	return len(s)
}

func (s sorter) Less(i, j int) bool {
	return s[i].i < s[j].i
}

func (s sorter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func searchItems(query string) (Items, error) {
	m := map[string]int{}
	toks := tokenize(query)
	for _, tok := range toks {
		vals, err := p.Cmd("ZRANGEBYLEX", autocompleteKey, "["+tok, "["+tok+"\xff").List()
		if err != nil {
			return Items{}, err
		}

		for _, val := range vals {
			p := strings.SplitN(val, "|", 2)
			if len(p) != 2 {
				continue
			}
			m[p[1]]++
		}
	}

	s := make(sorter, 0, len(m))
	for k, score := range m {
		s = append(s, sorterEl{s: k, i: score})
	}
	sort.Sort(sort.Reverse(s))

	keys := make([]string, len(s))
	for i := range s {
		keys[i] = s[i].s
	}

	return getItems(keys)
}
