package controller

type AuthResponse struct {
	Token string `json:"token,omitempty"`
}

type RegisterResponse struct {
	UserId string `json:"user_id,omitempty"`
}

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
