package desktop

import (
	"testing"
)

var locales = [...]struct{userloc string; locale Locale}{
	{`lang_COUNTRY.ENCODING@MODIFIER`, Locale{"lang", "COUNTRY", "ENCODING", "MODIFIER"}},
	{`lang_COUNTRY.ENCODING`, Locale{"lang", "COUNTRY", "ENCODING", ""}},
	{`lang_COUNTRY@MODIFIER`, Locale{"lang", "COUNTRY", "", "MODIFIER"}},
	{`lang.ENCODING@MODIFIER`, Locale{"lang", "", "ENCODING", "MODIFIER"}},
	{`lang.ENCODING`, Locale{"lang", "", "ENCODING", ""}},
	{`lang@MODIFIER`, Locale{"lang", "", "", "MODIFIER"}},
	{`en_AU`, Locale{"en", "AU", "", ""}},
	{`fr`, Locale{"fr", "", "", ""}},
	{`foo_BAR..BAZ@FOOBAR`, Locale{}},
	{`foo_BAR.baz@foobar`, Locale{}},
	{`fooBAR.baz@foobar`, Locale{}},
}

func TestParseLocale(t *testing.T) {
	for i, test := range locales {
		parsed, err := ParseLocale(test.userloc)
		if err != nil && i < 8 {
			t.Fatalf(`ParseLocale("%s") returned unexpected error: %v`, test.userloc, err)
		} else if err == nil && i > 7 {
			t.Fatalf(`ParseLocale("%s") did not return expected error`, test.userloc)
		}
		if parsed != test.locale {
			t.Fatalf(`ParseLocale("%s") equals %s not %s`, test.userloc, parsed, test.locale)
		}
	}
}
