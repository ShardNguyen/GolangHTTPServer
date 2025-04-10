package database

import "GolangHTTPServer/models/entity"

type Database interface {
	CreateUser(u *entity.User) error
	GetUserByID(id int) (*entity.User, error)
	GetAllUsers() (map[int]entity.User, error)
	UpdateUserByID(id int, u *entity.User) error
	DeleteUserByID(id int) error
	CloseConnection()
}
