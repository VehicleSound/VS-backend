package mock

import (
	"github.com/google/uuid"
	"github.com/timickb/transport-sound/internal/infrastructure/domain"
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
		err := r.CreateSound(&domain.Sound{
			Id:          uuid.NewString(),
			Name:        name,
			AuthorId:    authorId,
			AuthorLogin: "test",
			Tags:        make([]*domain.Tag, 0),
		})
		if err != nil {
			return err
		}
	}

	return nil
}
