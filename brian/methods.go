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

func updateUser(fn func(*User)) (User, error) {
	var u User
	u.Uploaded = &Uploaded{}
	_, err := getNS(&u)
	if err != nil && !os.IsNotExist(err) {
		return u, err
	}
	fn(&u)
	_, err = putNS(&u)
	return u, err
}

////////////////////////////////////////////////////////////////////////////////

type UploadSongArgs struct {
	Song Song `json:"song"`

	// optional, not needed if Song is using a URL
	SongPath string `json:"songPath"`
}

type UploadSongRes struct {
	Song Song `json:"song"`
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
	Replaces BlockID  `json:"replaces"`
}

type CreatePlaylistRes struct {
	Playlist Playlist `json:"playlist"`
}

func (_ Brian) CreatePlaylist(r *http.Request, args *CreatePlaylistArgs, res *CreatePlaylistRes) error {
	res.Playlist = args.Playlist
	res.Playlist.Uploaded = &Uploaded{}
	if _, err := put(&res.Playlist); err != nil {
		return err
	}

	_, err := updateUser(func(u *User) {
		p := res.Playlist
		p.Songs = nil
		if args.Replaces != "" {
			for i := range u.Playlists {
				if u.Playlists[i].ID == args.Replaces {
					u.Playlists[i] = p
					return
				}
			}
		}

		u.Playlists = append(u.Playlists, p)
	})
	return err
}

////////////////////////////////////////////////////////////////////////////////

type GetPlaylistArgs struct {
	PlaylistID BlockID `json:"playlistID"`
}

type GetPlaylistRes struct {
	Playlist Playlist `json:"playlist"`
}

func (_ Brian) GetPlaylist(r *http.Request, args *GetPlaylistArgs, res *GetPlaylistRes) error {
	res.Playlist.Uploaded = &Uploaded{}
	return get(args.PlaylistID, &res.Playlist)
}

////////////////////////////////////////////////////////////////////////////////

type GetUserArgs struct {
	// optional, returns this node's user if not set
	UserID BlockID `json:"userID"`
}

type GetUserRes struct {
	User User `json:"user"`
}

func (_ Brian) GetUser(r *http.Request, args *GetUserArgs, res *GetUserRes) error {
	var err error
	if args.UserID == "" {
		_, err = getNS(&res.User)
	} else {
		err = get(args.UserID, &res.User)
	}
	return err
}

////////////////////////////////////////////////////////////////////////////////

type SetUserNameArgs struct {
	Name string `json:"name"`
}

type SetUserNameRes struct {
	User User `json:"user"`
}

func (_ Brian) SetUserName(r *http.Request, args *SetUserNameArgs, res *SetUserNameRes) error {
	var err error
	res.User, err = updateUser(func(u *User) {
		u.Name = args.Name
	})
	return err
}
