package utilities

import (
	"GolangHTTPServer/models/entity"
	"database/sql"
)

func ConvertSQLRowsToUserMap(rows *sql.Rows) (map[int]entity.User, error) {
	defer rows.Close()

	uMap := make(map[int]entity.User)
	var u entity.User

	if err := u.Exists(); err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&u.Id, &u.Name)

		if err != nil {
			continue
		}

		uMap[u.Id] = u
	}

	return uMap, nil
}
