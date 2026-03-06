package nfcfyi_test

import (
	"testing"

	nfcfyi "github.com/fyipedia/nfcfyi-go"
)

func TestNewClient(t *testing.T) {
	c := nfcfyi.NewClient()
	if c.BaseURL != nfcfyi.DefaultBaseURL {
		t.Errorf("expected %s, got %s", nfcfyi.DefaultBaseURL, c.BaseURL)
	}
	if c.HTTPClient == nil {
		t.Error("expected non-nil HTTPClient")
	}
}

func TestSearch(t *testing.T) {
	c := nfcfyi.NewClient()
	result, err := c.Search("ntag")
	if err != nil {
		t.Fatalf("Search failed: %v", err)
	}
	if result.Total == 0 {
		t.Error("expected results, got 0")
	}
}
