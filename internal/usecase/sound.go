package usecase

import (
	"fmt"
	"github.com/timickb/transport-sound/internal/domain"
)

type SoundUseCase struct {
	r Repository
}

func NewSoundUseCase(r Repository) *SoundUseCase {
	return &SoundUseCase{r: r}
}

func (u *SoundUseCase) GetSoundById(id string) (*domain.Sound, error) {
	sound, err := u.r.GetSoundById(id)
	if err != nil {
		return nil, fmt.Errorf("get sound by id err: %w", err)
	}

	tags, err := u.r.GetTagsForSound(sound.Id)
	if err != nil {
		return nil, fmt.Errorf("get sound by id err: %w", err)
	}

	sound.Tags = tags
	return sound, nil
}

func (u *SoundUseCase) GetAllSounds() ([]*domain.Sound, error) {
	sounds, err := u.r.GetAllSounds()
	if err != nil {
		return nil, err
	}

	return sounds, nil
}
