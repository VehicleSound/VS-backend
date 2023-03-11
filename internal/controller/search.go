package controller

import (
	"github.com/timickb/transport-sound/internal/controller/dto"
	"github.com/timickb/transport-sound/internal/domain"
	"github.com/timickb/transport-sound/internal/usecase"
)

type SearchUseCase interface {
	Search(req *usecase.SearchRequest) ([]*domain.Sound, error)
}

type SearchController struct {
	u SearchUseCase
}

func NewSearchController(u SearchUseCase) *SearchController {
	return &SearchController{u: u}
}

func (c *SearchController) Search(req *dto.SearchRequest) ([]*domain.Sound, error) {
	res, err := c.u.Search(&usecase.SearchRequest{
		Name:       req.Name,
		TagIds:     req.TagIds,
		VehicleIds: req.VehicleIds,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
