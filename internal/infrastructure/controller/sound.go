package controller

import (
	dto2 "github.com/timickb/transport-sound/internal/infrastructure/controller/dto"
	"github.com/timickb/transport-sound/internal/infrastructure/domain"
	"github.com/timickb/transport-sound/internal/infrastructure/usecase"
)

type SoundUseCase interface {
	GetAllSounds() ([]*domain.Sound, error)
	GetSoundById(id string) (*domain.Sound, error)
	CreateSound(ctx usecase.UserContext, s *domain.Sound, tid []string) (string, error)
	GetRandomSounds(limit int) ([]*domain.Sound, error)
}

type SoundController struct {
	u SoundUseCase
}

func NewSoundController(u SoundUseCase) *SoundController {
	return &SoundController{u: u}
}

func (c *SoundController) GetAllSounds() ([]*dto2.SoundResponse, error) {
	sounds, err := c.u.GetAllSounds()
	if err != nil {
		return nil, err
	}

	resp := make([]*dto2.SoundResponse, len(sounds))
	for i, s := range sounds {
		resp[i] = mapSound(s)
	}
	return resp, nil
}

func (c *SoundController) GetSoundById(id string) (*dto2.SoundResponse, error) {
	sound, err := c.u.GetSoundById(id)
	if err != nil {
		return nil, err
	}
	return mapSound(sound), nil
}

func (c *SoundController) CreateSound(t *dto2.TokenResponse, req *dto2.CreateSoundRequest) (*dto2.CreateSoundResponse, error) {
	sound := &domain.Sound{
		Name:          req.Name,
		Description:   req.Description,
		AuthorId:      req.AuthorId,
		PictureFileId: req.PictureFileId,
		SoundFileId:   req.SoundFileId,
		VehicleId:     req.VehicleId,
	}

	id, err := c.u.CreateSound(nil, sound, req.TagIds)
	if err != nil {
		return nil, err
	}

	return &dto2.CreateSoundResponse{SoundId: id}, nil
}

func (c *SoundController) GetRandomSounds(limit int) ([]*dto2.SoundResponse, error) {
	sounds, err := c.u.GetRandomSounds(limit)
	if err != nil {
		return nil, err
	}

	resp := make([]*dto2.SoundResponse, len(sounds))
	for i, s := range sounds {
		resp[i] = mapSound(s)
	}
	return resp, nil
}