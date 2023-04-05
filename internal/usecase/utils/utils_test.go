package utils

import (
	"github.com/timickb/transport-sound/internal/usecase"
	"strings"
	"testing"
)

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

func TestValidatePassword(t *testing.T) {
	correct := []string{
		"8Ab8Ab8Ab",
		"A8b_A8b_A8b",
	}
	incorrect := []string{
		"",
		strings.Repeat("Abc", usecase.MaxPasswordLen/3),
		strings.Repeat("ABC", usecase.MaxPasswordLen/3),
		strings.Repeat("abc", usecase.MaxPasswordLen/3),
		strings.Repeat("123", usecase.MaxPasswordLen/3),
		strings.Repeat("A8b", usecase.MaxPasswordLen/3+1),
		strings.Repeat("A8b", usecase.MaxPasswordLen/3)[:usecase.MinPasswordLen-1],
	}

	for _, pwd := range correct {
		if !ValidatePassword(pwd) {
			t.Fatal("expected true on case", pwd)
		}
	}

	for _, pwd := range incorrect {
		if ValidatePassword(pwd) {
			t.Fatal("expected false on case", pwd)
		}
	}
}
