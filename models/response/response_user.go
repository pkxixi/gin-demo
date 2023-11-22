package response

import "go-blog/models"

type UserResponse struct {
	User models.User `json:"user"`
}

type LoginResponse struct {
	User      models.User `json:"user"`
	Token     string      `json:"token"`
	ExpiredAt int64       `json:"expiredAt"`
}
