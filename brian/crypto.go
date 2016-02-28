package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path"

	"github.com/mediocregopher/fiestaturtle/fiestatypes"
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

func encryptUser(u *fiestatypes.User) error {
	b, err := json.Marshal(u.UserPrivate)
	if err != nil {
		return err
	}

	benc, err := encrypt(b)
	if err != nil {
		return err
	}

	u.UserPrivateRaw = benc
	u.UserPrivate = fiestatypes.UserPrivate{}
	return nil
}

func decryptUser(u *fiestatypes.User) error {
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

func signUploader(up fiestatypes.Uploader) error {
	u := up.Get()
	u.UploaderID = getNodeID()
	u.UploaderSig = []byte{} // TODO actually do this
	up.Set(u)
	return nil
}
