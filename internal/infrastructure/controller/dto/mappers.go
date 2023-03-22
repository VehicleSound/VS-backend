package dto

import (
	"github.com/timickb/transport-sound/internal/infrastructure/domain"
)

func MapSound(s *domain.Sound) *SoundResponse {
	tags := make([]*TagResponse, len(s.Tags))

	for i, t := range s.Tags {
		tags[i] = &TagResponse{
			Id:    t.Id,
			Title: t.Title,
		}
	}

	return &SoundResponse{
		Id:             s.Id,
		Name:           s.Name,
		Description:    s.Description,
		PictureFileUrl: "/assets/images/" + s.PictureFileId + ".jpg",
		SoundFileUrl:   "/assets/sounds/" + s.SoundFileId + ".mp3",
		AuthorId:       s.AuthorId,
		VehicleId:      s.VehicleId,
		AuthorLogin:    s.AuthorLogin,
		VehicleName:    s.VehicleName,
		Tags:           tags,
	}
}
