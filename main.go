package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cockroachdb/pebble"
	"github.com/google/uuid"
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
	dec := json.NewDecoder(r.Body)
	var document map[string]interface{}
	err := dec.Decode(&document)
	if err != nil {
		jsonResponse(w, nil, err)
		return
	}

	id := uuid.New().String()

	bs, err := json.Marshal(document)
	if err != nil {
		jsonResponse(w, nil, err)
		return
	}
	err = s.db.Set([]byte(id), bs, pebble.Sync)
	if err != nil {
		jsonResponse(w, nil, err)
		return
	}

	jsonResponse(w, map[string]interface{}{
		"id": id,
	}, nil)
}

func (s server) searchDocuments(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	panic("Unimplemented")
}

func (s server) getDocument(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	panic("Unimplemented")
}

func jsonResponse(w http.ResponseWriter, body map[string]interface{}, err error) {
	data := map[string]interface{}{
		"body":   body,
		"status": "ok",
	}
	if err == nil {
		w.WriteHeader(http.StatusOK)
	} else {
		data["status"] = "error"
		data["error"] = err.Error()
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)
	err = enc.Encode(data)
	if err != nil {
		// TODO: set up panic handler?
		panic(err)
	}
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
