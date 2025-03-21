package data

import (
	"errors"

	"github.com/ShardNguyen/GolangCounter/pkg/entity"
)

type mapDatabase struct {
	userData map[int]entity.User
}

var instance *mapDatabase

// Get instance of the map database.
// If it exists, get that instance. If it doesn't, create a new instance of it and return that instance
func GetMapDatabaseInstance() *mapDatabase {
	// Check if instance is created
	if instance == nil {
		instance = &mapDatabase{
			userData: make(map[int]entity.User),
		}
	}

	return instance
}

// Create data by ID in the map database.
func (mapDB *mapDatabase) CreateUser(ur *entity.UserResponse) error {
	newID := generateId(mapDB.userData)
	ur.SetID(newID)

	// Convert response to user data
	newUser, err := ur.ConvertToUser()
	if err != nil {
		return err
	}

	mapDB.userData[newID] = *newUser
	return nil
}

// Get data by ID in the map database and return the said data.
// Return error if data is not retrievable
func (mapDB *mapDatabase) GetUser(id int) (u *entity.User, err error) {
	user, ok := mapDB.userData[id]
	if !ok {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

// Get all of the data contained in the hash map
func (mapDB *mapDatabase) GetAllUsers() (uMap map[int]entity.User, err error) {
	return mapDB.userData, nil
}

// Update data by ID in the map database with the response received from the user.
func (mapDB *mapDatabase) UpdateUser(id int, ur *entity.UserResponse) error {
	// Find if user with said ID exists
	if _, ok := mapDB.userData[id]; !ok {
		return errors.New("user not found")
	}

	// Set ID for user response
	ur.SetID(id)

	// Converting user response to user data
	updatedUser, err := ur.ConvertToUser()
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

// Basically get the highest ID and add 1 to it
func generateId(users map[int]entity.User) int {
	var maxId int

	for id := range users {
		if id > maxId {
			maxId = id
		}
	}

	return maxId + 1
}
