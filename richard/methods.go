package richard

import (
	"net/http"

	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
	"github.com/mediocregopher/fiestaturtle/fiestatypes"
)

type Richard struct{}

func RPC() http.Handler {
	s := rpc.NewServer()
	s.RegisterCodec(json2.NewCodec(), "application/json")
	s.RegisterService(Richard{}, "")
	return s
}

////////////////////////////////////////////////////////////////////////////////

type AddSongArgs struct {
	Song fiestatypes.Song `json:"song"`
}

func (_ Richard) AddSong(r *http.Request, args *AddSongArgs, res *struct{}) error {
	if err := setItem(args.Song); err != nil {
		return err
	}
	return indexItem(args.Song)
}

////////////////////////////////////////////////////////////////////////////////

type AddPlaylistArgs struct {
	Playlist fiestatypes.Playlist `json:"playlist"`
}

func (_ Richard) AddPlaylist(r *http.Request, args *AddPlaylistArgs, res *struct{}) error {
	if err := setItem(args.Playlist); err != nil {
		return err
	}
	return indexItem(args.Playlist)
}

////////////////////////////////////////////////////////////////////////////////

type AddUserArgs struct {
	User fiestatypes.User `json:"user"`
}

func (_ Richard) AddUser(r *http.Request, args *AddUserArgs, res *struct{}) error {
	if err := setItem(args.User); err != nil {
		return err
	}
	return indexItem(args.User)
}
