package controller

import "github.com/timickb/transport-sound/internal/domain"

type UserUseCase interface {
	CreateUser(login, email, password string) (*domain.User, error)
	ChangePassword(id, oPwd, nPwd string) error
	ChangeLogin(id, nLogin string) error
	ChangeEmail(id, nEmail string) error
	Deactivate(id string) error
}

type UserController struct {
	u UserUseCase
}

func NewUserController(u UserUseCase) *UserController {
	return &UserController{u: u}
}

func (c *UserController) Register(req *RegisterRequest) (*RegisterResponse, error) {
	user, err := c.u.CreateUser(req.Login, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &RegisterResponse{UserId: user.Id}, nil
}

func (c *UserController) ChangeLogin(req *ChangeLoginRequest) error {
	err := c.u.ChangeLogin(req.UserId, req.Login)
	if err != nil {
		return err
	}
	return nil
}

func (c *UserController) ChangeEmail(req *ChangeEmailRequest) error {
	err := c.u.ChangeEmail(req.UserId, req.Email)
	if err != nil {
		return err
	}
	return nil
}

func (c *UserController) ChangePassword(req *ChangePasswordRequest) error {
	err := c.u.ChangePassword(req.UserId, req.OldPassword, req.NewPassword)
	if err != nil {
		return err
	}
	return nil
}
