package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	. "testing"

	logging "gx/ipfs/Qmazh5oNUVsDZTs2g59rq8aYQqwpss8tcUWQzor5sCCEuH/go-log"

	"github.com/levenlabs/golib/rpcutil"
	"github.com/levenlabs/golib/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var log = logging.Logger("brian")

func init() {
	// TODO initialize a new ipfs for testing

	var err error
	n, err = newNode("~/.ipfs")
	if err != nil {
		log.Fatalf("newNode error: %s", err)
	}
}

func randSong(t *T) Song {
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

	s := Song{
		SongMeta: SongMeta{
			Title: testutil.RandStr(),
			Images: []SongImage{
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

func randPlaylist(t *T) Playlist {
	h := RPC()
	p := Playlist{
		Name: testutil.RandStr(),
		Songs: []Song{
			randSong(t),
			randSong(t),
			randSong(t),
		},
	}
	args := CreatePlaylistArgs{p}
	var res CreatePlaylistRes
	err := rpcutil.JSONRPC2CallHandler(h, &res, "Brian.CreatePlaylist", &args)
	require.Nil(t, err)

	return res.Playlist
}

func TestCreatePlaylist(t *T) {
	p := randPlaylist(t)
	assert.NotEmpty(t, p.Uploaded.ID)
	assert.NotEmpty(t, p.Uploaded.UploaderID)
}

func TestGetPlaylistByID(t *T) {
	h := RPC()
	p := randPlaylist(t)

	args := GetPlaylistByIDArgs{p.ID}
	var res GetPlaylistByIDRes
	err := rpcutil.JSONRPC2CallHandler(h, &res, "Brian.GetPlaylistByID", &args)
	require.Nil(t, err)
	assert.Equal(t, p, res.Playlist)
}
