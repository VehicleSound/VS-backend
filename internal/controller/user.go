package controller

import "github.com/timickb/transport-sound/internal/controller/dto"

type UserUseCase interface {
	CreateUser(login, email, password string) (string, error)
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

func (c *UserController) Register(req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	userId, err := c.u.CreateUser(req.Login, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &dto.RegisterResponse{UserId: userId}, nil
}

func (c *UserController) ChangeLogin(req *dto.ChangeLoginRequest) error {
	err := c.u.ChangeLogin(req.UserId, req.Login)
	if err != nil {
		return err
	}
	return nil
}

func (c *UserController) ChangeEmail(req *dto.ChangeEmailRequest) error {
	err := c.u.ChangeEmail(req.UserId, req.Email)
	if err != nil {
		return err
	}
	return nil
}

func (c *UserController) ChangePassword(req *dto.ChangePasswordRequest) error {
	err := c.u.ChangePassword(req.UserId, req.OldPassword, req.NewPassword)
	if err != nil {
		return err
	}
	return nil
}
