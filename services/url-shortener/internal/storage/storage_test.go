package storage

import (
	"errors"
	"testing"
)

func TestErrorSentinelsAreDistinct(t *testing.T) {
	if ErrConflict == nil {
		t.Fatal("ErrConflict is nil")
	}
	if ErrNotFound == nil {
		t.Fatal("ErrNotFound is nil")
	}
	if errors.Is(ErrConflict, ErrNotFound) {
		t.Fatal("ErrConflict matches ErrNotFound")
	}
}
