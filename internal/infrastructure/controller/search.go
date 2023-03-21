package controller

import (
	"context"
	"github.com/timickb/transport-sound/internal/infrastructure/controller/dto"
	"github.com/timickb/transport-sound/internal/infrastructure/domain"
	"github.com/timickb/transport-sound/internal/infrastructure/usecase/search"
)

type SearchUseCase interface {
	Search(req *search.Request) ([]*domain.Sound, error)
}

type SearchController struct {
	u SearchUseCase
}

func NewSearchController(u SearchUseCase) *SearchController {
	return &SearchController{u: u}
}

func (c *SearchController) Search(ctx context.Context, req *dto.SearchRequest) ([]*dto.SoundResponse, error) {
	sounds, err := c.u.Search(&search.Request{
		Name:       req.Name,
		TagIds:     req.TagIds,
		VehicleIds: req.VehicleIds,
	})
	if err != nil {
		return nil, err
	}

	resp := make([]*dto.SoundResponse, len(sounds))
	for i, s := range sounds {
		resp[i] = mapSound(s)
	}

	return resp, nil
}
