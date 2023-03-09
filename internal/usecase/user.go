package usecase

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/timickb/transport-sound/internal/repository"
)

type UserUseCase struct {
	r Repository
}

func NewUserUseCase(r Repository) *UserUseCase {
	return &UserUseCase{r: r}
}

func (u *UserUseCase) CreateUser(login, email, password string) (string, error) {
	if !validateLogin(login) {
		return "", errors.New("err create user: invalid login")
	}
	if !validatePassword(password) {
		return "", errors.New("err create user: invalid password")
	}
	if !validateEmail(email) {
		return "", errors.New("err create user: invalid email")
	}

	exLogin, errLogin := u.r.GetUserByLogin(login)
	if errLogin == nil && exLogin.Login == login {
		return "", errors.New("err create user: login already exists")
	}

	exEmail, errEmail := u.r.GetUserByEmail(email)
	if errEmail == nil && exEmail.Email == email {
		return "", errors.New("err create user: email already exists")
	}

	pwdHash := sha256.Sum256([]byte(password))
	pwdHashStr := fmt.Sprintf("%x", pwdHash)

	user, err := u.r.CreateUser(login, email, pwdHashStr)
	if err != nil {
		return "", fmt.Errorf("err create user: %w", err)
	}

	return user, nil
}

func (u *UserUseCase) ChangePassword(id, oPwd, nPwd string) error {
	if !validateLogin(nPwd) {
		return errors.New("err password too short")
	}

	user, err := u.r.GetUserById(id)
	if err != nil {
		return fmt.Errorf("err change password: %w", err)
	}

	oPwdHash := fmt.Sprintf("%x", sha256.Sum256([]byte(oPwd)))
	if oPwdHash != user.PasswordHash {
		return errors.New("err change password wrong old password")
	}

	_, err = u.r.EditUser(id, &repository.UserEditPayload{Password: nPwd})
	if err != nil {
		return fmt.Errorf("err change password: %w", err)
	}

	return nil
}

func (u *UserUseCase) ChangeLogin(id, nLogin string) error {
	if !validateLogin(nLogin) {
		return errors.New("err change login: login too short")
	}

	_, err := u.r.GetUserById(id)
	if err != nil {
		return fmt.Errorf("err change login: %w", err)
	}

	_, err = u.r.EditUser(id, &repository.UserEditPayload{Login: nLogin})
	if err != nil {
		return fmt.Errorf("err change login: %w", err)
	}

	return nil
}

func (u *UserUseCase) ChangeEmail(id, nEmail string) error {
	if !validateEmail(nEmail) {
		return errors.New("err change email: login too short")
	}

	_, err := u.r.GetUserById(id)
	if err != nil {
		return fmt.Errorf("err change email: %w", err)
	}

	_, err = u.r.EditUser(id, &repository.UserEditPayload{Email: nEmail})
	if err != nil {
		return fmt.Errorf("err change email: %w", err)
	}

	return nil
}

func (u *UserUseCase) Deactivate(id string) error {
	_, err := u.r.GetUserById(id)
	if err != nil {
		return fmt.Errorf("err deactivate user: %w", err)
	}

	_, err = u.r.EditUser(id, &repository.UserEditPayload{Activated: false})
	if err != nil {
		return fmt.Errorf("err deactivate user: %w", err)
	}

	return nil
}
