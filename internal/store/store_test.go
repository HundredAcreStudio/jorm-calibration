package store

import (
	"testing"

	"github.com/HundredAcreStudio/jorm-calibration/internal/model"
)

func TestMemoryStore_CreateAndGet(t *testing.T) {
	s := NewMemoryStore()

	created, err := s.Create(model.User{Name: "Alice", Email: "alice@example.com"})
	if err != nil {
		t.Fatalf("Create: %v", err)
	}
	if created.ID == "" {
		t.Fatal("expected non-empty ID")
	}

	got, err := s.Get(created.ID)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if got.Name != "Alice" {
		t.Errorf("Name = %q, want %q", got.Name, "Alice")
	}
}

func TestMemoryStore_List(t *testing.T) {
	s := NewMemoryStore()

	s.Create(model.User{Name: "Alice"})
	s.Create(model.User{Name: "Bob"})

	users, err := s.List()
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if len(users) != 2 {
		t.Errorf("len = %d, want 2", len(users))
	}
}

func TestMemoryStore_Delete(t *testing.T) {
	s := NewMemoryStore()

	created, _ := s.Create(model.User{Name: "Alice"})
	if err := s.Delete(created.ID); err != nil {
		t.Fatalf("Delete: %v", err)
	}

	_, err := s.Get(created.ID)
	if err == nil {
		t.Error("expected error after delete")
	}
}

func TestMemoryStore_GetNotFound(t *testing.T) {
	s := NewMemoryStore()

	_, err := s.Get("nonexistent")
	if err == nil {
		t.Error("expected error for missing user")
	}
}
