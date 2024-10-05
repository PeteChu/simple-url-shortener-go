package main

import (
	"fmt"
	"log"
	"net/http"
)

type HTTPServer struct {
	db     map[string]string
	server http.Server
}

func shortenUrl() {}

func (s *HTTPServer) index(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)

		if len(s.db) == 0 {
			fmt.Fprint(w, "No data. Satuaz")
			return
		}

		for k, v := range s.db {
			fmt.Fprintf(w, "Key: %s, Link: %q", k, v)
		}

	case http.MethodPost:
		w.WriteHeader(http.StatusCreated)
		url := r.FormValue("url")
		k := genUrlKey(url)
		s.db[k] = url
		fmt.Fprintf(w, "Link: localhost:3000/%s", k)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "%q", "Method not allowed.")
	}
}

func (s *HTTPServer) redirect(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	k := r.URL.Path[1:]

	val, ok := s.db[k]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	http.Redirect(w, r, val, http.StatusSeeOther)
}

func StartServer() {

	mux := http.NewServeMux()

	s := HTTPServer{
		db: make(map[string]string),
		server: http.Server{
			Addr:    ":3000",
			Handler: mux,
		},
	}

	mux.HandleFunc("/", s.redirect)
	mux.HandleFunc("/shrturl", s.index)

	log.Fatal(s.server.ListenAndServe())

}
