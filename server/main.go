package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func health(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s", r.Method, r.Host, r.URL.Path)

	if r.URL.Path != "/health" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "ok")
}

func main() {
	srv := http.NewServeMux()
	srv.HandleFunc("/", health)

	port := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	if port == ":" {
		port = ":8080"
	}

	log.Printf("Listening on port %v", port)
	http.ListenAndServe(port, srv)
}
