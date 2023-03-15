package postgres

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/timickb/transport-sound/internal/infrastructure/domain"
	"github.com/timickb/transport-sound/internal/infrastructure/repository"
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
	row := p.db.QueryRow(`SELECT * FROM users WHERE email=$1`, email)

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
	row := p.db.QueryRow(`SELECT * FROM users WHERE id=$1`, id)

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

func (p PqRepository) EditUser(id string, payload *repository.UserEditPayload) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (p PqRepository) AddFavourite(userId, soundId string) error {
	_, err := p.db.Exec(`INSERT INTO favourites (user_id, sound_id) VALUES ($1, $2)`, userId, soundId)

	if err != nil {
		return err
	}

	return nil
}

func NewPqRepository(db *sql.DB) *PqRepository {
	return &PqRepository{db: db}
}
