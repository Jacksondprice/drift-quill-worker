package handlers

import (
	"net/http/httptest"
	"testing"
)

func TestListAssignments(t *testing.T) {
	rec := httptest.NewRecorder()
	ListAssignments(rec, httptest.NewRequest("GET", "/api/assignments", nil))
	if rec.Code != 200 {
		t.Fatalf("status %d", rec.Code)
	}
}
