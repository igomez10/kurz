package kurz

import (
	"fmt"
	"net/http"
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
