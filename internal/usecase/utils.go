package usecase

func validateLogin(login string) bool {
	return len(login) >= 3
}

func validatePassword(password string) bool {
	return len(password) >= 3
}

func validateEmail(email string) bool {
	return true
}

func validateTag(title string) bool {
	return len(title) > 0
}
