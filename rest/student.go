package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"tgsca/database"
)

type CreateStudentRequest struct {
	StudentNumber int    `json:"StudentNumber"`
	FirstName     string `json:"FirstName"`
	LastName      string `json:"LastName"`
	MiddleName    string `json:"MiddleName"`
	Email         string `json:"Email"`
	DateOfBirth   string `json:"DateOfBirth"`
	GradeLevel    string `json:"GradeLevel"`
	Address       string `json:"Address"`
	ContactNumber string `json:"ContactNumber"`
}

type CreateStudentResponse struct {
	Successful bool   `json:"Successful"`
	Message    string `json:"Message"`
}

type ReadStudentRequest struct {
	StudentNumber int   `json:"StudentNumber"`
	UserID        int64 `json:"UserID"`
}
type ReadStudentResponse struct {
	Username      int    `json:"Username,omitempty"`
	Password      string `json:"Password,omitempty"`
	StudentNumber int    `json:"StudentNumber"`
	UserID        int    `json:"UserID"`
	Name          string `json:"Name"`
	FirstName     string `json:"FirstName"`
	MiddleName    string `json:"MiddleName"`
	LastName      string `json:"LastName"`
	Email         string `json:"Email"`
	DateOfBirth   string `json:"DateOfBirth"`
	GradeLevel    string `json:"GradeLevel"`
	Address       string `json:"Address"`
	ContactNumber string `json:"ContactNumber"`
}

type UpdateStudentRequest struct {
	StudentNumber int64  `json:"StudentNumber"`
	UserID        int64  `json:"UserID"`
	FirstName     string `json:"FirstName"`
	MiddleName    string `json:"MiddleName"`
	LastName      string `json:"LastName"`
	Email         string `json:"Email"`
	DateOfBirth   string `json:"DateOfBirth"`
	GradeLevel    string `json:"GradeLevel"`
	Address       string `json:"Address"`
	ContactNumber string `json:"ContactNumber"`
}

type UpdateStudentResponse struct {
	Successful bool   `json:"Successful"`
	Message    string `json:"Message"`
}

type DeleteStudentRequest struct {
	UserID int64 `json:"UserID"`
}

type DeleteStudentResponse struct {
	Successful bool   `json:"Successful"`
	Message    string `json:"Message"`
}

func (t *TGSCAConfiguration) CreateStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondJSON(w, 500, nil)
		return
	}

	// Restore request body after reading
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	req := &CreateStudentRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	generatedPassword, err := GeneratePassword()
	if err != nil {
		respondJSON(w, 500, nil)
	}

	fmt.Println(generatedPassword)

	md5HashPass := MD5HashPassword(generatedPassword)
	fmt.Println(md5HashPass)

	response := &CreateStudentResponse{}

	err = database.CreateStudent(t.TGSCAdb, req.StudentNumber, req.FirstName, req.LastName, req.MiddleName, md5HashPass, req.Email, req.DateOfBirth, req.GradeLevel, req.ContactNumber, req.Address, generatedPassword)
	if err != nil {
		response.Message = fmt.Sprintf(err.Error())
		respondJSON(w, 400, response)
		return
	}

	response.Successful = true

	respondJSON(w, 200, response)
}

func (t *TGSCAConfiguration) ReadStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondJSON(w, 500, nil)
		return
	}

	// Restore request body after reading
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	req := &ReadStudentRequest{}

	fmt.Println(string(body))

	err = json.Unmarshal(body, &req)
	if err != nil {
		fmt.Println(err)
		fmt.Println("ERROR")
		respondJSON(w, 400, nil)
		return
	}

	response := make([]*ReadStudentResponse, 0)

	dbResponse, err := database.ReadStudent(t.TGSCAdb, int(req.UserID))
	if err != nil {
		fmt.Println(err)
		respondJSON(w, 400, nil)
		return
	}

	if req.UserID != 0 {
		singleResponse := &ReadStudentResponse{
			StudentNumber: int(dbResponse[0].StudentNumber),
			UserID:        dbResponse[0].UserID,
			Name:          dbResponse[0].FirstName + " " + dbResponse[0].MiddleName + " " + dbResponse[0].LastName,
			FirstName:     dbResponse[0].FirstName,
			MiddleName:    dbResponse[0].MiddleName,
			LastName:      dbResponse[0].LastName,
			Email:         dbResponse[0].Email,
			DateOfBirth:   dbResponse[0].DateOfBirth,
			GradeLevel:    dbResponse[0].GradeLevel,
			Address:       dbResponse[0].Address,
			ContactNumber: dbResponse[0].ContactNumber,
		}

		respondJSON(w, 200, singleResponse)
		return
	}

	for _, student := range dbResponse {
		response = append(response, &ReadStudentResponse{
			Username:      int(student.StudentNumber),
			Password:      student.Password,
			StudentNumber: int(student.StudentNumber),
			UserID:        student.UserID,
			Name:          student.FirstName + " " + student.MiddleName + " " + student.LastName,
			FirstName:     student.FirstName,
			MiddleName:    student.MiddleName,
			LastName:      student.LastName,
			Email:         student.Email,
			DateOfBirth:   student.DateOfBirth,
			GradeLevel:    student.GradeLevel,
			Address:       student.Address,
			ContactNumber: student.ContactNumber,
		})
	}

	respondJSON(w, 200, response)
}

func (t *TGSCAConfiguration) UpdateStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondJSON(w, 500, nil)
		return
	}

	fmt.Println(string(body))

	// Restore request body after reading
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	req := &UpdateStudentRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	fmt.Println(req)

	response := &UpdateStudentResponse{}

	err = database.UpdateStudent(t.TGSCAdb, req.StudentNumber, req.UserID, req.FirstName, req.LastName, req.MiddleName, req.Email, req.DateOfBirth, req.GradeLevel, req.ContactNumber, req.Address)
	if err != nil {
		fmt.Println(err)
		response.Successful = false
		response.Message = fmt.Sprintf(err.Error())
		respondJSON(w, 400, response)
		return
	}
	response.Successful = true

	respondJSON(w, 200, response)
}

func (t *TGSCAConfiguration) DeleteStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondJSON(w, 500, nil)
		return
	}

	fmt.Println(string(body))

	// Restore request body after reading
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	req := &DeleteStudentRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	fmt.Println(req)

	response := &DeleteStudentResponse{}

	err = database.DeleteStudent(t.TGSCAdb, req.UserID)
	if err != nil {
		fmt.Println(err)
		response.Successful = false
		response.Message = fmt.Sprintf(err.Error())
		respondJSON(w, 400, response)
		return
	}

	response.Successful = true

	respondJSON(w, 200, response)
}
