package kurz

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type server URLShortener

// Start listens for connections on given port
func Start(port int) error {

	localdb := newLocalDB()
	shortener := newURLShortener(*localdb)
	newServer := server(*shortener)

	http.HandleFunc("/encode/", newServer.serveEncodeURL)
	formattedPort := fmt.Sprintf(":%d", port)

	err := http.ListenAndServe(formattedPort, nil)
	return err
}

func (s *server) serveEncodeURL(w http.ResponseWriter, r *http.Request) {

}

func getURLParam(fixedPath string, path string) (string, error) {
	urlToEncode := strings.TrimPrefix(path, fixedPath)

	// if no path is found
	if strings.Compare(urlToEncode, fixedPath) == 0 {
		errorToReturn := errors.New("invalid url to encode, check fixedPath in URL")
		return "", errorToReturn
	}

	// decode url to encode from the given path
	parsedValue, err := base64.StdEncoding.DecodeString(urlToEncode)
	if err != nil {
		errorMessage := fmt.Sprintf("Unable to decode base64 string: %s", err)
		return "", errors.New(errorMessage)
	}
	return string(parsedValue), nil
}
