# country

Provides minimal country data lookup via ISO 3166-1 alpha-2/alpha-3 codes.

## Installation

```bash
go get github.com/LUSHDigital/go-country
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/LUSHDigital/go-country"
)

func main() {
	if gb, ok := country.ByAlpha2("GB"); ok {
		fmt.Println(gb.Alpha2) // "GB"
		fmt.Println(gb.Alpha3) // "GBR"
		fmt.Println(gb.Name) // "United Kingdom"
	}

	if fr, ok := country.ByAlpha3("FRA"); ok {
		fmt.Println(fr.Alpha2) // "FR"
		fmt.Println(fr.Alpha3) // "FRA"
		fmt.Println(fr.Name) // "France"
	}
}
```