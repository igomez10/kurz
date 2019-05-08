package kurz

import (
	"net/http"
)

type server URLShortener

func Start() {

	localdb := newLocalDB()
	shortener := newURLShortener(*localdb)
	newServer := server(*shortener)

	http.HandleFunc("/encode/", newServer.serveEncodeURL)
}

func (s *server) serveEncodeURL(w http.ResponseWriter, r *http.Request) {

}
