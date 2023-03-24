package country

// Alpha2 returns a list of all ISO 3166-1 alpha-2 country codes.
func Alpha2() []string {
	codes := make([]string, 0, len(countries))

	for _, v := range countries {
		codes = append(codes, v.Alpha2)
	}

	return codes
}

// Alpha3 returns a list of all ISO 3166-1 alpha-3 country codes.
func Alpha3() []string {
	codes := make([]string, 0, len(countries))

	for _, v := range countries {
		codes = append(codes, v.Alpha3)
	}

	return codes
}
