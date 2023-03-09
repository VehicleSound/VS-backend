package postgres

import "github.com/timickb/transport-sound/internal/domain"

func (p PqRepository) GetAllSounds(limit int) ([]*domain.Sound, error) {
	//TODO implement me
	panic("implement me")
}

func (p PqRepository) GetSoundsNameLike(name string) ([]*domain.Sound, error) {
	//TODO implement me
	panic("implement me")
}

func (p PqRepository) GetSoundsByTagId(tagId string) ([]*domain.Sound, error) {
	//TODO implement me
	panic("implement me")
}

func (p PqRepository) GetSoundsByVehicleId(vehicleId string) ([]*domain.Sound, error) {
	//TODO implement me
	panic("implement me")
}
