package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func ChangePassword(db sqlx.Ext, userID int64, studentNumber int64, pass string, newpassword string, hashedNewPassword string) error {

	fmt.Println(studentNumber)
	fmt.Println(pass)
	fmt.Println(newpassword)
	fmt.Println(hashedNewPassword)
	var username int64
	var password string

	rows, err := db.Queryx(`SELECT u.Username, u.Password FROM Users as u
	WHERE u.Username=? AND u.Password=?`,
		studentNumber, pass)

	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&username, &password)
		if err != nil {
			return err
		}
	}

	if pass != password {
		return fmt.Errorf("Wrong password!")
	}

	_, err = db.Exec(`UPDATE Users SET Password= ?, PlainPassword = ? WHERE ID= ?`,
		hashedNewPassword,
		newpassword,
		userID,
	)

	if err != nil {
		return err
	}

	return nil
}
