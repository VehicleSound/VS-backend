package controller

import "github.com/timickb/transport-sound/internal/domain"

type SoundUseCase interface {
	GetAllSounds() ([]*domain.Sound, error)
	GetSoundById(id string) (*domain.Sound, error)
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
