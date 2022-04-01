package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type server struct {
	port string
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
	s := server{"8080"}
	router := httprouter.New()
	router.POST("/docs", s.addDocument)
	router.GET("/docs", s.searchDocuments)
	router.GET("/docs/:id", s.getDocument)

	log.Println("Listening on " + s.port)
	log.Fatal(http.ListenAndServe(":"+s.port, router))
}
