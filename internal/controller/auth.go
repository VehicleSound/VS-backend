package controller

import "github.com/timickb/transport-sound/internal/controller/dto"

type AuthUseCase interface {
	SignIn(email, password, secret string) (string, error)
}

type AuthController struct {
	u      AuthUseCase
	secret string
}

func NewAuthController(u AuthUseCase, secret string) *AuthController {
	return &AuthController{u: u}
}

func (c *AuthController) SignIn(req *dto.AuthRequest) (*dto.AuthResponse, error) {
	token, err := c.u.SignIn(req.Email, req.Password, c.secret)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{Token: token}, nil
}
