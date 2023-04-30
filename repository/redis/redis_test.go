package redis

import "testing"

func TestConnet(t *testing.T) {
	db := Connect()
	if db == nil {
		t.Errorf("Connect() == %v, want %v", db, "not nil")
	}
}
