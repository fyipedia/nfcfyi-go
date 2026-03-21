# nfcfyi-go

[![Go Reference](https://pkg.go.dev/badge/github.com/fyipedia/nfcfyi-go.svg)](https://pkg.go.dev/github.com/fyipedia/nfcfyi-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Go client for the [NFCFYI](https://nfcfyi.com) REST API. NFC chips, tag types, NDEF. Zero external dependencies beyond stdlib.

> **Try the interactive tools at [nfcfyi.com](https://nfcfyi.com)** — explore, search, and discover.

## Install

```bash
go get github.com/fyipedia/nfcfyi-go
```

## Quick Start

```go
package main

import (
    "fmt"
    nfcfyi "github.com/fyipedia/nfcfyi-go"
)

func main() {
    client := nfcfyi.NewClient()

    // Search across all content
    result, err := client.Search("query")
    if err != nil {
        panic(err)
    }
    fmt.Println(result)
}
```

## API Methods

| Method | Description |
|--------|-------------|
| `ChipFamilies()` | List chip families |
| `Chips()` | List chips |
| `Faqs()` | List faqs |
| `FrequencyBands()` | List frequency bands |
| `Glossary()` | List glossary |
| `Guides()` | List guides |
| `Manufacturers()` | List manufacturers |
| `NdefTypes()` | List ndef types |
| `OperatingModes()` | List operating modes |
| `SecurityProtocols()` | List security protocols |
| `Standards()` | List standards |
| `TagTypes()` | List tag types |
| `Tools()` | List tools |
| `UseCases()` | List use cases |
| `Search(query)` | Search across all content |

## Also Available

| Platform | Package | Link |
|----------|---------|------|
| **Python** | `pip install nfcfyi` | [PyPI](https://pypi.org/project/nfcfyi/) |
| **npm** | `npm install nfcfyi` | [npm](https://www.npmjs.com/package/nfcfyi) |
| **Go** | `go get github.com/fyipedia/nfcfyi-go` | [pkg.go.dev](https://pkg.go.dev/github.com/fyipedia/nfcfyi-go) |
| **Rust** | `cargo add nfcfyi` | [crates.io](https://crates.io/crates/nfcfyi) |
| **Ruby** | `gem install nfcfyi` | [rubygems](https://rubygems.org/gems/nfcfyi) |

## Tag FYI Family

Part of the [FYIPedia](https://fyipedia.com) open-source developer tools ecosystem.

| Site | Domain | Focus |
|------|--------|-------|
| BarcodeFYI | [barcodefyi.com](https://barcodefyi.com) | Barcode formats, EAN, UPC, ISBN standards |
| BLEFYI | [blefyi.com](https://blefyi.com) | Bluetooth Low Energy, GATT, beacons |
| **NFCFYI** | [nfcfyi.com](https://nfcfyi.com) | NFC chips, tag types, NDEF, standards |
| QRCodeFYI | [qrcodefyi.com](https://qrcodefyi.com) | QR code types, versions, encoding modes |
| RFIDFYI | [rfidfyi.com](https://rfidfyi.com) | RFID tags, frequency bands, standards |
| SmartCardFYI | [smartcardfyi.com](https://smartcardfyi.com) | Smart cards, EMV, APDU, Java Card |

## License

MIT
