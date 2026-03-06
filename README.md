# nfcfyi-go

[![Go Reference](https://pkg.go.dev/badge/github.com/fyipedia/nfcfyi-go.svg)](https://pkg.go.dev/github.com/fyipedia/nfcfyi-go)

Go client for the [NFCFYI](https://nfcfyi.com) API. Look up NFC chips, chip families, standards, operating modes, NDEF types, and use cases. Zero dependencies — stdlib only.

## Install

```bash
go get github.com/fyipedia/nfcfyi-go
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"

    nfcfyi "github.com/fyipedia/nfcfyi-go"
)

func main() {
    client := nfcfyi.NewClient()

    result, err := client.Search("ntag")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Found %d results for 'ntag'\n", result.Total)

    chip, err := client.Chip("ntag215")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s: %s\n", chip.Name, chip.Description)
}
```

## API Methods

| Method | Description |
|--------|-------------|
| `Search(query)` | Search chips, standards, and glossary |
| `Chip(slug)` | Get NFC chip details |
| `ChipFamily(slug)` | Get chip family details |
| `Standard(slug)` | Get standard details |
| `OperatingMode(slug)` | Get operating mode details |
| `NdefType(slug)` | Get NDEF type details |
| `UseCase(slug)` | Get use case details |
| `GlossaryTerm(slug)` | Get glossary term definition |
| `Compare(slugA, slugB)` | Compare two NFC chips |
| `Random()` | Get a random NFC chip |

## REST API

Free, no authentication required. Base URL: `https://nfcfyi.com/api`

```bash
curl https://nfcfyi.com/api/search/?q=ntag
curl https://nfcfyi.com/api/chip/ntag215/
curl https://nfcfyi.com/api/random/
```

Full OpenAPI spec: https://nfcfyi.com/api/openapi.json

## Also Available

| Language | Package | Install |
|----------|---------|---------|
| Python | [nfcfyi](https://pypi.org/project/nfcfyi/) | `pip install nfcfyi` |
| TypeScript | [nfcfyi](https://www.npmjs.com/package/nfcfyi) | `npm install nfcfyi` |
| Go | [nfcfyi-go](https://pkg.go.dev/github.com/fyipedia/nfcfyi-go) | `go get github.com/fyipedia/nfcfyi-go` |
| Rust | [nfcfyi](https://crates.io/crates/nfcfyi) | `cargo add nfcfyi` |
| Ruby | [nfcfyi](https://rubygems.org/gems/nfcfyi) | `gem install nfcfyi` |

## Code FYI Family

| Site | Domain | Focus |
|------|--------|-------|
| BarcodeFYI | [barcodefyi.com](https://barcodefyi.com) | Barcode symbologies and standards |
| QRCodeFYI | [qrcodefyi.com](https://qrcodefyi.com) | QR code types and encoding |
| NFCFYI | [nfcfyi.com](https://nfcfyi.com) | NFC chips and standards |
| BLEFYI | [blefyi.com](https://blefyi.com) | Bluetooth Low Energy profiles |
| RFIDFYI | [rfidfyi.com](https://rfidfyi.com) | RFID tags and frequencies |
| SmartCardFYI | [smartcardfyi.com](https://smartcardfyi.com) | Smart card platforms and standards |

## License

MIT
