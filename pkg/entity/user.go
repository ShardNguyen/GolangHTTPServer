package entity

import "errors"

type User struct {
	id   int
	name string
}

func (user *User) ConvertToResponse() (ur UserResponse, err error) {
	ur = UserResponse{}

	// Convert ID
	id, err := user.GetID()
	if err != nil {
		return
	}

	// Convert Name
	name, err := user.GetName()
	if err != nil {
		return
	}

	// Make Response to return
	ur = UserResponse{
		Id:   id,
		Name: name,
	}
	return
}

func NewUser(id int, name string) *User {
	u := new(User)
	u.id = id
	u.name = name
	return u
}

func (u *User) GetID() (int, error) {
	if u == nil {
		return 0, errors.New("User is not initialized")
	}

	return u.id, nil
}

func (u *User) GetName() (string, error) {
	if u == nil {
		return "", errors.New("User is not initialized")
	}

	return u.name, nil
}

func (u *User) SetID(id int) error {
	if u == nil {
		return errors.New("User is not initialized")
	}

	u.id = id
	return nil
}

func (u *User) SetName(name string) error {
	if u == nil {
		return errors.New("User is not initialized")
	}

	u.name = name
	return nil
}
