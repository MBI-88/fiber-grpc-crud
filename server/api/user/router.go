package user

import (
	"server/models"
)


func NewUser() *user {
	return &user{m: models.User{}}
}
