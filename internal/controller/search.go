package controller

import (
	"context"
	dto2 "github.com/timickb/transport-sound/internal/controller/dto"
	"github.com/timickb/transport-sound/internal/domain"
	"github.com/timickb/transport-sound/internal/usecase/search"
)

type SearchUseCase interface {
	Search(req *search.Request) ([]*domain.Sound, error)
}

type SearchController struct {
	u SearchUseCase
}

func NewSearch(u SearchUseCase) *SearchController {
	return &SearchController{u: u}
}

func (c *SearchController) Search(ctx context.Context, req *dto2.SearchRequest) ([]*dto2.SoundResponse, error) {
	sounds, err := c.u.Search(&search.Request{
		Name:       req.Name,
		TagIds:     req.TagIds,
		VehicleIds: req.VehicleIds,
	})
	if err != nil {
		return nil, err
	}

	resp := make([]*dto2.SoundResponse, len(sounds))
	for i, s := range sounds {
		resp[i] = dto2.MapSound(s)
	}

	return resp, nil
}
