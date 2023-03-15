package repository

type UserEditPayload struct {
	Login     string
	Email     string
	Password  string
	Activated bool
}
