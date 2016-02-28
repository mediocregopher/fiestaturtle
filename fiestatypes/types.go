package fiestatypes

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
	Richards []Richard `json:"richards"`
}

type Richard struct {
	Addr string `json:"addr"`
}

////////////////////////////////////////////////////////////////////////////////

type Uploader interface {
	Get() *Uploaded
	Set(*Uploaded)
	SetID(BlockID)
}

func (u *Uploaded) Get() *Uploaded {
	if u == nil {
		return &Uploaded{}
	}
	return u
}

func (u *Uploaded) Set(uin *Uploaded) {
	*u = *uin
}

func (u *Uploaded) SetID(id BlockID) {
	u.ID = id
}
