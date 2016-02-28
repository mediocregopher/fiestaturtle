package main

import "encoding/json"

// Identifies a block in ipfs
type BlockID string

// Identifies an ipfs node (effectively, identifies a user)
type NodeID string

// Most top level types will have this, it describes the identifier for the
// object, who uploaded it, and a verification of those two facts
type Uploaded struct {
	ID          BlockID `json:"id"`
	UploaderID  NodeID  `json:"uploaderID"`
	UploaderSig []byte  `json:"uploaderSig"`
}

type Song struct {
	Version int `json:"version"`
	*Uploaded

	SongMeta SongMeta `json:"meta"`

	// If the song was uploaded this will be populated
	UploadedSongMeta UploadedSongMeta `json:"uploadedSongMeta,omitempty"`

	// If the song is a reference to soundcloud or youtube or whatever, this
	// will be populated
	URL string `json:"url,omitempty"`
}

type SongMeta struct {
	Title        string      `json:"title,omitempty"`
	Artists      []string    `json:"artist,omitempty"`
	Album        string      `json:"album,omitempty"`
	AlbumArtists []string    `json:"albumartist,omitempty"`
	Year         string      `json:"year,omitempty"`
	Track        NoOf        `json:"track,omitempty"`
	Disk         NoOf        `json:"disk,omitempty"`
	Genres       []string    `json:"genre,omitempty"`
	Images       []SongImage `json:"picture,omitempty"`
	Duration     float64     `json:"duration,omitempty"`
}

type NoOf struct {
	No int `json:"no"`
	Of int `json:"of,omitempty"`
}

type SongImage struct {
	Format string  `json:"format"`
	Data   []byte  `json:"data"` // only populated by marco when creating song, ID is populated in all other cases
	ID     BlockID `json:"id"`
}

type UploadedSongMeta struct {
	DataID BlockID `json:"dataID"` // BlockID that holds the actual mp3 data
}

type Playlist struct {
	Version int `json:"version"`
	*Uploaded
	Name  string `json:"name"`
	Songs []Song `json:"songs"`
}

type User struct {
	Version int `json:"version"`
	*Uploaded
	Name      string     `json:"name"`
	Playlists []Playlist `json:"playlists"` // Songs won't be filled in on these

	UserPrivate    UserPrivate `json:"private"`
	UserPrivateRaw []byte      `json:"privateRaw"`
}

type UserPrivate struct {
	Richards []Richard
}

type Richard struct {
	Addr string
}

////////////////////////////////////////////////////////////////////////////////

type uploader interface {
	uploaded(BlockID)
	sign() error
}

func (u *Uploaded) sign() error {
	u.UploaderID = getNodeID()
	u.UploaderSig = []byte{} // TODO actually do this
	return nil
}

func (u *Uploaded) uploaded(id BlockID) {
	u.ID = id
}

////////////////////////////////////////////////////////////////////////////////

func (u *User) encrypt() error {
	b, err := json.Marshal(u.UserPrivate)
	if err != nil {
		return err
	}

	benc, err := encrypt(b)
	if err != nil {
		return err
	}

	u.UserPrivateRaw = benc
	u.UserPrivate = UserPrivate{}
	return nil
}

func (u *User) decrypt() error {
	if len(u.UserPrivateRaw) == 0 {
		return nil
	}

	b, err := decrypt(u.UserPrivateRaw)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, &u.UserPrivate); err != nil {
		return err
	}

	u.UserPrivateRaw = nil
	return nil
}
