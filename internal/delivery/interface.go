package delivery

import "github.com/timickb/transport-sound/internal/controller"

type AuthController interface {
	SignIn(req *controller.AuthRequest) (*controller.AuthResponse, error)
}
type UserController interface {
	Register(req *controller.RegisterRequest) (*controller.RegisterResponse, error)
	ChangeLogin(req *controller.ChangeLoginRequest) error
	ChangeEmail(req *controller.ChangeEmailRequest) error
	ChangePassword(req *controller.ChangePasswordRequest) error
}
