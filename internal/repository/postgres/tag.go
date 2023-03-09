package postgres

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/timickb/transport-sound/internal/domain"
)

func (p PqRepository) CreateTag(title string) (*domain.Tag, error) {
	st := `INSERT INTO tags (id, title) VALUES ($1, $2)`
	id := uuid.NewString()

	_, err := p.db.Exec(st, id, title)
	if err != nil {
		return nil, fmt.Errorf("db insertion err: %w", err)
	}

	return &domain.Tag{
		Id:    id,
		Title: title,
	}, nil
}

func (p PqRepository) GetTagById(id string) (*domain.Tag, error) {
	//TODO implement me
	panic("implement me")
}

func (p PqRepository) GetTagByTitle(title string) (*domain.Tag, error) {
	//TODO implement me
	panic("implement me")
}
