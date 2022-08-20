package store

import (
	"authorizationServer/model/user"
)

type UserRepository interface {
	Create(*user.User) error
	FindByID(int) (*user.User, error)
	FindByEmail(string) (*user.User, error)
	DeleteUserByEmail(string) error
	UpdateUser(*user.User) error
}
