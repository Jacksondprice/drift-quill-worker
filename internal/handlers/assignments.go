package handlers

import (
	"encoding/json"
	"net/http"
)

type Assignment struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ListAssignments(w http.ResponseWriter, _ *http.Request) {
	items := []Assignment{{ID: 1, Name: "sample"}}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(items)
}
