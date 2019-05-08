package kurz

import (
	"crypto/sha1"
	"encoding/base64"
)

// DataContainer is a general interface where the entries will be stored
type DataContainer interface {
	Put([]byte, []byte)
	Get([]byte) []byte
}

// URLShortener is the main struct
type URLShortener struct {
	data DataContainer
}

// EncodeURL encodes the initial URL and returns a shortened version
func (u *URLShortener) EncodeURL(initialURL string) (string, error) {
	// TODO VALIDATE URL

	hasher := sha1.New()

	_, err := hasher.Write([]byte(initialURL))
	if err != nil {
		return "", err
	}

	hashed := hasher.Sum(nil)
	encoded := base64.StdEncoding.EncodeToString(hashed)

	u.data.Put([]byte(encoded), []byte(initialURL))

	return string(encoded), nil
}

// DecodeURL returns the full URL associated to the given shortened URL
func (u *URLShortener) DecodeURL(shortenedURL []byte) []byte {
	return u.data.Get(shortenedURL)
}

func newURLShortener(d DataContainer) *URLShortener {
	return &URLShortener{data: d}
}
