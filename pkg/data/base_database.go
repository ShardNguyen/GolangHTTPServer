package data

import "github.com/ShardNguyen/GolangCounter/pkg/entity"

type Database interface {
	CreateUser(ur *entity.UserPublic) error
	GetUser(id int) (u *entity.User, err error)
	GetAllUsers() (uMap map[int]entity.User, err error)
	UpdateUser(id int, ur *entity.UserPublic) error
	DeleteUser(id int) error
}
