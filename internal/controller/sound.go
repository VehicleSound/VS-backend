package controller

import (
	"github.com/timickb/transport-sound/internal/controller/dto"
	"github.com/timickb/transport-sound/internal/domain"
)

type SoundUseCase interface {
	GetAllSounds() ([]*domain.Sound, error)
	GetSoundById(id string) (*domain.Sound, error)
	CreateSound(s *domain.Sound, tid []string) (string, error)
}

type SoundController struct {
	u SoundUseCase
}

func NewSoundController(u SoundUseCase) *SoundController {
	return &SoundController{u: u}
}

func (c *SoundController) GetAllSounds() ([]*domain.Sound, error) {
	sounds, err := c.u.GetAllSounds()
	if err != nil {
		return nil, err
	}
	return sounds, nil
}

func (c *SoundController) GetSoundById(id string) (*domain.Sound, error) {
	sound, err := c.u.GetSoundById(id)
	if err != nil {
		return nil, err
	}
	return sound, nil
}

func (c *SoundController) CreateSound(req *dto.CreateSoundRequest) (*dto.CreateSoundResponse, error) {
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
