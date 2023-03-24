//go:build ignore

package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"log"
	"os"
	"strings"

	"github.com/LUSHDigital/go-country"
)

type Country struct {
	Name   string `json:"name"`
	Alpha2 string `json:"alpha_2"`
	Alpha3 string `json:"alpha_3"`
}

func main() {
	countries, err := readCountries()
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
				Locales: %#v,
			},`,
			v.Name,
			v.Alpha2,
			v.Alpha3,
			v.Locales,
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

			// Prevent mutation of original slice.
			ret.Locales = append([]string(nil), ret.Locales...)

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

			// Prevent mutation of original slice.
			ret.Locales = append([]string(nil), ret.Locales...)

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

func readCountries() ([]country.Country, error) {
	iso3166, err := os.Open("./data/iso_3166-1.json")
	if err != nil {
		return nil, err
	}
	defer iso3166.Close()

	data := make(map[string][]Country)

	if err := json.NewDecoder(iso3166).Decode(&data); err != nil {
		return nil, fmt.Errorf("json decode: %v", err)
	}

	countryInfo, err := os.Open("./data/countryInfo.txt")
	if err != nil {
		return nil, err
	}
	defer countryInfo.Close()

	alpha3Tolocales := make(map[string]string)

	scanner := bufio.NewScanner(countryInfo)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "\t")

		locales := row[len(row)-4]
		if locales == "" {
			continue
		}

		alpha3Tolocales[row[1]] = locales
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	countries := make([]country.Country, 0, len(data["3166-1"]))

	for _, v := range data["3166-1"] {
		var locs []string
		if l, ok := alpha3Tolocales[v.Alpha3]; ok {
			locs = strings.Split(l, ",")
		}

		countries = append(countries, country.Country{
			Name:    v.Name,
			Alpha2:  v.Alpha2,
			Alpha3:  v.Alpha3,
			Locales: locs,
		})
	}

	return countries, nil
}
