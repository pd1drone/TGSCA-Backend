package database

import "github.com/jmoiron/sqlx"

type Subject struct {
	ID          int64  `json:"ID"`
	Subject     string `json:"Subject"`
	GradeLevel  string `json:"GradeLevel"`
	Schedule    string `json:"Schedule"`
	TeacherName string `json:"TeacherName"`
}

func CreateSubject(db sqlx.Ext, subject string, gradelevel string, schedule string, teachersID int64) error {

	_, err := db.Exec(`INSERT INTO Subjects (
		Subject,
		GradeLevel,
		Schedule,
		TeachersID
	)
	Values(?,?,?,?)`,
		subject,
		gradelevel,
		schedule,
		teachersID,
	)

	if err != nil {
		return err
	}

	return nil
}

func ReadSubject(db sqlx.Ext) ([]*Subject, error) {

	var id int64
	var subject string
	var gradelevel string
	var sched string
	var teachersID int64

	SubjectsArray := make([]*Subject, 0)

	rows, err := db.Queryx(`SELECT ID,Subject,GradeLevel,Schedule,TeachersID FROM Subjects`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id, &subject, &gradelevel, &sched, &teachersID)
		if err != nil {
			return nil, err
		}

		teacherName, err := GetTeachersName(db, teachersID)
		if err != nil {
			return nil, err
		}

		SubjectsArray = append(SubjectsArray, &Subject{
			ID:          id,
			Subject:     subject,
			GradeLevel:  gradelevel,
			Schedule:    sched,
			TeacherName: teacherName,
		})
	}

	return SubjectsArray, nil
}

func ReadSubjectForStudent(db sqlx.Ext, userID int64) ([]*Subject, error) {

	var id int64
	var subject string
	var gradelevel string
	var sched string
	var teachersID int64

	SubjectsArray := make([]*Subject, 0)

	GradeLVL, err := GetStudentGradeLevel(db, userID)
	if err != nil {
		return nil, err
	}

	rows, err := db.Queryx(`SELECT ID,Subject,GradeLevel,Schedule,TeachersID FROM Subjects WHERE GradeLevel= ?`, GradeLVL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id, &subject, &gradelevel, &sched, &teachersID)
		if err != nil {
			return nil, err
		}

		teacherName, err := GetTeachersName(db, teachersID)
		if err != nil {
			return nil, err
		}

		SubjectsArray = append(SubjectsArray, &Subject{
			ID:          id,
			Subject:     subject,
			GradeLevel:  gradelevel,
			Schedule:    sched,
			TeacherName: teacherName,
		})
	}

	return SubjectsArray, nil
}

func ReadSubjectGradeLevel(db sqlx.Ext, GradeLevel string) ([]*Subject, error) {

	var id int64
	var subject string
	var gradelevel string
	var sched string
	var teachersID int64

	SubjectsArray := make([]*Subject, 0)

	rows, err := db.Queryx(`SELECT ID,Subject,GradeLevel,Schedule,TeachersID FROM Subjects WHERE GradeLevel= ?`, GradeLevel)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id, &subject, &gradelevel, &sched, &teachersID)
		if err != nil {
			return nil, err
		}

		teacherName, err := GetTeachersName(db, teachersID)
		if err != nil {
			return nil, err
		}

		SubjectsArray = append(SubjectsArray, &Subject{
			ID:          id,
			Subject:     subject,
			GradeLevel:  gradelevel,
			Schedule:    sched,
			TeacherName: teacherName,
		})
	}

	return SubjectsArray, nil
}

func GetStudentGradeLevel(db sqlx.Ext, userID int64) (string, error) {

	var gradelvl string
	rows, err := db.Queryx(`SELECT GradeLevel FROM Students WHERE UserID= ?`, userID)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&gradelvl)
		if err != nil {
			return "", err
		}
	}

	return gradelvl, nil
}

func GetTeachersName(db sqlx.Ext, teacherID int64) (string, error) {
	var teachersName string
	rows, err := db.Queryx(`SELECT TeacherName FROM Teachers WHERE ID= ?`, teacherID)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&teachersName)
		if err != nil {
			return "", err
		}
	}
	return teachersName, nil
}

func DeleteSubject(db sqlx.Ext, subID int64) error {

	_, err := db.Exec(`DELETE FROM Subjects WHERE ID = ? `, subID)

	if err != nil {
		return err
	}

	return nil
}

func UpdateSubject(db sqlx.Ext, id int64, subject string, gradelevel string, sched string, teacherid int64) error {

	_, err := db.Exec(`UPDATE Subjects SET Subject= ?,GradeLevel=?,Schedule=?,TeachersID= ? WHERE ID= ?`,
		subject,
		gradelevel,
		sched,
		teacherid,
		id,
	)

	if err != nil {
		return err
	}

	return nil
}
