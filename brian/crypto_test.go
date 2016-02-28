package main

import (
	. "testing"

	"github.com/levenlabs/golib/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCrypto(t *T) {
	b := []byte(testutil.RandStr())
	t.Logf("b:   %x", b)

	enc, err := encrypt(b)
	t.Logf("enc: %x", enc)
	require.Nil(t, err)

	b2, err := decrypt(enc)
	t.Logf("b2:  %x", b2)
	require.Nil(t, err)
	assert.Equal(t, b, b2)
}
