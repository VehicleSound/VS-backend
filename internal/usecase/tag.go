package usecase

import (
	"errors"
	"fmt"
	"github.com/timickb/transport-sound/internal/domain"
)

type TagUseCase struct {
	repo Repository
}

func NewTagUseCase(repo Repository) *TagUseCase {
	return &TagUseCase{repo: repo}
}

func (u *TagUseCase) CreateTag(name string) error {
	if !validateTag(name) {
		return errors.New("err create tag title too short")
	}

	_, err := u.repo.CreateTag(name)
	if err != nil {
		return fmt.Errorf("err create tag: %w", err)
	}

	return nil
}

func (u *TagUseCase) GetTagById(id string) (*domain.Tag, error) {
	tag, err := u.repo.GetTagById(id)
	if err != nil {
		return nil, fmt.Errorf("err get tag by id: %w", err)
	}

	return tag, nil
}

func (u *TagUseCase) GetTagByTitle(title string) (*domain.Tag, error) {
	tag, err := u.repo.GetTagByTitle(title)
	if err != nil {
		return nil, fmt.Errorf("err get tag by title: %w", err)
	}

	return tag, nil
}

