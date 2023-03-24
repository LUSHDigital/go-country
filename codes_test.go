package country

import "testing"

func TestAlpha2(t *testing.T) {
	codes := Alpha2()

	var found bool

	sample := "GB"

	for _, v := range codes {
		if v == sample {
			found = true

			break
		}
	}

	if !found {
		t.Errorf("alpha2 country code %q could not be found", sample)
	}
}

func TestAlpha3(t *testing.T) {
	codes := Alpha3()

	var found bool

	sample := "GBR"

	for _, v := range codes {
		if v == sample {
			found = true

			break
		}
	}

	if !found {
		t.Errorf("alpha3 country code %q could not be found", sample)
	}
}
