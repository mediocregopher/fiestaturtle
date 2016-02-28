package main

import (
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
