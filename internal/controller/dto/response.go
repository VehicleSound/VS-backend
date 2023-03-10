package dto

type AuthResponse struct {
	Token string `json:"token,omitempty"`
}

type RegisterResponse struct {
	UserId string `json:"user_id,omitempty"`
}

type CreateTagResponse struct {
	TagId string `json:"tag_id,omitempty"`
}

type TagResponse struct {
	Id    string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

type UploadFileResponse struct {
	FileId string `json:"file_id,omitempty"`
}

type CreateSoundResponse struct {
	SoundId string `json:"sound_id,omitempty"`
}
