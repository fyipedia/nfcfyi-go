// Package nfcfyi provides a Go client for the NFCFYI API.
//
// NFCFYI is a comprehensive NFC reference covering chips, chip families,
// standards, operating modes, NDEF types, and use cases. This client requires
// no authentication and has zero external dependencies.
//
// Usage:
//
//	client := nfcfyi.NewClient()
//	result, err := client.Search("ntag")
package nfcfyi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// DefaultBaseURL is the default base URL for the NFCFYI API.
const DefaultBaseURL = "https://nfcfyi.com/api"

// Client is an NFCFYI API client.
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient creates a new NFCFYI API client with default settings.
func NewClient() *Client {
	return &Client{
		BaseURL:    DefaultBaseURL,
		HTTPClient: &http.Client{},
	}
}

func (c *Client) get(path string, result interface{}) error {
	resp, err := c.HTTPClient.Get(c.BaseURL + path)
	if err != nil {
		return fmt.Errorf("nfcfyi: request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("nfcfyi: HTTP %d: %s", resp.StatusCode, string(body))
	}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return fmt.Errorf("nfcfyi: decode failed: %w", err)
	}
	return nil
}

// Search searches across NFC chips, standards, and glossary terms.
func (c *Client) Search(query string) (*SearchResult, error) {
	var result SearchResult
	path := fmt.Sprintf("/search/?q=%s", url.QueryEscape(query))
	if err := c.get(path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Chip returns details for an NFC chip by slug.
func (c *Client) Chip(slug string) (*ChipDetail, error) {
	var result ChipDetail
	if err := c.get("/chip/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ChipFamily returns details for an NFC chip family by slug.
func (c *Client) ChipFamily(slug string) (*ChipFamilyDetail, error) {
	var result ChipFamilyDetail
	if err := c.get("/chip-family/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Standard returns details for an NFC standard by slug.
func (c *Client) Standard(slug string) (*StandardDetail, error) {
	var result StandardDetail
	if err := c.get("/standard/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// OperatingMode returns details for an NFC operating mode by slug.
func (c *Client) OperatingMode(slug string) (*OperatingModeDetail, error) {
	var result OperatingModeDetail
	if err := c.get("/operating-mode/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// NdefType returns details for an NDEF type by slug.
func (c *Client) NdefType(slug string) (*NdefTypeDetail, error) {
	var result NdefTypeDetail
	if err := c.get("/ndef-type/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UseCase returns details for an NFC use case by slug.
func (c *Client) UseCase(slug string) (*UseCaseDetail, error) {
	var result UseCaseDetail
	if err := c.get("/use-case/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GlossaryTerm returns a glossary term by slug.
func (c *Client) GlossaryTerm(slug string) (*GlossaryTerm, error) {
	var result GlossaryTerm
	if err := c.get("/glossary/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Compare compares two NFC chips.
func (c *Client) Compare(slugA, slugB string) (*CompareResult, error) {
	var result CompareResult
	path := fmt.Sprintf("/compare/?a=%s&b=%s", url.QueryEscape(slugA), url.QueryEscape(slugB))
	if err := c.get(path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Random returns a random NFC chip.
func (c *Client) Random() (*ChipDetail, error) {
	var result ChipDetail
	if err := c.get("/random/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}
