package kurz

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"
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

// HandleEncoder handles the call for encoding a URL
func (u *URLShortener) HandleEncoder(w http.ResponseWriter, r *http.Request) {
	path := r.URL.EscapedPath() // expected format /encode/<base64_encodedURL>
	splitted := strings.SplitN(path, "/encode/", 1)

	if len(splitted) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid resource")
		return
	}

	base64URLToEncode := splitted[1]

	urlToEncode := base64.StdEncoding.EncodeToString([]byte(base64URLToEncode))

	storedKey, err := u.EncodeURL(urlToEncode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Unexpected Error encoding URL")
		log.Println(err)
		return
	}

	fmt.Fprint(w, storedKey)
}
