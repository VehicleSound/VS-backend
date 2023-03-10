package domain

type Sound struct {
	Id            string
	Name          string
	Description   string
	AuthorId      string
	PictureFileId string
	SoundFileId   string
	VehicleId     string

	AuthorLogin string
	VehicleName string

	Tags []*Tag
}
