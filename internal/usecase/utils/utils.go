package utils

func ValidateLogin(login string) bool {
	return len(login) >= 3
}

func ValidatePassword(password string) bool {
	return len(password) >= 3
}

func ValidateEmail(email string) bool {
	return true
}

func ValidateTag(title string) bool {
	return len(title) > 0
}
