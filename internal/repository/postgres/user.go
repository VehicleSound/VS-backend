package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/timickb/transport-sound/internal/domain"
	"github.com/timickb/transport-sound/internal/repository"
)

type PqRepository struct {
	db *sql.DB
}

func (p PqRepository) CreateUser(login, email, pwdHash string) (string, error) {
	id := uuid.NewString()

	st := `INSERT INTO users (id, login, email, "passwordHash") 
		VALUES ($1, $2, $3, $4)`

	_, err := p.db.Exec(st, id, login, email, pwdHash)
	if err != nil {
		return "", fmt.Errorf("db insertion err %w", err)
	}

	return id, nil
}

func (p PqRepository) GetUserByLogin(login string) (*domain.User, error) {
	row := p.db.QueryRow(fmt.Sprintf("SELECT * FROM users WHERE login='%s';", login))

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

func (p PqRepository) EditUser(id string, payload *repository.UserEditPayload) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewPqRepository(db *sql.DB) *PqRepository {
	return &PqRepository{db: db}
}
