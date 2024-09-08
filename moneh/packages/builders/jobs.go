package builders

import (
	"database/sql"
)

func GetTotalCount(con *sql.DB, table string, view *string) (int, error) {
	var count int
	var sqlStatement string

	// Fix this. if table empty, there will be an error
	if view != nil {
		sqlStatement = "SELECT COUNT(*) FROM " + table + " WHERE " + *view
	} else {
		sqlStatement = "SELECT COUNT(*) FROM " + table
	}

	err := con.QueryRow(sqlStatement).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func ValidateOwner(con *sql.DB, table, token, id string) (bool, error) {
	sqlStatement := "SELECT 1 FROM " + table + " JOIN users_tokens ON users_tokens.context_id = " + table + ".created_by WHERE token = ? AND " + table + ".id = ? LIMIT 1"
	err := con.QueryRow(sqlStatement, token, id).Scan(new(int))
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func GetUserIdFromToken(con *sql.DB, token string) (string, error) {
	var id string

	sqlStatement := "SELECT users.id FROM users JOIN users_tokens ON users_tokens.context_id = users.id WHERE token = ? LIMIT 1"
	err := con.QueryRow(sqlStatement, token).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}

	return id, nil
}
