package fiestatypes

import (
	. "testing"

	"github.com/stretchr/testify/assert"
)

func TestUploaded(t *T) {
	assert.Implements(t, (*Uploader)(nil), new(Uploaded))
	assert.Implements(t, (*Uploader)(nil), new(User))
	assert.Implements(t, (*Uploader)(nil), new(Playlist))
}
