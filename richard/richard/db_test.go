package richard

import (
	"fmt"
	. "testing"

	"github.com/levenlabs/golib/testutil"
	"github.com/mediocregopher/fiestaturtle/fiestatypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSetGet(t *T) {
	s := fiestatypes.Song{
		Uploaded: &fiestatypes.Uploaded{
			ID: fiestatypes.BlockID(testutil.RandStr()),
		},
	}
	p := fiestatypes.Playlist{
		Uploaded: &fiestatypes.Uploaded{
			ID: fiestatypes.BlockID(testutil.RandStr()),
		},
	}
	u := fiestatypes.User{
		Uploaded: &fiestatypes.Uploaded{
			ID: fiestatypes.BlockID(testutil.RandStr()),
		},
	}

	require.Nil(t, setItem(s))
	require.Nil(t, setItem(p))
	require.Nil(t, setItem(u))

	keys := []string{
		itemKeyRaw("song", s.ID),
		itemKeyRaw("playlist", p.ID),
		itemKeyRaw("user", u.ID),
	}
	i, err := getItems(keys)
	require.Nil(t, err)
	assert.Len(t, i.Songs, 1)
	assert.Equal(t, s, i.Songs[0])
	assert.Len(t, i.Playlists, 1)
	assert.Equal(t, p, i.Playlists[0])
	assert.Len(t, i.Users, 1)
	assert.Equal(t, u, i.Users[0])
}

func TestTokenize(t *T) {
	ss := map[string][]string{
		"":                         []string{},
		"foo":                      []string{"foo"},
		"foo bar":                  []string{"foo", "bar"},
		"Foo Bar":                  []string{"foo", "bar"},
		"foo, bar":                 []string{"foo", "bar"},
		"this, that. other?":       []string{"this", "that", "other"},
		"okie-dokie, hokie-smokie": []string{"okie", "dokie", "hokie", "smokie"},
	}

	for in, expected := range ss {
		l := tokenize(in)
		t.Logf("%q -> %#v", in, l)
		assert.Equal(t, expected, l)
	}
}

func TestIndexing(t *T) {
	t1 := testutil.RandStr()
	t2 := testutil.RandStr()
	t3 := testutil.RandStr()
	s := fiestatypes.Song{
		Uploaded: &fiestatypes.Uploaded{
			ID: fiestatypes.BlockID(testutil.RandStr()),
		},
		SongMeta: fiestatypes.SongMeta{
			Title: fmt.Sprintf("%s %s, %s", t1, t2, t3),
		},
	}
	require.Nil(t, setItem(s))
	require.Nil(t, indexItem(s))

	it, err := searchItems(t1)
	require.Nil(t, err)
	assert.Equal(t, s, it.Songs[0])
}
