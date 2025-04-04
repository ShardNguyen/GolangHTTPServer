package data

import (
	"GolangHTTPServer/model/entity"
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

type postgresDatabase struct {
	userData *sql.DB
}

var postgresInstance *postgresDatabase
var mutexPostgres = &sync.Mutex{}

// Get instance of PostgreSQL. If the instance doesn't exist, create a new instance and return the new instance. Otherwise, return the PostgreSQL instance
func GetPostgreSQLInstance() (*postgresDatabase, error) {
	mutexPostgres.Lock()

	if postgresInstance == nil {
		// Opens a connection to a Postgres Database using DATABASE URL
		db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		postgresInstance = &postgresDatabase{
			userData: db,
		}

	}

	mutexPostgres.Unlock()
	return postgresInstance, nil
}

// Create data by received JSON file and add user to the PostgreSQL database.
func (postgresDB *postgresDatabase) CreateUser(up *entity.UserPublic) error {
	queryString := fmt.Sprintf("INSERT INTO users (name) "+"VALUES ('%s') RETURNING id", up.Name)
	// Scan copy the ID created from SQL Database
	// and pasted it into the id in User Public struct
	err := postgresDB.userData.QueryRow(queryString).Scan(&up.Id)
	return err
}

// Get user by ID in the Postgres database and return the said user.
// Return error if the database returns no row.
func (postgresDB *postgresDatabase) GetUser(id int) (*entity.User, error) {
	var up entity.UserPublic
	queryString := fmt.Sprintf("SELECT * FROM users WHERE id = %d;", id)

	// The Scan function is used to copy values
	// from the SQL database to the user's public struct
	err := postgresDB.userData.QueryRow(queryString).Scan(&up.Id, &up.Name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	u, err := up.ConvertToUser()
	return u, err
}

// Get all of the data contained in the Postgres database
func (postgresDB *postgresDatabase) GetAllUsers() (map[int]entity.User, error) {
	rows, err := postgresDB.userData.Query("SELECT * FROM users;")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	uMap := make(map[int]entity.User)
	var up entity.UserPublic

	for rows.Next() {
		// Scan for user data, skip the current line if the scan returns error
		err := rows.Scan(&up.Id, &up.Name)
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
	queryString := fmt.Sprintf("UPDATE users SET name = %s WHERE id = %d;", up.Name, id)
	_, err := postgresDB.userData.Exec(queryString)
	return err
}

// Delete data in the Postgres Database by ID.
func (postgresDB *postgresDatabase) DeleteUser(id int) error {
	_, err := postgresDB.GetUser(id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	queryString := fmt.Sprintf("DELETE FROM users WHERE id = %d", id)
	_, err = postgresDB.userData.Exec(queryString)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (postgresDB *postgresDatabase) CloseConnection() {
	postgresDB.userData.Close()
}
