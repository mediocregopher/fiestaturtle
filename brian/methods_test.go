package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	. "testing"

	"github.com/levenlabs/golib/rpcutil"
	"github.com/levenlabs/golib/testutil"
	"github.com/mediocregopher/fiestaturtle/fiestatypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	var err error
	n, err = newNode(false)
	if err != nil {
		panic(err)
	}

	noRichardOps = true

	os.Remove(nsFilePath())
	os.Remove(keyFilePath())
}

func getUser(t *T) fiestatypes.User {
	h := RPC()
	args := GetUserArgs{}
	res := GetUserRes{}
	err := rpcutil.JSONRPC2CallHandler(h, &res, "Brian.GetUser", &args)
	require.Nil(t, err)
	return res.User
}

func randSong(t *T) fiestatypes.Song {
	h := RPC()
	songContent := []byte(testutil.RandStr())
	imgContent := []byte(testutil.RandStr())

	tmpDir, err := ioutil.TempDir("", "brian")
	require.Nil(t, err)
	songPath := path.Join(tmpDir, testutil.RandStr())

	songF, err := os.Create(songPath)
	require.Nil(t, err)
	_, err = fmt.Fprint(songF, string(songContent))
	require.Nil(t, err)
	songF.Close()

	s := fiestatypes.Song{
		SongMeta: fiestatypes.SongMeta{
			Title: testutil.RandStr(),
			Images: []fiestatypes.SongImage{
				{
					Data: imgContent,
				},
			},
		},
	}
	args := UploadSongArgs{
		Song:     s,
		SongPath: songPath,
	}
	var res UploadSongRes
	err = rpcutil.JSONRPC2CallHandler(h, &res, "Brian.UploadSongData", &args)
	require.Nil(t, err)

	return res.Song
}

func TestUploadSongData(t *T) {
	s := randSong(t)
	assert.NotEmpty(t, s.Uploaded.ID)
	assert.NotEmpty(t, s.Uploaded.UploaderID)

	assert.NotEmpty(t, s.UploadedSongMeta.DataID)

	assert.Empty(t, s.SongMeta.Images[0].Data)
	assert.NotEmpty(t, s.SongMeta.Images[0].ID)
}

func randPlaylist(t *T) fiestatypes.Playlist {
	h := RPC()
	p := fiestatypes.Playlist{
		Name: testutil.RandStr(),
		Songs: []fiestatypes.Song{
			randSong(t),
			randSong(t),
			randSong(t),
		},
	}
	args := CreatePlaylistArgs{Playlist: p}
	var res CreatePlaylistRes
	err := rpcutil.JSONRPC2CallHandler(h, &res, "Brian.CreatePlaylist", &args)
	require.Nil(t, err)

	return res.Playlist
}

func TestCreatePlaylist(t *T) {
	p := randPlaylist(t)
	assert.NotEmpty(t, p.Uploaded.ID)
	assert.NotEmpty(t, p.Uploaded.UploaderID)

	p.Songs = nil
	u := getUser(t)
	require.Len(t, u.Playlists, 1)
	assert.Equal(t, p, u.Playlists[0])

	p.Name = testutil.RandStr()
	h := RPC()
	args := CreatePlaylistArgs{Playlist: p, Replaces: p.ID}
	var res CreatePlaylistRes
	err := rpcutil.JSONRPC2CallHandler(h, &res, "Brian.CreatePlaylist", &args)
	require.Nil(t, err)

	u = getUser(t)
	require.Len(t, u.Playlists, 1)
	assert.Equal(t, res.Playlist, u.Playlists[0])
}

func TestGetPlaylist(t *T) {
	h := RPC()
	p := randPlaylist(t)

	args := GetPlaylistArgs{p.ID}
	var res GetPlaylistRes
	err := rpcutil.JSONRPC2CallHandler(h, &res, "Brian.GetPlaylist", &args)
	require.Nil(t, err)
	assert.Equal(t, p, res.Playlist)
}

func TestSetUserName(t *T) {
	h := RPC()
	args := SetUserNameArgs{Name: testutil.RandStr()}
	var res SetUserNameRes
	err := rpcutil.JSONRPC2CallHandler(h, &res, "Brian.SetUserName", &args)
	require.Nil(t, err)
	assert.Equal(t, args.Name, res.User.Name)

	u := getUser(t)
	assert.Equal(t, args.Name, u.Name)
}

func TestDeletePlaylist(t *T) {
	h := RPC()
	p1 := randPlaylist(t)

	args := DeletePlaylistArgs{PlaylistID: p1.ID}
	err := rpcutil.JSONRPC2CallHandler(h, &struct{}{}, "Brian.DeletePlaylist", &args)
	require.Nil(t, err)

	u := getUser(t)
	for _, p := range u.Playlists {
		assert.NotEqual(t, p1.ID, p)
	}
}

func TestSetRichards(t *T) {
	h := RPC()
	args := SetRichardsArgs{
		Richards: []fiestatypes.Richard{
			{Addr: testutil.RandStr()},
			{Addr: testutil.RandStr()},
			{Addr: testutil.RandStr()},
		},
	}
	err := rpcutil.JSONRPC2CallHandler(h, &struct{}{}, "Brian.SetRichards", &args)
	require.Nil(t, err)

	u := getUser(t)
	assert.Equal(t, args.Richards, u.UserPrivate.Richards)
}
