package auth

import "inventory-ticketing-system/domain/entity"

type LoginResponse struct {
	Token string        `json:"token"`
	User  *entity.User  `json:"user"`
}