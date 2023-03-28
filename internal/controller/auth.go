package controller

import (
	"context"
	dto2 "github.com/timickb/transport-sound/internal/controller/dto"
	"github.com/timickb/transport-sound/internal/domain"
)

type AuthUseCase interface {
	SignIn(email, password, secret string) (string, error)
	GetUserByToken(tokenRaw, secret string) (*domain.User, error)
}

type AuthController struct {
	u      AuthUseCase
	secret string
}

func NewAuth(u AuthUseCase, secret string) *AuthController {
	return &AuthController{u: u, secret: secret}
}

func (c *AuthController) SignIn(ctx context.Context, req *dto2.AuthRequest) (*dto2.AuthResponse, error) {
	token, err := c.u.SignIn(req.Email, req.Password, c.secret)
	if err != nil {
		return nil, err
	}

	return &dto2.AuthResponse{Token: token}, nil
}

func (c *AuthController) GetUserByToken(ctx context.Context, token string) (*dto2.TokenResponse, error) {
	user, err := c.u.GetUserByToken(token, c.secret)
	if err != nil {
		return nil, err
	}

	return &dto2.TokenResponse{
		Id:        user.Id,
		Login:     user.Login,
		Email:     user.Email,
		Confirmed: user.Confirmed,
		Active:    user.Active,
	}, nil
}
