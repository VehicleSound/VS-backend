package usecase

import (
	"crypto/sha256"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type AuthUseCase struct {
	repo Repository
}

func NewAuthUseCase(r Repository) *AuthUseCase {
	return &AuthUseCase{repo: r}
}

func (u *AuthUseCase) SignIn(email, password, secret string) (string, error) {
	user, err := u.repo.GetUserByEmail(email)
	if err != nil {
		return "", fmt.Errorf("err sign in: %w", err)
	}

	pHash := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
	if user.PasswordHash != pHash {
		return "", fmt.Errorf("err sign in: wrong password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"login":      user.Login,
		"email":      user.Email,
		"authorized": true,
		"exp":        time.Now().Add(10 * time.Minute),
	})

	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("err sign in: %w", err)
	}

	return tokenStr, nil
}
