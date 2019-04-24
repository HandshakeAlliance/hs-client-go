package http

import (
	"testing"
)

func TestInfo(t *testing.T) {
	hostname := "http://localhost:13037"
	hsClient := New(hostname)

	info, err := hsClient.Info()

	if err != nil {
		t.Errorf("Info Endpoint Failed with the following error: %v", err)
	}
}
