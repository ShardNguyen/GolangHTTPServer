package postgresql

import (
	"GolangHTTPServer/models/entity"
	"GolangHTTPServer/utilities"
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

type postgresql struct {
	database *sql.DB
}

var postgresqlInstance *postgresql
var mutexPostgres = &sync.Mutex{}

func GetPostgresqlInstance() (*postgresql, error) {
	defer mutexPostgres.Unlock()
	mutexPostgres.Lock()

	if postgresqlInstance == nil {
		db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

		if err != nil {
			return nil, err
		}

		postgresqlInstance = &postgresql{
			database: db,
		}
	}

	return postgresqlInstance, nil
}

func (postgres *postgresql) CreateUser(u *entity.User) error {
	if err := u.Exists(); err != nil {
		return err
	}

	queryString := fmt.Sprintf("INSERT INTO users (name) "+"VALUES ('%s') RETURNING id;", u.Name)
	err := postgres.database.QueryRow(queryString).Scan(&u.Id)

	return err
}

func (postgres *postgresql) GetUserByID(id int) (*entity.User, error) {
	queryString := fmt.Sprintf("SELECT * FROM users WHERE id = %d;", id)

	var u entity.User
	if err := u.Exists(); err != nil {
		return nil, err
	}

	err := postgres.database.QueryRow(queryString).Scan(&u.Id, &u.Name)

	if err != nil {
		return nil, err
	}

	return &u, err
}

func (postgres *postgresql) GetAllUsers() (map[int]entity.User, error) {
	rows, err := postgres.database.Query("SELECT * FROM users;")

	if err != nil {
		return nil, err
	}

	return utilities.ConvertSQLRowsToUserMap(rows)
}

func (postgres *postgresql) UpdateUserByID(id int, u *entity.User) error {
	if err := u.Exists(); err != nil {
		return err
	}

	if err := u.SetID(id); err != nil {
		return err
	}

	queryString := fmt.Sprintf("UPDATE users SET name = '%s' WHERE id = %d;", u.Name, u.Id)
	_, err := postgres.database.Exec(queryString)

	return err
}

func (postgres *postgresql) DeleteUserByID(id int) error {
	_, err := postgres.GetUserByID(id)

	if err != nil {
		return err
	}

	queryString := fmt.Sprintf("DELETE FROM users WHERE id = %d;", id)

	_, err = postgres.database.Exec(queryString)
	if err != nil {
		return err
	}

	return nil
}

func (postgres *postgresql) CloseConnection() {
	postgres.database.Close()
}
