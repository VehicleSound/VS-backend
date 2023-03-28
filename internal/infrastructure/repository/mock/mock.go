package mock

import (
	"github.com/google/uuid"
	domain2 "github.com/timickb/transport-sound/internal/domain"
	"github.com/timickb/transport-sound/internal/infrastructure/repository/memory"
)

type Repository struct {
	memory.Repository
}

func NewRepository() *Repository {
	return &Repository{*memory.NewRepository()}
}

func (r *Repository) CreateTestSounds(names []string) error {
	authorId, err := r.CreateUser("test", "test", "test")
	if err != nil {
		return err
	}

	for _, name := range names {
		err := r.CreateSound(&domain2.Sound{
			Id:          uuid.NewString(),
			Name:        name,
			AuthorId:    authorId,
			AuthorLogin: "test",
			Tags:        make([]*domain2.Tag, 0),
		})
		if err != nil {
			return err
		}
	}

	return nil
}
