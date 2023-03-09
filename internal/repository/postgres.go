package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/timickb/transport-sound/internal/domain"
	"time"
)

type PqRepository struct {
	db *sql.DB
}

func (p PqRepository) CreateUser(login, email, pwdHash string) (*domain.User, error) {
	id := uuid.NewString()

	st := `INSERT INTO users (id, login, email, "passwordHash") 
		VALUES ($1, $2, $3, $4)`

	_, err := p.db.Exec(st, id, login, email, pwdHash)
	if err != nil {
		return nil, fmt.Errorf("db insertion err %w", err)
	}

	return &domain.User{
		Id:           id,
		Login:        "0",
		Email:        "0",
		Confirmed:    false,
		Active:       false,
		DateCreated:  time.Now(),
		PasswordHash: "",
	}, nil

}

func (p PqRepository) GetUserByLogin(login string) (*domain.User, error) {
	row := p.db.QueryRow(fmt.Sprintf("SELECT * FROM users WHERE login='%s';", login))

	if row == nil {
		return nil, errors.New("no such user")
	}

	result := &domain.User{}

	err := row.Scan(
		result.Id,
		result.Login,
		result.Email, result.Confirmed,
		result.PasswordHash,
		result.Active)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p PqRepository) GetUserByEmail(email string) (*domain.User, error) {
	row := p.db.QueryRow(fmt.Sprintf("SELECT * FROM users WHERE email='%s';", email))

	if row == nil {
		return nil, errors.New("no such user")
	}

	result := &domain.User{}

	err := row.Scan(
		&result.Id,
		&result.Login,
		&result.Email,
		&result.Confirmed,
		&result.PasswordHash,
		&result.Active)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p PqRepository) GetUserById(id string) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (p PqRepository) EditUser(id string, payload *UserEditPayload) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (p PqRepository) CreateTag(title string) (*domain.Tag, error) {
	//TODO implement me
	panic("implement me")
}

func (p PqRepository) GetTagById(id string) (*domain.Tag, error) {
	//TODO implement me
	panic("implement me")
}

func (p PqRepository) GetTagByTitle(title string) (*domain.Tag, error) {
	//TODO implement me
	panic("implement me")
}

func (p PqRepository) GetAllSounds(limit int) ([]*domain.Sound, error) {
	//TODO implement me
	panic("implement me")
}

func (p PqRepository) GetSoundsNameLike(name string) ([]*domain.Sound, error) {
	//TODO implement me
	panic("implement me")
}

func (p PqRepository) GetSoundsByTagId(tagId string) ([]*domain.Sound, error) {
	//TODO implement me
	panic("implement me")
}

func (p PqRepository) GetSoundsByVehicleId(vehicleId string) ([]*domain.Sound, error) {
	//TODO implement me
	panic("implement me")
}

func NewPqRepository(db *sql.DB) *PqRepository {
	return &PqRepository{db: db}
}
