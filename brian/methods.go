package main

import (
	"bytes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
)

type Brian struct{}

func RPC() http.Handler {
	s := rpc.NewServer()
	s.RegisterCodec(json2.NewCodec(), "application/json")
	s.RegisterService(Brian{}, "")
	return s
}

////////////////////////////////////////////////////////////////////////////////

type UploadSongArgs struct {
	Song Song

	// optional, not needed if Song is using a URL
	SongPath string
}

type UploadSongRes struct {
	Song Song
}

func (_ Brian) UploadSongData(r *http.Request, args *UploadSongArgs, res *UploadSongRes) error {
	res.Song = args.Song
	var err error
	if args.SongPath != "" {
		if err := uploadSong(&res.Song, args.SongPath); err != nil {
			return err
		}
	}

	if err := uploadSongImgs(&res.Song); err != nil {
		return err
	}

	res.Song.Uploaded = &Uploaded{}
	_, err = put(&res.Song)
	return err
}

func uploadSong(s *Song, path string) error {
	log.Printf("uploading song file at %q", path)
	songF, err := os.Open(path)
	if err != nil {
		return err
	}

	songID, err := putRaw(songF)
	songF.Close()
	if err != nil {
		return err
	}

	s.UploadedSongMeta.DataID = songID
	log.Printf("done uploading %q (%q)", path, songID)
	return nil
}

func uploadSongImgs(s *Song) error {
	for i, img := range s.SongMeta.Images {
		log.Printf("uploading song image")
		imgID, err := putRaw(bytes.NewBuffer(img.Data))
		if err != nil {
			return err
		}
		s.SongMeta.Images[i].ID = imgID
		s.SongMeta.Images[i].Data = nil
		log.Printf("done uploading song image (%q)", imgID)
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type CreateSongArgs struct {
	Song Song
}
