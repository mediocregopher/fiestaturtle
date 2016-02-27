package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	. "testing"

	"github.com/levenlabs/golib/rpcutil"
	"github.com/levenlabs/golib/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	// TODO initialize a new ipfs for testing

	var err error
	n, err = newNode("~/.ipfs")
	if err != nil {
		log.Printf("newNode error: %s", err)
	}
}

func TestUploadSongData(t *T) {
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

	assert.NotEmpty(t, res.Song.Uploaded.ID)
	assert.NotEmpty(t, res.Song.Uploaded.UploaderID)
	s.Uploaded = res.Song.Uploaded

	assert.NotEmpty(t, res.Song.UploadedSongMeta.DataID)
	s.UploadedSongMeta = res.Song.UploadedSongMeta

	assert.Empty(t, res.Song.SongMeta.Images[0].Data)
	assert.NotEmpty(t, res.Song.SongMeta.Images[0].ID)
	s.SongMeta.Images = res.Song.SongMeta.Images

	assert.Equal(t, s, res.Song)
}
