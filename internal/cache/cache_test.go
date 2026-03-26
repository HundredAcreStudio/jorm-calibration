package cache

import (
	"testing"
	"time"
)

func TestCache_SetAndGet(t *testing.T) {
	c := New(1 * time.Minute)
	c.Set("key", "value")

	got := c.Get("key")
	if got != "value" {
		t.Errorf("Get = %v, want %q", got, "value")
	}
}

func TestCache_GetExpired(t *testing.T) {
	c := New(1 * time.Millisecond)
	c.Set("key", "value")

	time.Sleep(5 * time.Millisecond)

	got := c.Get("key")
	if got != nil {
		t.Errorf("expected nil for expired key, got %v", got)
	}
}

func TestCache_Delete(t *testing.T) {
	c := New(1 * time.Minute)
	c.Set("key", "value")
	c.Delete("key")

	got := c.Get("key")
	if got != nil {
		t.Errorf("expected nil after delete, got %v", got)
	}
}
