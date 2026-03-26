package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/HundredAcreStudio/jorm-calibration/internal/store"
)

func TestCreateUser(t *testing.T) {
	db := store.NewMemoryStore()
	h := CreateUser(db)

	body := `{"name":"Alice","email":"alice@example.com"}`
	req := httptest.NewRequest("POST", "/users", strings.NewReader(body))
	w := httptest.NewRecorder()
	h(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("status = %d, want %d", w.Code, http.StatusCreated)
	}
	if !strings.Contains(w.Body.String(), "Alice") {
		t.Error("response should contain user name")
	}
}

func TestListUsers(t *testing.T) {
	db := store.NewMemoryStore()
	h := ListUsers(db)

	req := httptest.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	h(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", w.Code, http.StatusOK)
	}
}
