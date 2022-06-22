package model

type RegisterRequest struct {
	User *User `json:"user"`
}

type RegisterResponse struct{}
