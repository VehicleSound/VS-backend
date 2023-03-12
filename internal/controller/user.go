package controller

import (
	"fmt"
	"github.com/timickb/transport-sound/internal/controller/dto"
	"github.com/timickb/transport-sound/internal/domain"
)

type UserUseCase interface {
	GetUserByLoginOrEmailOrId(cred string) (*domain.User, error)
	CreateUser(login, email, password string) (string, error)
	ChangePassword(id, oPwd, nPwd string) error
	ChangeLogin(id, nLogin string) error
	ChangeEmail(id, nEmail string) error
	Deactivate(id string) error
	AddToFav(userId, soundId string) error
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

func (c *UserController) GetUserById(id string) (*dto.GetUserResponse, error) {
	user, err := c.u.GetUserByLoginOrEmailOrId(id)
	if err != nil {
		return nil, err
	}

	return c.mapUser(user), nil
}

func (c *UserController) GetUser(req *dto.GetUserRequest) (*dto.GetUserResponse, error) {
	if req.Login != "" {
		user, err := c.u.GetUserByLoginOrEmailOrId(req.Login)
		if err != nil {
			return nil, err
		}

		return c.mapUser(user), nil
	}

	if req.Email != "" {
		user, err := c.u.GetUserByLoginOrEmailOrId(req.Email)
		if err != nil {
			return nil, err
		}

		return c.mapUser(user), nil
	}

	return nil, fmt.Errorf("err get user: wrong credentials")
}

func (c *UserController) AddToFav(req *dto.AddToFavRequest) error {
	if err := c.u.AddToFav(req.UserId, req.SoundId); err != nil {
		return err
	}

	return nil
}

func (c *UserController) mapUser(user *domain.User) *dto.GetUserResponse {
	return &dto.GetUserResponse{
		Id:        user.Id,
		Login:     user.Login,
		Email:     user.Email,
		Active:    user.Active,
		Confirmed: user.Confirmed,
	}
}
