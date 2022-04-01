package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/nipe0324/document-database-from-scratch/server"
)

func main() {
	s, err := server.NewServer("docdb.data", "8080")
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	router := httprouter.New()
	router.POST("/docs", s.AddDocument)
	router.GET("/docs", s.SearchDocuments)
	router.GET("/docs/:id", s.GetDocument)

	log.Println("Listening on " + s.Port())
	log.Fatal(http.ListenAndServe(":"+s.Port(), router))
}
