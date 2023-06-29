package database

import (
	"github.com/jmoiron/sqlx"
)

func LoginStudent(db sqlx.Ext, username string, password string, dateOfBirth string) (bool, int64, error) {

	counter := 0
	var UserID int64
	var UserName int64
	var Pass string
	var dbo string

	rows, err := db.Queryx(`SELECT u.ID, u.Username, u.Password, s.DateOfBirth FROM Users as u
	JOIN Students as s 
	ON s.UserID = u.ID
	WHERE u.Username=? AND u.Password=? AND s.DateOfBirth =?`,
		username, password, dateOfBirth)

	if err != nil {
		return false, 0, err
	}
	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&UserID, &UserName, &Pass, &dbo)
		if err != nil {
			return false, 0, err
		}
		counter++
	}

	if err := rows.Err(); err != nil {
		return false, 0, err
	}

	if counter == 0 {
		return false, 0, err
	}

	return true, UserID, nil
}

func LoginAdmin(db sqlx.Ext, username string, password string) (bool, error) {

	counter := 0

	rows, err := db.Queryx(`SELECT u.ID, u.Username, u.Password FROM Users as u
	WHERE u.Username=? AND u.Password=? AND IsAdmin=true`,
		username, password)

	if err != nil {
		return false, err
	}
	defer rows.Close()

	for rows.Next() {

		counter++
	}

	if err := rows.Err(); err != nil {
		return false, err
	}

	if counter == 0 {
		return false, nil
	}

	return true, nil
}
