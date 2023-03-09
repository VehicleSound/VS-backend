package usecase

import "github.com/timickb/transport-sound/internal/domain"

type SearchRequest struct {
	Name       string
	TagIds     []string
	VehicleIds []string
}

type SearchUseCase struct {
	repo Repository
}

func NewSearchUseCase(repo Repository) *SearchUseCase {
	return &SearchUseCase{repo: repo}
}

func (u *SearchUseCase) Search(req *SearchRequest) ([]*domain.Sound, error) {
	return nil, nil
}
