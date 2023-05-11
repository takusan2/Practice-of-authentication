package database

import (
	"testing"
)

func TestConnect(t *testing.T) {
	db := Connect()
	if db == nil {
		t.Errorf("Connect() == %v, want %v", db, "not nil")
	}
}
