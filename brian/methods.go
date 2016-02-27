package main

import (
	"bytes"
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
	return nil
}

func uploadSongImgs(s *Song) error {
	for i, img := range s.SongMeta.Images {
		imgID, err := putRaw(bytes.NewBuffer(img.Data))
		if err != nil {
			return err
		}
		s.SongMeta.Images[i].ID = imgID
		s.SongMeta.Images[i].Data = nil
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type CreatePlaylistArgs struct {
	Playlist Playlist `json:"playlist"`
}

type CreatePlaylistRes struct {
	Playlist Playlist `json:"playlist"`
}

func (_ Brian) CreatePlaylist(r *http.Request, args *CreatePlaylistArgs, res *CreatePlaylistRes) error {
	res.Playlist = args.Playlist
	res.Playlist.Uploaded = &Uploaded{}
	_, err := put(&res.Playlist)
	return err
}

////////////////////////////////////////////////////////////////////////////////

type GetPlaylistByIDArgs struct {
	PlaylistID BlockID `json:"playlistID"`
}

type GetPlaylistByIDRes struct {
	Playlist Playlist `json:"playlist"`
}

func (_ Brian) GetPlaylistByID(r *http.Request, args *GetPlaylistByIDArgs, res *GetPlaylistByIDRes) error {
	res.Playlist.Uploaded = &Uploaded{}
	return get(args.PlaylistID, &res.Playlist)
}
