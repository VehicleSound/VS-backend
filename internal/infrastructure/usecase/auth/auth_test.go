package auth

import (
	"crypto/sha256"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/timickb/transport-sound/internal/infrastructure/domain"
	"github.com/timickb/transport-sound/internal/infrastructure/repository/memory"
	"testing"
	"time"
)

func TestSignIn(t *testing.T) {
	pwd := "password"
	pwdHash := fmt.Sprintf("%x", sha256.Sum256([]byte(pwd)))

	user := domain.User{
		Id:           uuid.NewString(),
		Login:        "test_user",
		Email:        "test@example.com",
		PasswordHash: pwdHash,
		Confirmed:    false,
		Active:       false,
		DateCreated:  time.Now(),
	}

	secret := "jwt_secret"

	r := memory.NewRepository()
	if err := r.CreateUser(user); err != nil {
		t.Fatal(err)
	}

	authService := New(r, logrus.New())

	// sign in with right credentials
	_, err := authService.SignIn(user.Email, pwd, secret)
	if err != nil {
		t.Fatal(err)
	}

	// sign in with wrong password
	_, err = authService.SignIn(user.Email, "", secret)
	if err == nil {
		t.Fatal("expected wrong password srverrors")
	}

	// sign in with wrong email
	_, err = authService.SignIn("bread", pwd, secret)
	if err == nil {
		t.Fatal("expected user not found srverrors")
	}
}

func TestValidateToken(t *testing.T) {
	secret := "secret"
	r := memory.NewRepository()
	authService := New(r, logrus.New())

	user := &domain.User{
		Id:           "12345",
		Login:        "login",
		Email:        "email",
		PasswordHash: "pwd_hash",
		Confirmed:    false,
		Active:       false,
		DateCreated:  time.Now(),
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
		t.Fatal(err)
	}

	decoded, err := authService.GetUserByToken(tokenStr, secret)
	if err != nil {
		t.Fatal(err)
	}

	if decoded.Email != user.Email {
		t.Error("email field corrupted")
	}
	if decoded.Id != user.Id {
		t.Errorf("id field corrupted")
	}
	if decoded.Login != user.Login {
		t.Errorf("login field corrupted")
	}
	if decoded.Confirmed != user.Confirmed {
		t.Errorf("confirmed field corrupted")
	}
	if decoded.Active != user.Active {
		t.Errorf("user field corrupted")
	}
}
