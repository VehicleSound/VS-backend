package controller

import (
	"context"
	dto2 "github.com/timickb/transport-sound/internal/controller/dto"
	"github.com/timickb/transport-sound/internal/domain"
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

func (c *TagController) CreateTag(ctx context.Context, req *dto2.CreateTagRequest) (*dto2.CreateTagResponse, error) {
	tagId, err := c.u.CreateTag(req.Title)
	if err != nil {
		return nil, err
	}

	return &dto2.CreateTagResponse{TagId: tagId}, nil
}

func (c *TagController) GetAllTags(context.Context) ([]*dto2.TagResponse, error) {
	tags, err := c.u.GetAllTags()
	if err != nil {
		return nil, err
	}

	var response []*dto2.TagResponse

	for _, tag := range tags {
		response = append(response, &dto2.TagResponse{
			Id:    tag.Id,
			Title: tag.Title,
		})
	}

	return response, nil
}

func (c *TagController) GetTagById(ctx context.Context, id string) (*dto2.TagResponse, error) {
	tag, err := c.u.GetTagById(id)
	if err != nil {
		return nil, err
	}

	return &dto2.TagResponse{
		Id:    tag.Id,
		Title: tag.Title,
	}, nil
}
