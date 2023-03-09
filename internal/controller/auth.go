package controller

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

func (c *AuthController) SignIn(req *AuthRequest) (*AuthResponse, error) {
	token, err := c.u.SignIn(req.Email, req.Password, c.secret)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{Token: token}, nil
}
