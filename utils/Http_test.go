package utils

import "testing"

func TestNewClient(t *testing.T) {
	_, err := Http.NewClientSimple()
	if err != nil {
		t.Error(err)
	}
}
