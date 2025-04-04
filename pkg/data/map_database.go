package data

import (
	"errors"

	"GolangHTTPServer/pkg/entity"
)

type mapDatabase struct {
	userData map[int]entity.User
}

var mapInstance *mapDatabase

// Get mapInstance of the map database.
// If it exists, get that mapInstance. If it doesn't, create a new mapInstance of it and return that mapInstance
func GetMapDatabaseInstance() *mapDatabase {
	// Check if mapInstance is created
	if mapInstance == nil {
		mapInstance = &mapDatabase{
			userData: make(map[int]entity.User),
		}
	}

	return mapInstance
}

// Create data by received JSON file and add user to the map database.
func (mapDB *mapDatabase) CreateUser(up *entity.UserPublic) error {
	// Check if id already exists in the map database
	if _, taken := mapDB.userData[up.Id]; taken {
		return errors.New("id is already taken")
	}

	// Convert public to private
	newUser, err := up.ConvertToUser()
	if err != nil {
		return err
	}

	mapDB.userData[up.Id] = *newUser
	return nil
}

// Get user by ID in the map database and return the said user.
// Return error if id doesn't exist
func (mapDB *mapDatabase) GetUser(id int) (u *entity.User, err error) {
	user, ok := mapDB.userData[id]
	if !ok {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

// Get all of the data contained in the map database
func (mapDB *mapDatabase) GetAllUsers() (uMap map[int]entity.User, err error) {
	return mapDB.userData, nil
}

// Update data by ID in the map database with the response received from the user.
func (mapDB *mapDatabase) UpdateUser(id int, up *entity.UserPublic) error {
	// Find if user with said ID exists
	if _, ok := mapDB.userData[id]; !ok {
		return errors.New("user not found")
	}

	// Set ID for user response
	up.SetID(id)

	// Converting user response to user data
	updatedUser, err := up.ConvertToUser()
	if err != nil {
		return errors.New("cannot edit this user")
	}

	mapDB.userData[id] = *updatedUser
	return nil
}

// Delete data by ID in the map database.
func (mapDB *mapDatabase) DeleteUser(id int) error {
	// Find if user with said ID exists
	if _, ok := mapDB.userData[id]; !ok {
		return errors.New("user not found")
	}

	delete(mapDB.userData, id)
	return nil
}
