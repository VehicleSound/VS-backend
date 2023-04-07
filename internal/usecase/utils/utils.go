package utils

import (
	"github.com/timickb/transport-sound/internal/usecase"
	"regexp"
)

func ValidateLogin(login string) bool {
	return len(login) >= usecase.MinLoginLen && len(login) <= usecase.MaxLoginLen
}

func ValidatePassword(password string) bool {
	if len(password) < usecase.MinPasswordLen || len(password) > usecase.MaxPasswordLen {
		return false
	}

	digitRegex := regexp.MustCompile("[0-9]")
	upperCaseRegex := regexp.MustCompile("[A-Z]")
	lowerCaseRegex := regexp.MustCompile("[a-z]")

	if !digitRegex.MatchString(password) {
		return false
	}

	if !upperCaseRegex.MatchString(password) {
		return false
	}

	if !lowerCaseRegex.MatchString(password) {
		return false
	}

	return true
}

func ValidateEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(email)

}

func ValidateTag(title string) bool {
	return len(title) > 0
}
