package database

import (
	"github.com/jmoiron/sqlx"
)

type Enrolled struct {
	ID            int64 `json:"ID"`
	StudentNumber int64 `json:"StudentNumber"`
	SubjectID     int64 `json:"SubjectID"`
}

type ReadEnrolledResponse struct {
	ID            int64  `json:"ID"`
	StudentNumber int64  `json:"StudentNumber"`
	StudentName   string `json:"StudentName"`
	GradeLevel    string `json:"GradeLevel"`
	Subject       string `json:"Subject"`
	Schedule      string `json:"Schedule"`
	TeacherName   string `json:"TeacherName"`
	SubjectID     int64  `json:"SubjectID"`
}

func CreateEnrolled(db sqlx.Ext, studentNumber int64, SubjectID int64) error {

	_, err := db.Exec(`INSERT INTO Enrolled (
		StudentNumber,
		SubjectID
	)
	Values(?,?)`,
		studentNumber,
		SubjectID,
	)

	if err != nil {
		return err
	}

	return nil
}

func ReadEnrolled(db sqlx.Ext) ([]*ReadEnrolledResponse, error) {

	var id int64
	var stdNumber int64
	var subID int64

	EnrolledArr := make([]*Enrolled, 0)
	EnrolledDetailsArray := make([]*ReadEnrolledResponse, 0)

	rows, err := db.Queryx(`SELECT ID, StudentNumber,SubjectID FROM Enrolled`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id, &stdNumber, &subID)
		if err != nil {
			return nil, err
		}

		EnrolledArr = append(EnrolledArr, &Enrolled{
			ID:            id,
			StudentNumber: stdNumber,
			SubjectID:     subID,
		})
	}

	for _, enrollee := range EnrolledArr {

		enrolleeDetails, err := GetEnrolleeDetails(db, enrollee.StudentNumber, enrollee.SubjectID)
		if err != nil {
			return nil, err
		}
		EnrolledDetailsArray = append(EnrolledDetailsArray, &ReadEnrolledResponse{
			ID:            enrollee.ID,
			StudentNumber: enrolleeDetails.StudentNumber,
			StudentName:   enrolleeDetails.StudentName,
			GradeLevel:    enrolleeDetails.GradeLevel,
			Subject:       enrolleeDetails.Subject,
			Schedule:      enrolleeDetails.Schedule,
			TeacherName:   enrolleeDetails.TeacherName,
			SubjectID:     enrollee.SubjectID,
		})
	}

	return EnrolledDetailsArray, nil
}

func GetEnrolleeDetails(db sqlx.Ext, stdnumber int64, subjectID int64) (*ReadEnrolledResponse, error) {

	var subject string
	var gradelvl string
	var schedule string
	var teachername string

	rows, err := db.Queryx(`SELECT sub.Subject, sub.GradeLevel, sub.Schedule, t.TeacherName FROM Subjects as sub
	JOIN Teachers as t
	ON sub.TeachersID = t.ID
	WHERE sub.ID = ?`, subjectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&subject, &gradelvl, &schedule, &teachername)
		if err != nil {
			return nil, err
		}
	}

	var studentnumber int64
	var firstname string
	var lastname string
	var middlename string

	rows, err = db.Queryx(`SELECT StudentNumber, FirstName,LastName,MiddleName FROM Students Where StudentNumber = ?`, stdnumber)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&studentnumber, &firstname, &lastname, &middlename)
		if err != nil {
			return nil, err
		}
	}

	return &ReadEnrolledResponse{
		ID:            0,
		StudentNumber: studentnumber,
		StudentName:   firstname + " " + middlename + " " + lastname,
		GradeLevel:    gradelvl,
		Subject:       subject,
		Schedule:      schedule,
		TeacherName:   teachername,
	}, nil

}

func DeleteEnrolled(db sqlx.Ext, enrolledID int64) error {

	_, err := db.Exec(`DELETE FROM Enrolled WHERE ID = ? `, enrolledID)

	if err != nil {
		return err
	}

	return nil
}

func UpdateEnrolled(db sqlx.Ext, id int64, studentnumber int64, subjectID int64) error {

	_, err := db.Exec(`UPDATE Enrolled SET StudentNumber= ?,SubjectID= ? WHERE ID= ?`,
		studentnumber,
		subjectID,
		id,
	)

	if err != nil {
		return err
	}

	return nil
}

func ReadEnrolledSubjects(db sqlx.Ext, userID int64) ([]*ReadEnrolledResponse, error) {

	var studentnum int64

	rows, err := db.Queryx(`SELECT Username FROM Users WHERE ID = ?`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&studentnum)
		if err != nil {
			return nil, err
		}
	}

	var id int64
	var subID int64

	EnrolledArr := make([]*Enrolled, 0)
	EnrolledDetailsArray := make([]*ReadEnrolledResponse, 0)

	rows, err = db.Queryx(`SELECT ID,SubjectID FROM Enrolled WHERE StudentNumber = ?`, studentnum)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id, &subID)
		if err != nil {
			return nil, err
		}

		EnrolledArr = append(EnrolledArr, &Enrolled{
			ID:        id,
			SubjectID: subID,
		})
	}

	for _, enrollee := range EnrolledArr {

		enrolleeDetails, err := GetEnrolleeDetails(db, studentnum, enrollee.SubjectID)
		if err != nil {
			return nil, err
		}
		EnrolledDetailsArray = append(EnrolledDetailsArray, &ReadEnrolledResponse{
			ID:            enrollee.ID,
			StudentNumber: enrolleeDetails.StudentNumber,
			StudentName:   enrolleeDetails.StudentName,
			GradeLevel:    enrolleeDetails.GradeLevel,
			Subject:       enrolleeDetails.Subject,
			Schedule:      enrolleeDetails.Schedule,
			TeacherName:   enrolleeDetails.TeacherName,
			SubjectID:     enrollee.SubjectID,
		})
	}

	return EnrolledDetailsArray, nil
}
