package nfcfyi

// SearchResult represents the API search response.
type SearchResult struct {
	Query   string       `json:"query"`
	Results []SearchItem `json:"results"`
	Total   int          `json:"total"`
}

// SearchItem represents a single search result item.
type SearchItem struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

// ChipDetail represents an NFC chip.
type ChipDetail struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Family      string `json:"family"`
	URL         string `json:"url"`
}

// ChipFamilyDetail represents an NFC chip family.
type ChipFamilyDetail struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// StandardDetail represents an NFC standard.
type StandardDetail struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// OperatingModeDetail represents an NFC operating mode.
type OperatingModeDetail struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// NdefTypeDetail represents an NDEF type.
type NdefTypeDetail struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// UseCaseDetail represents an NFC use case.
type UseCaseDetail struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// GlossaryTerm represents a glossary term.
type GlossaryTerm struct {
	Term       string `json:"term"`
	Slug       string `json:"slug"`
	Definition string `json:"definition"`
	URL        string `json:"url"`
}

// CompareResult represents a comparison between two NFC chips.
type CompareResult struct {
	ChipA interface{} `json:"chip_a"`
	ChipB interface{} `json:"chip_b"`
	URL   string      `json:"url"`
}
