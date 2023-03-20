package auth

import (
	"crypto/sha256"
	"fmt"
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

	user := &domain.User{
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
	if _, err := r.CreateUser(user.Login, user.Email, user.PasswordHash); err != nil {
		t.Fatal(err)
	}

	authService := NewAuthUseCase(r, logrus.New())

	// sign in with right credentials
	_, err := authService.SignIn(user.Email, pwd, secret)
	if err != nil {
		t.Fatal(err)
	}

	// sign in with wrong password
	_, err = authService.SignIn(user.Email, "", secret)
	if err == nil {
		t.Fatal("expected wrong password error")
	}

	// sign in with wrong email
	_, err = authService.SignIn("bread", pwd, secret)
	if err == nil {
		t.Fatal("expected user not found error")
	}
}
