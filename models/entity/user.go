package entity

import "errors"

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func NewUser(id int, name string) *User {
	u := new(User)
	u.Id = id
	u.Name = name
	return u
}

func (u *User) Exists() error {
	if u == nil {
		return errors.New("user does not exist")
	}

	return nil
}

func (u *User) SetID(id int) error {
	if err := u.Exists(); err != nil {
		return err
	}

	u.Id = id
	return nil
}

func (u *User) SetName(name string) error {
	if err := u.Exists(); err != nil {
		return err
	}

	u.Name = name
	return nil
}
