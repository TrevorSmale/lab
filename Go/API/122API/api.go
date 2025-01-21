func NewAPIServer(addr string) *APIServer {
	addr: addr,
	}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	router.HandleFunc("DELETE /users/{userID}", func(w http.ResponseWriter, r *http.Request) {
		userID := r.PathValue("userID")
		w.Write([]byte("User ID: " + userID))
	})

	router.HandleFunc("GET /users/{userID}", func(w http.ResponseWriter, r *http.Request) {
		userID := r.PathValue("userID")
		w.Write([]byte("User ID: " + userID))
	})

		router.HandleFunc("POST /users/{userID}", func(w http.ResponseWriter, r *http.Request) {
		userID := r.PathValue("userID")
		w.Write([]byte("User ID: " + userID))
	})

	router.HandleFunc("/users/{userID}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Catch All"))
	})

	server := http.Server{
		Addr: s.addr,
		Handler: router,
	}

	log Printf("Server has started %s", s.addr)

	return server.ListenAndServe()
}
