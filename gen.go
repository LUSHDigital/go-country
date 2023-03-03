//go:build ignore

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"log"
	"os"

	"github.com/LUSHDigital/go-country"
)

type Country struct {
	Name   string `json:"name"`
	Alpha2 string `json:"alpha_2"`
	Alpha3 string `json:"alpha_3"`
}

func main() {
	f, err := os.Open("./data/iso_3166-1.json")
	if err != nil {
		log.Fatal(err)
	}

	countries, err := readCountries(f)
	if err != nil {
		log.Fatal(err)
	}

	var countriesBuf, alpha2Buf, alpha3Buf bytes.Buffer

	for i, v := range countries {
		fmt.Fprintf(
			&countriesBuf, `
			{
				Name: "%s",
				Alpha2: "%s",
				Alpha3: "%s",
			},`,
			v.Name,
			v.Alpha2,
			v.Alpha3,
		)

		fmt.Fprintf(&alpha2Buf, `
			case "%s":
				ret = countries[%d]`,
			v.Alpha2,
			i,
		)

		fmt.Fprintf(&alpha3Buf, `
			case "%s":
				ret = countries[%d]`,
			v.Alpha3,
			i,
		)
	}

	var buf bytes.Buffer

	fmt.Fprintln(&buf, "package country")
	fmt.Fprintf(
		&buf,
		`var countries = [%d]Country{%s
		}

		// ByAlpha2 looks up a country by its ISO 3166-1 alpha-2 code.
		// Callers should check the boolean before using the returned country.
		func ByAlpha2(code string) (*Country, bool) {
			var ret Country

			// Prevent escape to the heap by using switch over map.
			switch code {%s
			default:
				return nil, false
			}

			return &ret, true
		}

		// ByAlpha3 looks up a country by its ISO 3166-1 alpha-3 code.
		// Callers should check the boolean before using the returned country.
		func ByAlpha3(code string) (*Country, bool) {
			var ret Country

			// Prevent escape to the heap by using switch over map.
			switch code {%s
			default:
				return nil, false
			}

			return &ret, true
		}`,
		len(countries),
		countriesBuf.Bytes(),
		alpha2Buf.Bytes(),
		alpha3Buf.Bytes(),
	)

	res, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatalf("go fmt: %v", err)
	}

	if err := os.WriteFile("./country_gen.go", res, 0644); err != nil {
		log.Fatal(err)
	}
}

func readCountries(f *os.File) ([]country.Country, error) {
	defer f.Close()

	data := make(map[string][]Country)

	if err := json.NewDecoder(f).Decode(&data); err != nil {
		return nil, fmt.Errorf("json decode: %v", err)
	}

	countries := make([]country.Country, 0, len(data["3166-1"]))

	for _, v := range data["3166-1"] {
		countries = append(countries, country.Country{
			Name:   v.Name,
			Alpha2: v.Alpha2,
			Alpha3: v.Alpha3,
		})
	}

	return countries, nil
}
