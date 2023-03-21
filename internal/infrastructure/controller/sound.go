package controller

import (
	"context"
	"github.com/timickb/transport-sound/internal/infrastructure/controller/dto"
	"github.com/timickb/transport-sound/internal/infrastructure/domain"
)

type SoundUseCase interface {
	GetAllSounds() ([]*domain.Sound, error)
	GetSoundById(id string) (*domain.Sound, error)
	CreateSound(s *domain.Sound, tid []string) (string, error)
	GetRandomSounds(limit int) ([]*domain.Sound, error)
}

type SoundController struct {
	u SoundUseCase
}

func NewSoundController(u SoundUseCase) *SoundController {
	return &SoundController{u: u}
}

func (c *SoundController) GetAllSounds(context.Context) ([]*dto.SoundResponse, error) {
	sounds, err := c.u.GetAllSounds()
	if err != nil {
		return nil, err
	}

	resp := make([]*dto.SoundResponse, len(sounds))
	for i, s := range sounds {
		resp[i] = mapSound(s)
	}
	return resp, nil
}

func (c *SoundController) GetSoundById(ctx context.Context, id string) (*dto.SoundResponse, error) {
	sound, err := c.u.GetSoundById(id)
	if err != nil {
		return nil, err
	}
	return mapSound(sound), nil
}

func (c *SoundController) CreateSound(ctx context.Context, t *dto.TokenResponse, req *dto.CreateSoundRequest) (*dto.CreateSoundResponse, error) {
	sound := &domain.Sound{
		Name:          req.Name,
		Description:   req.Description,
		AuthorId:      req.AuthorId,
		PictureFileId: req.PictureFileId,
		SoundFileId:   req.SoundFileId,
		VehicleId:     req.VehicleId,
	}

	id, err := c.u.CreateSound(sound, req.TagIds)
	if err != nil {
		return nil, err
	}

	return &dto.CreateSoundResponse{SoundId: id}, nil
}

func (c *SoundController) GetRandomSounds(ctx context.Context, limit int) ([]*dto.SoundResponse, error) {
	sounds, err := c.u.GetRandomSounds(limit)
	if err != nil {
		return nil, err
	}

	resp := make([]*dto.SoundResponse, len(sounds))
	for i, s := range sounds {
		resp[i] = mapSound(s)
	}
	return resp, nil
}
