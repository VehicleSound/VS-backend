package user

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/timickb/transport-sound/internal/infrastructure/domain"
	"github.com/timickb/transport-sound/internal/infrastructure/repository"
	"github.com/timickb/transport-sound/internal/infrastructure/usecase"
	"github.com/timickb/transport-sound/internal/infrastructure/usecase/utils"
	"github.com/timickb/transport-sound/internal/interfaces"
	"time"
)

type UseCase struct {
	r   usecase.Repository
	log interfaces.Logger
}

func New(r usecase.Repository, log interfaces.Logger) *UseCase {
	return &UseCase{r: r}
}

func (u *UseCase) ValidateRegistration(login, email, password string) error {
	if !utils.ValidatePassword(password) {
		return errors.New("err validate reg: invalid password")
	}
	if !utils.ValidateEmail(email) {
		return errors.New("err validate reg: invalid email")
	}
	if !utils.ValidateLogin(login) {
		return errors.New("err validate reg: invalid login")
	}

	exLogin, errLogin := u.r.GetUserByLogin(login)
	if errLogin == nil && exLogin.Login == login {
		return errors.New("err validate reg: login already exists")
	}

	exEmail, errEmail := u.r.GetUserByEmail(email)
	if errEmail == nil && exEmail.Email == email {
		return errors.New("err validate reg: email already exists")
	}

	return nil
}

func (u *UseCase) CreateUser(login, email, password string) (string, error) {
	userId := uuid.NewString()
	pwdHashStr := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))

	user := domain.User{
		Id:           userId,
		Login:        login,
		Email:        email,
		PasswordHash: pwdHashStr,
		Confirmed:    false,
		Active:       true,
		DateCreated:  time.Now(),
	}

	if err := u.r.CreateUser(user); err != nil {
		return "", fmt.Errorf("err create user: %w", err)
	}

	return userId, nil
}

func (u *UseCase) ChangePassword(id, oPwd, nPwd string) error {
	if !utils.ValidateLogin(nPwd) {
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

func (u *UseCase) ChangeLogin(id, nLogin string) error {
	if !utils.ValidateLogin(nLogin) {
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

func (u *UseCase) ChangeEmail(id, nEmail string) error {
	if !utils.ValidateEmail(nEmail) {
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

func (u *UseCase) AddToFav(userId, soundId string) error {
	if err := u.r.AddFavourite(userId, soundId); err != nil {
		return err
	}

	return nil
}

func (u *UseCase) Deactivate(id string) error {
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

func (u *UseCase) GetUserByLoginOrEmailOrId(cred string) (*domain.User, error) {
	userById, errId := u.r.GetUserById(cred)
	if errId == nil {
		return userById, nil
	}

	userByLogin, errLogin := u.r.GetUserByLogin(cred)
	if errLogin == nil {
		return userByLogin, nil
	}

	userByEmail, errEmail := u.r.GetUserByEmail(cred)
	if errEmail == nil {
		return userByEmail, nil
	}

	errStr := fmt.Sprintf("err get user: %s, %s, %s",
		errId.Error(),
		errLogin.Error(),
		errEmail.Error())

	return nil, errors.New(errStr)
}
