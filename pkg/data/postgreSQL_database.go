package data

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/ShardNguyen/GolangCounter/pkg/entity"
	_ "github.com/lib/pq"
)

type postgresDatabase struct {
	userData *sql.DB
}

var postgresInstance *postgresDatabase

// Get instance of PostgreSQL. If the instance doesn't exist, create a new instance and return the new instance. Otherwise, return the PostgreSQL instance
func GetPostgreSQLInstance() (*postgresDatabase, error) {
	if postgresInstance == nil {
		// Opens a connection to a Postgres Database using DATABASE URL
		db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

		if err != nil {
			return nil, err
		}

		postgresInstance = &postgresDatabase{
			userData: db,
		}
	}
	return postgresInstance, nil
}

// Create data by received JSON file and add user to the PostgreSQL database.
func (postgresDB *postgresDatabase) CreateUser(up *entity.UserPublic) error {
	queryString := fmt.Sprintf("INSERT INTO users (id, name) "+"VALUES (%d, %s)", up.Id, up.Name)
	return postgresDB.userData.QueryRow(queryString).Scan(&up)
}

// Get user by ID in the Postgres database and return the said user.
// Return error if the database returns no row.
func (postgresDB *postgresDatabase) GetUser(id int) (*entity.User, error) {
	queryString := fmt.Sprintf("SELECT * FROM users WHERE id = %d", id)

	var up entity.UserPublic
	// Scan through the postgres database and return the first data row that has the said id
	err := postgresDB.userData.QueryRow(queryString).Scan(&up)

	if err != nil {
		return nil, err
	}

	u, err := up.ConvertToUser()
	return u, err
}

// Get all of the data contained in the Postgres database
func (postgresDB *postgresDatabase) GetAllUsers() (map[int]entity.User, error) {
	rows, err := postgresDB.userData.Query("SELECT * FROM users")

	if err != nil {
		return nil, err
	}

	uMap := make(map[int]entity.User)

	for rows.Next() {
		var up entity.UserPublic

		// Scan for user data, skip the current line if the scan returns error
		err := rows.Scan(&up)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Convert public to private, skip the current line if this process returns error
		u, err := up.ConvertToUser()
		if err != nil {
			fmt.Println(err)
			continue
		}

		uMap[up.Id] = *u
	}

	return uMap, err
}

// Update data in the Postgres Database by ID and the response received from the user
func (postgresDB *postgresDatabase) UpdateUser(id int, up *entity.UserPublic) error {
	queryString := fmt.Sprintf("UPDATE users SET name = %s WHERE id = %d", up.Name, id)
	return postgresDB.userData.QueryRow(queryString).Scan(&up)
}

// Delete data in the Postgres Database by ID.
func (postgresDB *postgresDatabase) DeleteUser(id int) error {
	queryString := fmt.Sprintf("DELETE FROM users WHERE id = %d", id)
	_, err := postgresDB.userData.Exec(queryString)
	return err
}
