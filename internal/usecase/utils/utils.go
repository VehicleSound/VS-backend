package utils

import "regexp"

func ValidateLogin(login string) bool {
	return len(login) >= 3
}

func ValidatePassword(password string) bool {
	return len(password) >= 3
}

func ValidateEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(email)

}

func ValidateTag(title string) bool {
	return len(title) > 0
}
