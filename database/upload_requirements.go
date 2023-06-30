package database

import (
	"github.com/jmoiron/sqlx"
)

type Requirements struct {
	ID              int64  `json:"ID"`
	StudentNumber   int64  `json:"StudentNumber"`
	UploadedFile    string `json:"UploadedFile"`
	RequirementType string `json:"RequirementType"`
}

func CreateRequirements(db sqlx.Ext, filepath string, UserID int64, RequirementType string) error {

	var studentNumber int64
	rows, err := db.Queryx(`SELECT Username FROM Users WHERE ID = ?`, UserID)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&studentNumber)
		if err != nil {
			return err
		}
	}

	// create User
	_, err = db.Exec(`INSERT INTO Requirements (
		StudentNumber,
		UploadedFile,
		RequirementType
	)
	Values(?,?,?)`,
		studentNumber,
		filepath,
		RequirementType,
	)
	if err != nil {
		return err
	}

	return nil
}

func ReadRequirements(db sqlx.Ext, userID int64) ([]*Requirements, error) {

	requirementArray := make([]*Requirements, 0)
	var stdnum int64
	var ID int64
	var uploadedFilePath string
	var requirementType string

	if userID != 0 {
		var studentNumber int64

		rows, err := db.Queryx(`SELECT Username FROM Users WHERE ID = ?`, userID)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			err = rows.Scan(&studentNumber)
			if err != nil {
				return nil, err
			}
		}

		rows, err = db.Queryx(`SELECT ID, StudentNumber, UploadedFile,RequirementType FROM Requirements WHERE StudentNumber = ?`, studentNumber)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			err = rows.Scan(&ID, &stdnum, &uploadedFilePath, &requirementType)
			if err != nil {
				return nil, err
			}
			requirementArray = append(requirementArray, &Requirements{
				ID:              ID,
				StudentNumber:   stdnum,
				UploadedFile:    uploadedFilePath,
				RequirementType: requirementType,
			})
		}

		return requirementArray, nil
	}

	rows, err := db.Queryx(`SELECT ID, StudentNumber, UploadedFile,RequirementType FROM Requirements`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&ID, &stdnum, &uploadedFilePath, &requirementType)
		if err != nil {
			return nil, err
		}
		requirementArray = append(requirementArray, &Requirements{
			ID:              ID,
			StudentNumber:   stdnum,
			UploadedFile:    uploadedFilePath,
			RequirementType: requirementType,
		})
	}

	return requirementArray, nil

}

func DeleteRequirement(db sqlx.Ext, ID int64) error {

	_, err := db.Exec(`DELETE FROM Requirements WHERE ID = ? `, ID)

	if err != nil {
		return err
	}

	return nil
}
