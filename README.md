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

## FYIPedia Developer Tools

Part of the [FYIPedia](https://fyipedia.com) open-source developer tools ecosystem.

| Package | PyPI | npm | Go | Description |
|---------|------|-----|----|-------------|
| airlinefyi | [PyPI](https://pypi.org/project/airlinefyi/) | [npm](https://www.npmjs.com/package/airlinefyi) | [Go](https://pkg.go.dev/github.com/fyipedia/airlinefyi-go) | Airlines, fleets, alliances — [airlinefyi.com](https://airlinefyi.com) |
| airportfyi | [PyPI](https://pypi.org/project/airportfyi/) | [npm](https://www.npmjs.com/package/airportfyi) | [Go](https://pkg.go.dev/github.com/fyipedia/airportfyi-go) | 4,500+ airports, IATA/ICAO — [airportfyi.com](https://airportfyi.com) |
| alloyfyi | [PyPI](https://pypi.org/project/alloyfyi/) | [npm](https://www.npmjs.com/package/alloyfyi) | [Go](https://pkg.go.dev/github.com/fyipedia/alloyfyi-go) | Metal alloys, compositions — [alloyfyi.com](https://alloyfyi.com) |
| anatomyfyi | [PyPI](https://pypi.org/project/anatomyfyi/) | [npm](https://www.npmjs.com/package/anatomyfyi) | [Go](https://pkg.go.dev/github.com/fyipedia/anatomyfyi-go) | 14,692 anatomical structures — [anatomyfyi.com](https://anatomyfyi.com) |
| barcodefyi | [PyPI](https://pypi.org/project/barcodefyi/) | [npm](https://www.npmjs.com/package/barcodefyi) | [Go](https://pkg.go.dev/github.com/fyipedia/barcodefyi-go) | Barcode formats, EAN, UPC, ISBN — [barcodefyi.com](https://barcodefyi.com) |
| beerfyi | [PyPI](https://pypi.org/project/beerfyi/) | [npm](https://www.npmjs.com/package/beerfyi) | [Go](https://pkg.go.dev/github.com/fyipedia/beerfyi-go) | 112 beer styles, hops, malts, yeast — [beerfyi.com](https://beerfyi.com) |
| birdfyi | [PyPI](https://pypi.org/project/birdfyi/) | [npm](https://www.npmjs.com/package/birdfyi) | [Go](https://pkg.go.dev/github.com/fyipedia/birdfyi-go) | 11,251 birds, orders, conservation — [birdfyi.com](https://birdfyi.com) |
| blefyi | [PyPI](https://pypi.org/project/blefyi/) | [npm](https://www.npmjs.com/package/blefyi) | [Go](https://pkg.go.dev/github.com/fyipedia/blefyi-go) | Bluetooth Low Energy, GATT, beacons — [blefyi.com](https://blefyi.com) |
| boardgamefyi | [PyPI](https://pypi.org/project/boardgamefyi/) | [npm](https://www.npmjs.com/package/boardgamefyi) | [Go](https://pkg.go.dev/github.com/fyipedia/boardgamefyi-go) | Board games, rules, reviews — [boardgamefyi.com](https://boardgamefyi.com) |
| brewfyi | [PyPI](https://pypi.org/project/brewfyi/) | [npm](https://www.npmjs.com/package/brewfyi) | [Go](https://pkg.go.dev/github.com/fyipedia/brewfyi-go) | 72 coffee varieties, roasting, brewing — [brewfyi.com](https://brewfyi.com) |
| cablefyi | [PyPI](https://pypi.org/project/cablefyi/) | [npm](https://www.npmjs.com/package/cablefyi) | [Go](https://pkg.go.dev/github.com/fyipedia/cablefyi-go) | Submarine cables, connectors — [cablefyi.com](https://cablefyi.com) |
| calcfyi | [PyPI](https://pypi.org/project/calcfyi/) | [npm](https://www.npmjs.com/package/calcfyi) | [Go](https://pkg.go.dev/github.com/fyipedia/calcfyi-go) | 200+ calculators — [calcfyi.com](https://calcfyi.com) |
| chemfyi | [PyPI](https://pypi.org/project/chemfyi/) | [npm](https://www.npmjs.com/package/chemfyi) | [Go](https://pkg.go.dev/github.com/fyipedia/chemfyi-go) | Periodic table, elements, compounds — [chemfyi.com](https://chemfyi.com) |
| cocktailfyi | [PyPI](https://pypi.org/project/cocktailfyi/) | [npm](https://www.npmjs.com/package/cocktailfyi) | [Go](https://pkg.go.dev/github.com/fyipedia/cocktailfyi-go) | 636 cocktail recipes, ABV, calories — [cocktailfyi.com](https://cocktailfyi.com) |
| colorfyi | [PyPI](https://pypi.org/project/colorfyi/) | [npm](https://www.npmjs.com/package/colorfyi) | [Go](https://pkg.go.dev/github.com/fyipedia/colorfyi-go) | Color science, WCAG contrast, 16.7M hex colors — [colorfyi.com](https://colorfyi.com) |
| dinofyi | [PyPI](https://pypi.org/project/dinofyi/) | [npm](https://www.npmjs.com/package/dinofyi) | [Go](https://pkg.go.dev/github.com/fyipedia/dinofyi-go) | 6,142 dinosaurs, paleontology — [dinofyi.com](https://dinofyi.com) |
| distancefyi | [PyPI](https://pypi.org/project/distancefyi/) | [npm](https://www.npmjs.com/package/distancefyi) | [Go](https://pkg.go.dev/github.com/fyipedia/distancefyi-go) | Haversine distance & travel times — [distancefyi.com](https://distancefyi.com) |
| drugfyi | [PyPI](https://pypi.org/project/drugfyi/) | [npm](https://www.npmjs.com/package/drugfyi) | [Go](https://pkg.go.dev/github.com/fyipedia/drugfyi-go) | Drug interactions, pharmacology — [drugfyi.com](https://drugfyi.com) |
| emojifyi | [PyPI](https://pypi.org/project/emojifyi/) | [npm](https://www.npmjs.com/package/emojifyi) | [Go](https://pkg.go.dev/github.com/fyipedia/emojifyi-go) | Emoji metadata & encoding for 3,953 emojis — [emojifyi.com](https://emojifyi.com) |
| fishfyi | [PyPI](https://pypi.org/project/fishfyi/) | [npm](https://www.npmjs.com/package/fishfyi) | [Go](https://pkg.go.dev/github.com/fyipedia/fishfyi-go) | Fish species, marine biology — [fishfyi.com](https://fishfyi.com) |
| fontfyi | [PyPI](https://pypi.org/project/fontfyi/) | [npm](https://www.npmjs.com/package/fontfyi) | [Go](https://pkg.go.dev/github.com/fyipedia/fontfyi-go) | Google Fonts metadata & CSS — [fontfyi.com](https://fontfyi.com) |
| gemfyi | [PyPI](https://pypi.org/project/gemfyi/) | [npm](https://www.npmjs.com/package/gemfyi) | [Go](https://pkg.go.dev/github.com/fyipedia/gemfyi-go) | Gemstones, mineralogy, grading — [gemfyi.com](https://gemfyi.com) |
| holidayfyi | [PyPI](https://pypi.org/project/holidayfyi/) | [npm](https://www.npmjs.com/package/holidayfyi) | [Go](https://pkg.go.dev/github.com/fyipedia/holidayfyi-go) | Holiday dates & Easter calculation — [holidayfyi.com](https://holidayfyi.com) |
| ipfyi | [PyPI](https://pypi.org/project/ipfyi/) | [npm](https://www.npmjs.com/package/ipfyi) | [Go](https://pkg.go.dev/github.com/fyipedia/ipfyi-go) | IP geolocation, ASN lookup — [ipfyi.com](https://ipfyi.com) |
| mineralfyi | [PyPI](https://pypi.org/project/mineralfyi/) | [npm](https://www.npmjs.com/package/mineralfyi) | [Go](https://pkg.go.dev/github.com/fyipedia/mineralfyi-go) | 6,215 minerals, crystal systems — [mineralfyi.com](https://mineralfyi.com) |
| mountainfyi | [PyPI](https://pypi.org/project/mountainfyi/) | [npm](https://www.npmjs.com/package/mountainfyi) | [Go](https://pkg.go.dev/github.com/fyipedia/mountainfyi-go) | Mountains, peaks, elevation — [mountainfyi.com](https://mountainfyi.com) |
| namefyi | [PyPI](https://pypi.org/project/namefyi/) | [npm](https://www.npmjs.com/package/namefyi) | [Go](https://pkg.go.dev/github.com/fyipedia/namefyi-go) | Korean romanization & Five Elements — [namefyi.com](https://namefyi.com) |
| **nfcfyi** | [PyPI](https://pypi.org/project/nfcfyi/) | [npm](https://www.npmjs.com/package/nfcfyi) | [Go](https://pkg.go.dev/github.com/fyipedia/nfcfyi-go) | NFC chips, tag types, NDEF — [nfcfyi.com](https://nfcfyi.com) |
| nihonshufyi | [PyPI](https://pypi.org/project/nihonshufyi/) | [npm](https://www.npmjs.com/package/nihonshufyi) | [Go](https://pkg.go.dev/github.com/fyipedia/nihonshufyi-go) | 80 sake, rice varieties, breweries — [nihonshufyi.com](https://nihonshufyi.com) |
| nutrifyi | [PyPI](https://pypi.org/project/nutrifyi/) | [npm](https://www.npmjs.com/package/nutrifyi) | [Go](https://pkg.go.dev/github.com/fyipedia/nutrifyi-go) | Nutrition data, food composition — [nutrifyi.com](https://nutrifyi.com) |
| pillfyi | [PyPI](https://pypi.org/project/pillfyi/) | [npm](https://www.npmjs.com/package/pillfyi) | [Go](https://pkg.go.dev/github.com/fyipedia/pillfyi-go) | Pill identification, FDA database — [pillfyi.com](https://pillfyi.com) |
| planefyi | [PyPI](https://pypi.org/project/planefyi/) | [npm](https://www.npmjs.com/package/planefyi) | [Go](https://pkg.go.dev/github.com/fyipedia/planefyi-go) | Aircraft models, specifications — [planefyi.com](https://planefyi.com) |
| plantfyi | [PyPI](https://pypi.org/project/plantfyi/) | [npm](https://www.npmjs.com/package/plantfyi) | [Go](https://pkg.go.dev/github.com/fyipedia/plantfyi-go) | 379,774 plants, taxonomy — [plantfyi.com](https://plantfyi.com) |
| qrcodefyi | [PyPI](https://pypi.org/project/qrcodefyi/) | [npm](https://www.npmjs.com/package/qrcodefyi) | [Go](https://pkg.go.dev/github.com/fyipedia/qrcodefyi-go) | QR code types, versions, encoding — [qrcodefyi.com](https://qrcodefyi.com) |
| quakefyi | [PyPI](https://pypi.org/project/quakefyi/) | [npm](https://www.npmjs.com/package/quakefyi) | [Go](https://pkg.go.dev/github.com/fyipedia/quakefyi-go) | Earthquakes, seismic data — [quakefyi.com](https://quakefyi.com) |
| rfidfyi | [PyPI](https://pypi.org/project/rfidfyi/) | [npm](https://www.npmjs.com/package/rfidfyi) | [Go](https://pkg.go.dev/github.com/fyipedia/rfidfyi-go) | RFID tags, frequency bands, standards — [rfidfyi.com](https://rfidfyi.com) |
| salaryfyi | [PyPI](https://pypi.org/project/salaryfyi/) | [npm](https://www.npmjs.com/package/salaryfyi) | [Go](https://pkg.go.dev/github.com/fyipedia/salaryfyi-go) | Salary comparison, tax calculators — [salaryfyi.com](https://salaryfyi.com) |
| smartcardfyi | [PyPI](https://pypi.org/project/smartcardfyi/) | [npm](https://www.npmjs.com/package/smartcardfyi) | [Go](https://pkg.go.dev/github.com/fyipedia/smartcardfyi-go) | Smart cards, EMV, APDU, Java Card — [smartcardfyi.com](https://smartcardfyi.com) |
| speciesfyi | [PyPI](https://pypi.org/project/speciesfyi/) | [npm](https://www.npmjs.com/package/speciesfyi) | [Go](https://pkg.go.dev/github.com/fyipedia/speciesfyi-go) | Species taxonomy, classification — [speciesfyi.com](https://speciesfyi.com) |
| starfyi | [PyPI](https://pypi.org/project/starfyi/) | [npm](https://www.npmjs.com/package/starfyi) | [Go](https://pkg.go.dev/github.com/fyipedia/starfyi-go) | Stars, constellations, exoplanets — [starfyi.com](https://starfyi.com) |
| statuscodefyi | [PyPI](https://pypi.org/project/statuscodefyi/) | [npm](https://www.npmjs.com/package/statuscodefyi) | [Go](https://pkg.go.dev/github.com/fyipedia/statuscodefyi-go) | HTTP status codes, protocols — [statuscodefyi.com](https://statuscodefyi.com) |
| symbolfyi | [PyPI](https://pypi.org/project/symbolfyi/) | [npm](https://www.npmjs.com/package/symbolfyi) | [Go](https://pkg.go.dev/github.com/fyipedia/symbolfyi-go) | Symbol encoding in 11 formats — [symbolfyi.com](https://symbolfyi.com) |
| teafyi | [PyPI](https://pypi.org/project/teafyi/) | [npm](https://www.npmjs.com/package/teafyi) | [Go](https://pkg.go.dev/github.com/fyipedia/teafyi-go) | 60 tea varieties, teaware, brewing — [teafyi.com](https://teafyi.com) |
| timefyi | [PyPI](https://pypi.org/project/timefyi/) | [npm](https://www.npmjs.com/package/timefyi) | [Go](https://pkg.go.dev/github.com/fyipedia/timefyi-go) | Timezone ops & business hours — [timefyi.com](https://timefyi.com) |
| tldfyi | [PyPI](https://pypi.org/project/tldfyi/) | [npm](https://www.npmjs.com/package/tldfyi) | [Go](https://pkg.go.dev/github.com/fyipedia/tldfyi-go) | TLD registry, domain extensions — [tldfyi.com](https://tldfyi.com) |
| trainfyi | [PyPI](https://pypi.org/project/trainfyi/) | [npm](https://www.npmjs.com/package/trainfyi) | [Go](https://pkg.go.dev/github.com/fyipedia/trainfyi-go) | Railway stations, routes, networks — [trainfyi.com](https://trainfyi.com) |
| unicodefyi | [PyPI](https://pypi.org/project/unicodefyi/) | [npm](https://www.npmjs.com/package/unicodefyi) | [Go](https://pkg.go.dev/github.com/fyipedia/unicodefyi-go) | Unicode lookup with 17 encodings — [unicodefyi.com](https://unicodefyi.com) |
| unitfyi | [PyPI](https://pypi.org/project/unitfyi/) | [npm](https://www.npmjs.com/package/unitfyi) | [Go](https://pkg.go.dev/github.com/fyipedia/unitfyi-go) | Unit conversion, 220 units — [unitfyi.com](https://unitfyi.com) |
| univfyi | [PyPI](https://pypi.org/project/univfyi/) | [npm](https://www.npmjs.com/package/univfyi) | [Go](https://pkg.go.dev/github.com/fyipedia/univfyi-go) | University rankings, programs — [univfyi.com](https://univfyi.com) |
| vinofyi | [PyPI](https://pypi.org/project/vinofyi/) | [npm](https://www.npmjs.com/package/vinofyi) | [Go](https://pkg.go.dev/github.com/fyipedia/vinofyi-go) | Wines, grapes, regions, wineries — [vinofyi.com](https://vinofyi.com) |
| whiskeyfyi | [PyPI](https://pypi.org/project/whiskeyfyi/) | [npm](https://www.npmjs.com/package/whiskeyfyi) | [Go](https://pkg.go.dev/github.com/fyipedia/whiskeyfyi-go) | 80 whiskey expressions, distilleries — [whiskeyfyi.com](https://whiskeyfyi.com) |
| zipfyi | [PyPI](https://pypi.org/project/zipfyi/) | [npm](https://www.npmjs.com/package/zipfyi) | [Go](https://pkg.go.dev/github.com/fyipedia/zipfyi-go) | ZIP/postal codes, geocoding — [zipfyi.com](https://zipfyi.com) |

## Links

- **Site**: [nfcfyi.com](https://nfcfyi.com)
- **API Docs**: [nfcfyi.com/api/v1/](https://nfcfyi.com/api/v1/)
- **OpenAPI**: [nfcfyi.com/api/v1/schema/](https://nfcfyi.com/api/v1/schema/)

## License

MIT
