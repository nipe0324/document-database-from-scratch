package main

import (
	"log"
	"net/http"

	"github.com/cockroachdb/pebble"
	"github.com/julienschmidt/httprouter"
)

type server struct {
	db   *pebble.DB
	port string
}

func newServer(database string, port string) (*server, error) {
	s := server{db: nil, port: port}
	var err error
	s.db, err = pebble.Open(database, &pebble.Options{})
	return &s, err
}

func (s server) addDocument(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	panic("Unimplemented")
}

func (s server) searchDocuments(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	panic("Unimplemented")
}

func (s server) getDocument(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	panic("Unimplemented")
}

func main() {
	s, err := newServer("docdb.data", "8080")
	if err != nil {
		log.Fatal(err)
	}
	defer s.db.Close()

	router := httprouter.New()
	router.POST("/docs", s.addDocument)
	router.GET("/docs", s.searchDocuments)
	router.GET("/docs/:id", s.getDocument)

	log.Println("Listening on " + s.port)
	log.Fatal(http.ListenAndServe(":"+s.port, router))
}
