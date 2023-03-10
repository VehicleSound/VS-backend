package dto

import "mime/multipart"

type AuthRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type RegisterRequest struct {
	Email    string `json:"email,omitempty"`
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
}

type ChangeLoginRequest struct {
	UserId string `json:"user_id,omitempty"`
	Login  string `json:"login,omitempty" json:"login,omitempty"`
}

type ChangeEmailRequest struct {
	UserId string `json:"user_id,omitempty"`
	Email  string `json:"email,omitempty" json:"email,omitempty"`
}

type ChangePasswordRequest struct {
	UserId      string `json:"user_id,omitempty"`
	OldPassword string `json:"old_password,omitempty"`
	NewPassword string `json:"new_password,omitempty"`
}

type CreateTagRequest struct {
	Title string `json:"title,omitempty"`
}

type CreateSoundRequest struct {
	Name          string   `json:"name,omitempty"`
	Description   string   `json:"description,omitempty"`
	PictureFileId string   `json:"picture_id,omitempty"`
	SoundFileId   string   `json:"sound_id,omitempty"`
	TagIds        []string `json:"tag_ids,omitempty"`
	VehicleId     string   `json:"vehicle_id,omitempty"`
	AuthorId      string   `json:"author_id,omitempty"`
}

type UploadFileRequest struct {
	File *multipart.FileHeader `form:"file"`
}
