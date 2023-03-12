package controller

import (
	"github.com/timickb/transport-sound/internal/controller/dto"
	"github.com/timickb/transport-sound/internal/domain"
)

func mapSound(s *domain.Sound) *dto.SoundResponse {
	return &dto.SoundResponse{
		Id:             s.Id,
		Name:           s.Name,
		Description:    s.Description,
		PictureFileUrl: "/assets/images/" + s.PictureFileId + ".jpg",
		SoundFileUrl:   "/assets/sounds/" + s.SoundFileId + ".mp3",
		AuthorId:       s.AuthorId,
		VehicleId:      s.VehicleId,
		AuthorLogin:    s.AuthorLogin,
		VehicleName:    s.VehicleName,
	}
}
