package utils

import "testing"

func TestValidateEmail(t *testing.T) {
	correct := []string{
		"somebody@somewhere.com",
		"sub.domain@site.net",
		"sub.domain@subdomain.site.net",
	}
	incorrect := []string{
		"somebody@somewhere",
		"some_text",
		"domain.com",
		"@",
		"text@",
		"@text",
		"text.text@",
		"@text.text",
		"a@b.c",
	}

	for _, email := range correct {
		if !ValidateEmail(email) {
			t.Fatal("expected true on case", email)
		}
	}
	for _, email := range incorrect {
		if ValidateEmail(email) {
			t.Fatal("expected false on case", email)
		}
	}
}
