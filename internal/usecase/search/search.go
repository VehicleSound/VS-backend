package search

import (
	"fmt"
	"github.com/timickb/transport-sound/internal/domain"
	"github.com/timickb/transport-sound/internal/interfaces"
	"github.com/timickb/transport-sound/internal/usecase"
)

type Request struct {
	Name       string
	TagIds     []string
	VehicleIds []string
}

type UseCase struct {
	repo usecase.Repository
	log  interfaces.Logger
}

func New(repo usecase.Repository, log interfaces.Logger) *UseCase {
	return &UseCase{repo: repo, log: log}
}

func (u *UseCase) Search(req *Request) ([]*domain.Sound, error) {
	byName, err := u.repo.GetSoundsNameLike(req.Name)
	if err != nil {
		return nil, fmt.Errorf("search err: %w", err)
	}

	suitable := make([]*domain.Sound, 0)

	for _, sound := range byName {
		if len(req.VehicleIds) > 0 && !contains(req.VehicleIds, sound.VehicleId) {
			continue
		}
		if len(req.TagIds) > 0 {
			for _, tag := range sound.Tags {
				if !contains(req.TagIds, tag.Id) {
					continue
				}
			}
		}

		tags, err := u.repo.GetTagsForSound(sound.Id)
		if err != nil {
			return nil, fmt.Errorf("search err: %w", err)
		}

		sound.Tags = tags
		suitable = append(suitable, sound)
	}

	return suitable, nil
}

func contains(slice []string, s string) bool {
	for _, a := range slice {
		if a == s {
			return true
		}
	}
	return false
}
