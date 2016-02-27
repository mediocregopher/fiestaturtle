package main

import "time"

// Identifies a block in ipfs
type BlockID string

// Identifies an ipfs node (effectively, identifies a user)
type NodeID string

// Most top level types will have this, it describes the identifier for the
// object, who uploaded it, and a verification of those two facts
type Uploaded struct {
	ID          BlockID
	UploaderID  NodeID
	UploaderSig []byte
}

type Song struct {
	Version int
	Uploaded
	Name        string
	Artist      string
	Album       string
	ReleaseDate time.Time

	// If the song was uploaded this will be populated
	UploadedSongMeta UploadedSongMeta `json:",omitempty"`

	// If the song is a reference to soundcloud or youtube or whatever, this
	// will be populated
	URL string `json:",omitempty"`
}

type UploadedSongMeta struct {
	DataID  BlockID // BlockID that holds the actual mp3 data
	ImageID BlockID // BlockID that may hold the image
}

type Playlist struct {
	Version int
	Uploaded
	Songs []Song
}
