package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
	"os"
	"path"
)

func keyFilePath() string {
	return path.Join(n.rdir, "super_secret_key")
}

func getKey() ([]byte, error) {
	kpath := keyFilePath()
	b, err := ioutil.ReadFile(kpath)
	if os.IsNotExist(err) {
		key := make([]byte, 16)
		if _, err := io.ReadFull(rand.Reader, key); err != nil {
			return nil, err
		}
		err := ioutil.WriteFile(kpath, key, 0400)
		return key, err
	}
	return b, err
}

func encrypt(b []byte) ([]byte, error) {
	key, err := getKey()
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aead, err := cipher.NewGCMWithNonceSize(block, 16)
	if err != nil {
		return nil, err
	}

	noncel := aead.NonceSize()
	out := make([]byte, noncel+len(b)+aead.Overhead())
	nonce := out[:noncel]
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	nout := aead.Seal(out[noncel:noncel], nonce, b, nil)
	return out[:noncel+len(nout)], nil
}

func decrypt(b []byte) ([]byte, error) {
	key, err := getKey()
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aead, err := cipher.NewGCMWithNonceSize(block, 16)
	if err != nil {
		return nil, err
	}

	noncel := aead.NonceSize()
	nonce := b[:noncel]
	b = b[noncel:]
	out := make([]byte, len(b))
	return aead.Open(out[:0], nonce, b, nil)
}
