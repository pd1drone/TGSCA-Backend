package database

import (
	"github.com/jmoiron/sqlx"
)

type ReadTeachersTable struct {
	ID          int64  `json:"ID"`
	TeacherName string `json:"TeacherName"`
}

func CreateTeachers(db sqlx.Ext, TeacherName string) error {

	_, err := db.Exec(`INSERT INTO Teachers (
		TeacherName
	)
	Values(?)`,
		TeacherName,
	)

	if err != nil {
		return err
	}

	return nil
}

func ReadTeachers(db sqlx.Ext) ([]*ReadTeachersTable, error) {

	var id int64
	var teachername string

	Teachers := make([]*ReadTeachersTable, 0)

	rows, err := db.Queryx(`SELECT ID, TeacherName FROM Teachers`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id, &teachername)
		if err != nil {
			return nil, err
		}

		Teachers = append(Teachers, &ReadTeachersTable{
			ID:          id,
			TeacherName: teachername,
		})
	}

	return Teachers, nil
}

func UpdateTeachers(db sqlx.Ext, ID int64, TeacherName string) error {

	_, err := db.Exec(`UPDATE Teachers SET TeacherName= ? WHERE ID= ?`,
		TeacherName,
		ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func DeleteTeachers(db sqlx.Ext, ID int64) error {

	_, err := db.Exec(`DELETE FROM Teachers WHERE ID = ? `, ID)

	if err != nil {
		return err
	}

	return nil
}
