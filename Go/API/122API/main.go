package main

import (
	"log"
	"net/http"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()
	router.HandleFunc("/users/{userID}", func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Path
		w.Write([]byte("User ID: " + userID))
	})

	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	log.Printf("Server has started %s", s.addr)

	return server.ListenAndServe()
}

func main() {
	apiServer := NewAPIServer(":8080")
	log.Fatal(apiServer.Run())
}
