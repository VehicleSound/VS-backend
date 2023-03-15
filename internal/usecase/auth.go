package usecase

import (
	"crypto/sha256"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/timickb/transport-sound/internal/domain"
	"github.com/timickb/transport-sound/internal/interfaces"
	"time"
)

type AuthUseCase struct {
	repo Repository
	log  interfaces.Logger
}

func NewAuthUseCase(r Repository, log interfaces.Logger) *AuthUseCase {
	return &AuthUseCase{repo: r, log: log}
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
		"login":     user.Login,
		"email":     user.Email,
		"id":        user.Id,
		"active":    user.Active,
		"confirmed": user.Confirmed,
		"iat":       time.Now().Unix(),
		"exp":       time.Now().Add(time.Hour).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("err sign in: %w", err)
	}

	return tokenStr, nil
}

func (u *AuthUseCase) ValidateToken(tokenRaw, secret string) (*domain.User, error) {
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(tokenRaw, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("token validation failed: %w", err)
	}

	user := &domain.User{
		Id:        fmt.Sprintf("%v", claims["id"]),
		Login:     fmt.Sprintf("%v", claims["login"]),
		Email:     fmt.Sprintf("%v", claims["email"]),
		Confirmed: claims["confirmed"].(bool),
		Active:    claims["active"].(bool),
	}

	return user, nil
}
