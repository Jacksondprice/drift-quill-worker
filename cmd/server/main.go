package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Jacksondprice/drift-quill-worker/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/assignments", handlers.ListAssignments)
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}
	log.Fatal(http.ListenAndServe(addr, mux))
}
