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

type GetUserResponse struct {
	Id        string `json:"id,omitempty"`
	Login     string `json:"login,omitempty"`
	Email     string `json:"email,omitempty"`
	Active    bool   `json:"active,omitempty"`
	Confirmed bool   `json:"confirmed,omitempty"`
}

type TokenResponse struct {
	Id        string `json:"id,omitempty"`
	Login     string `json:"login,omitempty"`
	Email     string `json:"email,omitempty"`
	Confirmed bool   `json:"confirmed,omitempty"`
	Active    bool   `json:"active,omitempty"`
}

type SoundResponse struct {
	Id             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Description    string `json:"description,omitempty"`
	PictureFileUrl string `json:"picture_file_url,omitempty"`
	SoundFileUrl   string `json:"sound_file_url,omitempty"`
	AuthorId       string `json:"author_id,omitempty"`
	VehicleId      string `json:"vehicle_id"`
	AuthorLogin    string `json:"author_login,omitempty"`
	VehicleName    string `json:"vehicle_name"`
}
