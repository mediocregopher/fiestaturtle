package main

import (
	"bytes"
	"net/http"
	"os"

	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
	"github.com/levenlabs/golib/rpcutil"
	"github.com/mediocregopher/fiestaturtle/fiestatypes"
	"github.com/mediocregopher/fiestaturtle/richard/richard"
)

type Brian struct{}

func RPC() http.Handler {
	s := rpc.NewServer()
	s.RegisterCodec(json2.NewCodec(), "application/json")
	s.RegisterService(Brian{}, "")
	return s
}

func updateUser(fn func(*fiestatypes.User)) (fiestatypes.User, error) {
	u, err := getNSUser()
	if err != nil {
		return u, err
	}
	fn(&u)
	return putNSUser(u)
}

////////////////////////////////////////////////////////////////////////////////

type UploadSongArgs struct {
	Song fiestatypes.Song `json:"song"`

	// optional, not needed if Song is using a URL
	SongPath string `json:"songPath"`
}

type UploadSongRes struct {
	Song fiestatypes.Song `json:"song"`
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

	res.Song.Uploaded = &fiestatypes.Uploaded{}
	if _, err = put(&res.Song); err != nil {
		return err
	}

	return addRichardSong(res.Song)
}

func uploadSong(s *fiestatypes.Song, path string) error {
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

func uploadSongImgs(s *fiestatypes.Song) error {
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
	Playlist fiestatypes.Playlist `json:"playlist"`
	Replaces fiestatypes.BlockID  `json:"replaces"`
}

type CreatePlaylistRes struct {
	Playlist fiestatypes.Playlist `json:"playlist"`
}

func (_ Brian) CreatePlaylist(r *http.Request, args *CreatePlaylistArgs, res *CreatePlaylistRes) error {
	res.Playlist = args.Playlist
	res.Playlist.Uploaded = &fiestatypes.Uploaded{}
	if _, err := put(&res.Playlist); err != nil {
		return err
	}

	_, err := updateUser(func(u *fiestatypes.User) {
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
	if err != nil {
		return err
	}

	return addRichardPlaylist(res.Playlist)
}

////////////////////////////////////////////////////////////////////////////////

type GetPlaylistArgs struct {
	PlaylistID fiestatypes.BlockID `json:"playlistID"`
}

type GetPlaylistRes struct {
	Playlist fiestatypes.Playlist `json:"playlist"`
}

func (_ Brian) GetPlaylist(r *http.Request, args *GetPlaylistArgs, res *GetPlaylistRes) error {
	res.Playlist.Uploaded = &fiestatypes.Uploaded{}
	return get(args.PlaylistID, &res.Playlist)
}

////////////////////////////////////////////////////////////////////////////////

type GetUserArgs struct {
	// optional, returns this node's user if not set
	UserID fiestatypes.BlockID `json:"userID"`
}

type GetUserRes struct {
	User fiestatypes.User `json:"user"`
}

func (_ Brian) GetUser(r *http.Request, args *GetUserArgs, res *GetUserRes) error {
	var err error
	if args.UserID == "" {
		res.User, err = getNSUser()
		return err
	}

	return get(args.UserID, &res.User)
}

////////////////////////////////////////////////////////////////////////////////

type SetUserNameArgs struct {
	Name string `json:"name"`
}

type SetUserNameRes struct {
	User fiestatypes.User `json:"user"`
}

func (_ Brian) SetUserName(r *http.Request, args *SetUserNameArgs, res *SetUserNameRes) error {
	var err error
	res.User, err = updateUser(func(u *fiestatypes.User) {
		u.Name = args.Name
	})
	return err
}

////////////////////////////////////////////////////////////////////////////////

type DeletePlaylistArgs struct {
	PlaylistID fiestatypes.BlockID `json:"playlistID"`
}

func (_ Brian) DeletePlaylist(r *http.Request, args *DeletePlaylistArgs, res *struct{}) error {
	_, err := updateUser(func(u *fiestatypes.User) {
		newp := make([]fiestatypes.Playlist, 0, len(u.Playlists))
		for _, p := range u.Playlists {
			if p.ID == args.PlaylistID {
				continue
			}
			newp = append(newp, p)
		}
		u.Playlists = newp
	})
	return err
}

////////////////////////////////////////////////////////////////////////////////

type SetRichardsArgs struct {
	Richards []fiestatypes.Richard `json:"richards"`
}

func (_ Brian) SetRichards(r *http.Request, args *SetRichardsArgs, res *struct{}) error {
	_, err := updateUser(func(u *fiestatypes.User) {
		u.UserPrivate.Richards = args.Richards
	})
	return err
}

////////////////////////////////////////////////////////////////////////////////

type SearchArgs struct {
	Query string `json:"query"`
}

type SearchRes struct {
	Items richard.Items `json:"items"`
}

func (_ Brian) Search(r *http.Request, args *SearchArgs, res *SearchRes) error {
	var i richard.Items
	rargs := richard.SearchArgs{Query: args.Query}
	forEachRichard(func(r fiestatypes.Richard) error {
		res := richard.SearchRes{}
		err := rpcutil.JSONRPC2Call("http://"+r.Addr+"/rpc", &res, "Richard.Search", &rargs)
		if err != nil {
			return err
		}

		i.Songs = append(i.Songs, res.Items.Songs...)
		i.Playlists = append(i.Playlists, res.Items.Playlists...)
		i.Users = append(i.Users, res.Items.Users...)
		return nil
	})
	res.Items = i
	return nil
}
