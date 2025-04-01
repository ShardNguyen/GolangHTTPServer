package entity

import "errors"

type UserPublic struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (up *UserPublic) ConvertToUser() (u *User, err error) {
	u = &User{}

	if up == nil {
		err = errors.New("User Response is not detected")
		return
	}

	u.SetID(up.Id)
	u.SetName(up.Name)
	return
}

func (up *UserPublic) SetID(id int) {
	up.Id = id
}
