package entity

import "errors"

type UserResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (ur *UserResponse) ConvertToUser() (u *User, err error) {
	u = &User{}

	if ur == nil {
		err = errors.New("User Response is not detected")
		return
	}

	u.SetID(ur.Id)
	u.SetName(ur.Name)
	return
}

func (ur *UserResponse) SetID(id int) {
	ur.Id = id
}
