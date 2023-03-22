package controller

import (
	"context"
	"github.com/timickb/transport-sound/internal/infrastructure/controller/dto"
	"github.com/timickb/transport-sound/internal/infrastructure/domain"
)

type TagUseCase interface {
	CreateTag(name string) (string, error)
	GetTagById(id string) (*domain.Tag, error)
	GetTagByTitle(title string) (*domain.Tag, error)
	GetAllTags() ([]*domain.Tag, error)
}

type TagController struct {
	u TagUseCase
}

func NewTag(u TagUseCase) *TagController {
	return &TagController{u: u}
}

func (c *TagController) CreateTag(ctx context.Context, req *dto.CreateTagRequest) (*dto.CreateTagResponse, error) {
	tagId, err := c.u.CreateTag(req.Title)
	if err != nil {
		return nil, err
	}

	return &dto.CreateTagResponse{TagId: tagId}, nil
}

func (c *TagController) GetAllTags(context.Context) ([]*dto.TagResponse, error) {
	tags, err := c.u.GetAllTags()
	if err != nil {
		return nil, err
	}

	var response []*dto.TagResponse

	for _, tag := range tags {
		response = append(response, &dto.TagResponse{
			Id:    tag.Id,
			Title: tag.Title,
		})
	}

	return response, nil
}

func (c *TagController) GetTagById(ctx context.Context, id string) (*dto.TagResponse, error) {
	tag, err := c.u.GetTagById(id)
	if err != nil {
		return nil, err
	}

	return &dto.TagResponse{
		Id:    tag.Id,
		Title: tag.Title,
	}, nil
}
