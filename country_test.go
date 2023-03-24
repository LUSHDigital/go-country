package country

import (
	"reflect"
	"testing"
)

func TestByAlpha2(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		c, b := ByAlpha2("GB")

		if b != true {
			t.Errorf("b = %v, want %v", b, true)
		}

		if c.Alpha2 != "GB" {
			t.Errorf("c.Alpha2 = %v, want %v", c.Alpha2, "GB")
		}

		if c.Alpha3 != "GBR" {
			t.Errorf("c.Alpha3 = %v, want %v", c.Alpha3, "GBR")
		}

		if !reflect.DeepEqual(c.Locales, []string{"en-GB", "cy-GB", "gd"}) {
			t.Errorf("c.Locales = %q, want %q", c.Locales, []string{"en-GB", "cy-GB", "gd"})
		}
	})

	t.Run("invalid", func(t *testing.T) {
		_, b := ByAlpha2("XX")

		if b != false {
			t.Errorf("b = %v, want %v", b, false)
		}
	})
}

func TestByAlpha3(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		c, b := ByAlpha3("GBR")

		if b != true {
			t.Errorf("b = %v, want %v", b, true)
		}

		if c.Alpha2 != "GB" {
			t.Errorf("c.Alpha2 = %v, want %v", c.Alpha2, "GB")
		}

		if c.Alpha3 != "GBR" {
			t.Errorf("c.Alpha3 = %v, want %v", c.Alpha3, "GBR")
		}

		if !reflect.DeepEqual(c.Locales, []string{"en-GB", "cy-GB", "gd"}) {
			t.Errorf("c.Locales = %q, want %q", c.Locales, []string{"en-GB", "cy-GB", "gd"})
		}
	})

	t.Run("invalid", func(t *testing.T) {
		_, b := ByAlpha2("FOO")

		if b != false {
			t.Errorf("b = %v, want %v", b, false)
		}
	})
}

func BenchmarkByAlpha2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ByAlpha2("GB")
		ByAlpha2("DE")
		ByAlpha2("US")
	}
}

func BenchmarkByAlpha3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ByAlpha3("GBR")
		ByAlpha3("DEU")
		ByAlpha3("USA")
	}
}
