package database

import "github.com/jmoiron/sqlx"

type EnrolledPending struct {
	ID            int64 `json:"ID"`
	StudentNumber int64 `json:"StudentNumber"`
	SubjectID     int64 `json:"SubjectID"`
}

type ReadEnrolledPendingResponse struct {
	ID            int64  `json:"ID"`
	StudentNumber int64  `json:"StudentNumber"`
	StudentName   string `json:"StudentName"`
	GradeLevel    string `json:"GradeLevel"`
	Subject       string `json:"Subject"`
	Schedule      string `json:"Schedule"`
	TeacherName   string `json:"TeacherName"`
}

func CreateEnrolledPending(db sqlx.Ext, studentNumber int64, SubjectID int64) error {

	_, err := db.Exec(`INSERT INTO EnrolledPending (
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

func ReadEnrolledPending(db sqlx.Ext, userID int64) ([]*ReadEnrolledPendingResponse, error) {

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
	var stdNumber int64
	var subID int64

	EnrolledArr := make([]*EnrolledPending, 0)
	EnrolledDetailsArray := make([]*ReadEnrolledPendingResponse, 0)

	rows, err = db.Queryx(`SELECT ID, StudentNumber,SubjectID FROM EnrolledPending WHERE StudentNumber = ?`, studentnum)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id, &stdNumber, &subID)
		if err != nil {
			return nil, err
		}

		EnrolledArr = append(EnrolledArr, &EnrolledPending{
			ID:            id,
			StudentNumber: stdNumber,
			SubjectID:     subID,
		})
	}

	for _, enrollee := range EnrolledArr {

		enrolleeDetails, err := GetEnrolleeDetailsPending(db, enrollee.StudentNumber, enrollee.SubjectID)
		if err != nil {
			return nil, err
		}
		EnrolledDetailsArray = append(EnrolledDetailsArray, &ReadEnrolledPendingResponse{
			ID:            enrollee.ID,
			StudentNumber: enrolleeDetails.StudentNumber,
			StudentName:   enrolleeDetails.StudentName,
			GradeLevel:    enrolleeDetails.GradeLevel,
			Subject:       enrolleeDetails.Subject,
			Schedule:      enrolleeDetails.Schedule,
			TeacherName:   enrolleeDetails.TeacherName,
		})
	}

	return EnrolledDetailsArray, nil
}

func GetEnrolleeDetailsPending(db sqlx.Ext, stdnumber int64, subjectID int64) (*ReadEnrolledPendingResponse, error) {

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

	return &ReadEnrolledPendingResponse{
		ID:            0,
		StudentNumber: studentnumber,
		StudentName:   firstname + " " + middlename + " " + lastname,
		GradeLevel:    gradelvl,
		Subject:       subject,
		Schedule:      schedule,
		TeacherName:   teachername,
	}, nil

}

func DeleteEnrolledPending(db sqlx.Ext, enrolledID int64) error {

	_, err := db.Exec(`DELETE FROM EnrolledPending WHERE ID = ? `, enrolledID)

	if err != nil {
		return err
	}

	return nil
}

func UpdateEnrolledPending(db sqlx.Ext, id int64, studentnumber int64, subjectID int64) error {

	_, err := db.Exec(`UPDATE EnrolledPending SET StudentNumber= ?,SubjectID= ? WHERE ID= ?`,
		studentnumber,
		subjectID,
		id,
	)

	if err != nil {
		return err
	}

	return nil
}
