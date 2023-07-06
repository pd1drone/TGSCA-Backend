package database

import "github.com/jmoiron/sqlx"

type StudentNumberID struct {
	StudentNumber int64 `json:"StudentNumber"`
}

type ApproveRejectResponse struct {
	StudentNumber      int64  `json:"StudentNumber"`
	StudentName        string `json:"StudentName"`
	GradeLevel         string `json:"GradeLevel"`
	ProgressCard       string `json:"ProgressCard"`
	ProgressCardStatus string `json:"ProgressCardStatus"`
	Form137            string `json:"Form137"`
	Form137Status      string `json:"Form137Status"`
	GoodMoral          string `json:"GoodMoral"`
	GoodMoralStatus    string `json:"GoodMoralStatus"`
	RegFee             string `json:"RegFee"`
	RegFeeStatus       string `json:"RegFeeStatus"`
}

func GetStudentEnrollmentStatus(db sqlx.Ext) ([]*ApproveRejectResponse, error) {

	userIds := make([]*StudentNumberID, 0)
	var StudentNumber int64

	rows, err := db.Queryx(`SELECT DISTINCT StudentNumber FROM EnrolledPending`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&StudentNumber)
		if err != nil {
			return nil, err
		}
		userIds = append(userIds, &StudentNumberID{
			StudentNumber: StudentNumber,
		})
	}

	approveRejectArr := make([]*ApproveRejectResponse, 0)
	for _, student := range userIds {
		var fname string
		var lname string
		var mname string
		var gradelvl string

		rowsSTD, err := db.Queryx(`SELECT FirstName,LastName, MiddleName, GradeLevel FROM Students WHERE StudentNumber= ?`, student.StudentNumber)

		if err != nil {
			return nil, err
		}
		defer rowsSTD.Close()

		for rowsSTD.Next() {
			err = rowsSTD.Scan(&fname, &lname, &mname, &gradelvl)
			if err != nil {
				return nil, err
			}
		}

		var progresscard string
		var form137 string
		var goodmoral string
		var regfee string

		rowsForm, err := db.Queryx(`SELECT UploadedFile FROM Requirements WHERE StudentNumber= ? AND RequirementType= ?`, student.StudentNumber, "Form 137")

		if err != nil {
			return nil, err
		}
		defer rowsForm.Close()

		for rowsForm.Next() {
			err = rowsForm.Scan(&form137)
			if err != nil {
				return nil, err
			}
		}

		rowsCert, err := db.Queryx(`SELECT UploadedFile FROM Requirements WHERE StudentNumber= ? AND RequirementType= ?`, student.StudentNumber, "Good Moral Certificate")

		if err != nil {
			return nil, err
		}
		defer rowsCert.Close()

		for rowsCert.Next() {
			err = rowsCert.Scan(&goodmoral)
			if err != nil {
				return nil, err
			}
		}

		rowsCard, err := db.Queryx(`SELECT UploadedFile FROM Requirements WHERE StudentNumber= ? AND RequirementType= ?`, student.StudentNumber, "Progress Card")

		if err != nil {
			return nil, err
		}
		defer rowsCard.Close()

		for rowsCard.Next() {
			err = rowsCard.Scan(&progresscard)
			if err != nil {
				return nil, err
			}
		}

		rowsFee, err := db.Queryx(`SELECT UploadedFile FROM Requirements WHERE StudentNumber= ? AND RequirementType= ?`, student.StudentNumber, "Registration Fee")

		if err != nil {
			return nil, err
		}
		defer rowsFee.Close()

		for rowsFee.Next() {
			err = rowsFee.Scan(&regfee)
			if err != nil {
				return nil, err
			}
		}

		progresscardStatus := "Not Uploaded"
		form137Status := "Not Uploaded"
		goodmoralStatus := "Not Uploaded"
		regfeeStatus := "Not Uploaded"

		if progresscard != "" {
			progresscardStatus = "Uploaded"
		}
		if form137 != "" {
			form137Status = "Uploaded"
		}
		if goodmoral != "" {
			goodmoralStatus = "Uploaded"
		}
		if regfee != "" {
			regfeeStatus = "Uploaded"
		}

		approveRejectArr = append(approveRejectArr, &ApproveRejectResponse{
			StudentNumber:      student.StudentNumber,
			StudentName:        fname + " " + mname + " " + lname,
			GradeLevel:         gradelvl,
			ProgressCard:       progresscard,
			ProgressCardStatus: progresscardStatus,
			Form137:            form137,
			Form137Status:      form137Status,
			GoodMoral:          goodmoral,
			GoodMoralStatus:    goodmoralStatus,
			RegFee:             regfee,
			RegFeeStatus:       regfeeStatus,
		})
	}

	return approveRejectArr, nil
}

func ApproveStudentEnrollment(db sqlx.Ext, studentNumber int64) error {

	var id int64
	var stdNumber int64
	var subID int64

	EnrolledArr := make([]*EnrolledPending, 0)

	rows, err := db.Queryx(`SELECT ID, StudentNumber,SubjectID FROM EnrolledPending WHERE StudentNumber = ?`, studentNumber)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id, &stdNumber, &subID)
		if err != nil {
			return err
		}

		EnrolledArr = append(EnrolledArr, &EnrolledPending{
			ID:            id,
			StudentNumber: stdNumber,
			SubjectID:     subID,
		})
	}

	for _, subjectEnrolled := range EnrolledArr {
		var count int64
		Exists := false

		rows, err := db.Queryx(`SELECT COUNT(ID) FROM Enrolled WHERE StudentNumber = ? And SubjectID = ?`, subjectEnrolled.StudentNumber, subjectEnrolled.SubjectID)
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			err = rows.Scan(&count)
			if err != nil {
				return err
			}

			if count > 0 {
				Exists = true
			} else {
				Exists = false
			}
		}

		if !Exists {
			err = CreateEnrolled(db, subjectEnrolled.StudentNumber, subjectEnrolled.SubjectID)
			if err != nil {
				return err
			}
		}

		err = DeleteEnrolledPending(db, subjectEnrolled.ID)
		if err != nil {
			return err
		}

	}

	return nil
}
