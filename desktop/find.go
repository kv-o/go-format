package desktop

// Locale represents a valid locale for a desktop entry.
type Locale struct {
	Lang     string
	Country  string
	Encoding string
	Modifier string
}

// Get attempts to get the value associated with the specified key in the
// specified group, for locale. Returns an error if the key is not found.
func Get(group any, key string, locale Locale) (any, error) {
	return nil, nil
}

// Groups returns a slice of key-value pair groups from the desktop entry at v.
func Groups(v any) []any {
	return nil
}

// ParseLocale returns a Locale for string s, which must be of the form
// lang_COUNTRY.ENCODING@MODIFIER
func ParseLocale(s string) (Locale, error) {
	return Locale{}, nil
}
