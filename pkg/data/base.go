package data

import "github.com/ShardNguyen/GolangCounter/pkg/entity"

type Database interface {
	CreateUser(ur *entity.UserResponse) error
	GetUser(id int) (u *entity.User, err error)
	GetAllUsers() (uMap map[int]entity.User, err error)
	UpdateUser(id int, ur *entity.UserResponse) error
	DeleteUser(id int) error
}
