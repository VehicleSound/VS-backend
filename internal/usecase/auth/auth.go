package auth

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mitchellh/mapstructure"
	"github.com/timickb/transport-sound/internal/domain"
	"github.com/timickb/transport-sound/internal/interfaces"
	"github.com/timickb/transport-sound/internal/usecase"
	"time"
)

type UseCase struct {
	repo usecase.Repository
	log  interfaces.Logger
}

func New(r usecase.Repository, log interfaces.Logger) *UseCase {
	return &UseCase{repo: r, log: log}
}

func (u *UseCase) SignIn(email, password, secret string) (string, error) {
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

func (u *UseCase) GetUserByToken(tokenRaw, secret string) (*domain.User, error) {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenRaw, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	user := &domain.User{}
	if err = mapstructure.Decode(claims, user); err != nil {
		return nil, fmt.Errorf("failed to decode claims: %w", err)
	}

	return user, nil
}
