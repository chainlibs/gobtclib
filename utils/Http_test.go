package utils

import "testing"

func TestNewClient(t *testing.T) {
	_, err := Http.NewSimpleClient()
	if err != nil {
		t.Error(err)
	}
}
