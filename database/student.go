package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ReadStudentResponse struct {
	Username      int64  `json:"Username"`
	Password      string `json:"Password,omitempty"`
	StudentNumber int64  `json:"StudentNumber"`
	UserID        int    `json:"UserID"`
	FirstName     string `json:"FirstName"`
	LastName      string `json:"LastName"`
	MiddleName    string `json:"MiddleName"`
	Email         string `json:"Email"`
	DateOfBirth   string `json:"DateOfBirth"`
	GradeLevel    string `json:"GradeLevel"`
	Address       string `json:"Address"`
	ContactNumber string `json:"ContactNumber"`
}

func CreateStudent(db sqlx.Ext, StudentNumber int, FirstName string, LastName string, MiddleName, Password string, Email string, DateOfBirth string, GradeLevel string, ContactNum string, Address string, plainpass string) error {

	// check if studentNumber exists
	var Count int64
	Exists := false
	rows, err := db.Queryx(`SELECT COUNT(Username) FROM Users WHERE Username = ?`, StudentNumber)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&Count)
		if err != nil {
			return err
		}

		if Count > 0 {
			Exists = true
		} else {
			Exists = false
		}
	}

	if Exists {
		return fmt.Errorf("Student Number should be unique")
	}

	// create User
	res, err := db.Exec(`INSERT INTO Users (
		Username,
		Password,
		IsAdmin,
		PlainPassword
	)
	Values(?,?,?,?)`,
		StudentNumber,
		Password,
		false,
		plainpass,
	)
	if err != nil {
		return err
	}

	UserID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	_, err = db.Exec(`INSERT INTO Students (
		StudentNumber,
		FirstName,
		LastName,
		MiddleName,
		UserID,
		Email,
		DateOfBirth,
		GradeLevel,
		ContactNumber,
		Address
	)
	Values(?,?,?,?,?,?,?,?,?,?)`,
		StudentNumber,
		FirstName,
		LastName,
		MiddleName,
		UserID,
		Email,
		DateOfBirth,
		GradeLevel,
		ContactNum,
		Address,
	)

	if err != nil {
		return err
	}

	return nil
}

func ReadStudent(db sqlx.Ext, studentNumber int) ([]*ReadStudentResponse, error) {

	var username int64
	var password string
	var studentnum int64
	var userID int64
	var firstname string
	var lastname string
	var middlename string
	var email string
	var dateofbirth string
	var gradelevel string
	var contactnum string
	var address string

	readStdntResponse := make([]*ReadStudentResponse, 0)

	if studentNumber == 0 {

		rows, err := db.Queryx(`SELECT u.Username, u.PlainPassword, s.StudentNumber, s.UserID, s.FirstName, s.LastName, s.MiddleName, s.Email, s.DateOfBirth, s.GradeLevel, s.ContactNumber,s.Address FROM Users as u
	JOIN Students as s
	ON s.UserID = u.ID`)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			err = rows.Scan(&username, &password, &studentnum, &userID, &firstname, &lastname, &middlename, &email, &dateofbirth, &gradelevel, &contactnum, &address)
			if err != nil {
				return nil, err
			}
			readStdntResponse = append(readStdntResponse, &ReadStudentResponse{
				Username:      username,
				Password:      password,
				StudentNumber: studentnum,
				UserID:        int(userID),
				FirstName:     firstname,
				LastName:      lastname,
				MiddleName:    middlename,
				Email:         email,
				DateOfBirth:   dateofbirth,
				GradeLevel:    gradelevel,
				Address:       address,
				ContactNumber: contactnum,
			})
		}

		return readStdntResponse, nil

	}

	rows, err := db.Queryx(`SELECT u.Username, u.Password, s.StudentNumber, s.UserID, s.FirstName, s.LastName, s.MiddleName, s.Email, s.DateOfBirth, s.GradeLevel, s.ContactNumber,s.Address FROM Users as u
	JOIN Students as s
	ON s.UserID = u.ID
	WHERE u.Username = ?`, studentNumber)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&username, &password, &studentnum, &userID, &firstname, &lastname, &middlename, &email, &dateofbirth, &gradelevel, &contactnum, &address)
		if err != nil {
			return nil, err
		}

		readStdntResponse = append(readStdntResponse, &ReadStudentResponse{
			Username:      username,
			StudentNumber: studentnum,
			UserID:        int(userID),
			FirstName:     firstname,
			LastName:      lastname,
			MiddleName:    middlename,
			Email:         email,
			DateOfBirth:   dateofbirth,
			GradeLevel:    gradelevel,
			Address:       address,
			ContactNumber: contactnum,
		})
	}

	return readStdntResponse, nil
}

func UpdateStudent(db sqlx.Ext, studentID int64, userID int64, firstname string, lastname string, middlename string, email string, dateofbirth string, gradelevel string, contactnum string, address string) error {

	if studentID != 0 {

		// check if studentNumber exists
		var Count int64
		Exists := false
		rows, err := db.Queryx(`SELECT COUNT(Username) FROM Users WHERE Username = ?`, studentID)
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			err = rows.Scan(&Count)
			if err != nil {
				return err
			}

			if Count > 0 {
				Exists = true
			} else {
				Exists = false
			}
		}

		if Exists {
			return fmt.Errorf("Student Number should be unique")
		}

		_, err = db.Exec(`UPDATE Users SET Username = ? WHERE ID= ?`,
			studentID,
			userID,
		)
		if err != nil {
			return err
		}

		_, err = db.Exec(`UPDATE Students SET StudentNumber = ? ,FirstName= ?, LastName = ?, MiddleName = ?, Email = ?, DateOfBirth =? , GradeLevel = ?, ContactNumber = ? , Address = ? WHERE UserID= ?`,
			studentID,
			firstname,
			lastname,
			middlename,
			email,
			dateofbirth,
			gradelevel,
			contactnum,
			address,
			userID,
		)

		if err != nil {
			return err
		}

		return nil
	}

	_, err := db.Exec(`UPDATE Students SET FirstName= ?, LastName = ?, MiddleName = ?, Email = ?, DateOfBirth =? , GradeLevel = ?, ContactNumber = ? , Address = ? WHERE UserID= ?`,
		firstname,
		lastname,
		middlename,
		email,
		dateofbirth,
		gradelevel,
		contactnum,
		address,
		userID,
	)

	if err != nil {
		return err
	}

	return nil
}

func DeleteStudent(db sqlx.Ext, userID int64) error {

	_, err := db.Exec(`DELETE FROM Users WHERE ID = ? `, userID)

	if err != nil {
		return err
	}

	_, err = db.Exec(`DELETE FROM Students WHERE UserID = ? `, userID)

	if err != nil {
		return err
	}

	return nil
}
