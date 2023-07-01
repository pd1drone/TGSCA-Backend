package database

import "github.com/jmoiron/sqlx"

type Appointment struct {
	ID                     int64  `json:"ID"`
	Name                   string `json:"Name"`
	Email                  string `json:"Email"`
	ContactNumber          string `json:"ContactNumber"`
	StudentNumber          int64  `json:"StudentNumber"`
	AppointmentType        string `json:"AppointmentType"`
	AppointmentDescription string `json:"AppointmentDescription"`
	AppointmentDate        string `json:"AppointmentDate"`
}

func CreateAppointment(db sqlx.Ext, name string, email string, contactnum string, studentnum int64, appointmenttype string, appointmentdesc string, appointmentDate string) error {

	_, err := db.Exec(`INSERT INTO Appointments (
		Name,
		Email,
		ContactNumber,
		StudentNumber,
		AppointmentType,
		AppointmentDescription,
		AppointmentDate
	)
	Values(?,?,?,?,?,?,?)`,
		name,
		email,
		contactnum,
		studentnum,
		appointmenttype,
		appointmentdesc,
		appointmentDate,
	)

	if err != nil {
		return err
	}

	return nil
}

func ReadAppointment(db sqlx.Ext) ([]*Appointment, error) {

	var id int64
	var name string
	var email string
	var contactnum string
	var studentnum int64
	var appointmenttype string
	var appointmentdesc string
	var appointmentDate string

	AppointmentArray := make([]*Appointment, 0)

	rows, err := db.Queryx(`SELECT ID, Name, Email, ContactNumber, StudentNumber,AppointmentType,AppointmentDescription,AppointmentDate FROM Appointments`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id, &name, &email, &contactnum, &studentnum, &appointmenttype, &appointmentdesc, &appointmentDate)
		if err != nil {
			return nil, err
		}

		AppointmentArray = append(AppointmentArray, &Appointment{
			ID:                     id,
			Name:                   name,
			Email:                  email,
			ContactNumber:          contactnum,
			StudentNumber:          studentnum,
			AppointmentType:        appointmenttype,
			AppointmentDescription: appointmentdesc,
			AppointmentDate:        appointmentDate,
		})
	}

	return AppointmentArray, nil
}
