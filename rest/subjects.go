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

type CreateSubjectRequest struct {
	Subject    string `json:"Subject"`
	GradeLevel string `json:"GradeLevel"`
	Schedule   string `json:"Schedule"`
	TeachersID int64  `json:"TeachersID"`
}

type CreateSubjectResponse struct {
	Successful bool   `json:"Successful"`
	Message    string `json:"Message"`
}

type ReadSubjectForStudentRequest struct {
	UserID  int64  `json:"UserID"`
	Subject string `json:"Subject"`
}

type DeleteSubjectRequest struct {
	ID int64 `json:"ID"`
}

type DeleteSubjectResponse struct {
	Successful bool   `json:"Successful"`
	Message    string `json:"Message"`
}

type UpdateSubjectRequest struct {
	ID         int64  `json:"ID"`
	Subject    string `json:"Subject"`
	GradeLevel string `json:"GradeLevel"`
	Schedule   string `json:"Schedule"`
	TeachersID int64  `json:"TeachersID"`
}

type UpdateSubjectResponse struct {
	Successful bool   `json:"Successful"`
	Message    string `json:"Message"`
}

type ReadSubjectGradeLevelRequest struct {
	GradeLevel string `json:"GradeLevel"`
}

func (t *TGSCAConfiguration) CreateSubject(w http.ResponseWriter, r *http.Request) {

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

	req := &CreateSubjectRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	response := &CreateSubjectResponse{}

	err = database.CreateSubject(t.TGSCAdb, req.Subject, req.GradeLevel, req.Schedule, req.TeachersID)
	if err != nil {
		response.Message = fmt.Sprintf(err.Error())
		respondJSON(w, 400, response)
		return
	}

	response.Successful = true

	respondJSON(w, 200, response)
}

func (t *TGSCAConfiguration) ReadSubject(w http.ResponseWriter, r *http.Request) {

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

	dbResponse, err := database.ReadSubject(t.TGSCAdb)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, dbResponse)
}

func (t *TGSCAConfiguration) ReadSubjectForStudent(w http.ResponseWriter, r *http.Request) {

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

	req := &ReadSubjectForStudentRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	dbResponse, err := database.ReadSubjectForStudent(t.TGSCAdb, req.UserID)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, dbResponse)
}

func (t *TGSCAConfiguration) ReadSubjectSchedule(w http.ResponseWriter, r *http.Request) {

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

	req := &ReadSubjectForStudentRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	dbResponse, err := database.ReadSubjectSchedule(t.TGSCAdb, req.UserID, req.Subject)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, dbResponse)
}

func (t *TGSCAConfiguration) ReadSubjectForStudentDropDown(w http.ResponseWriter, r *http.Request) {

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

	req := &ReadSubjectForStudentRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	dbResponse, err := database.ReadSubjectForStudentDropDown(t.TGSCAdb, req.UserID)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, dbResponse)
}

func (t *TGSCAConfiguration) ReadSubjectGradeLevel(w http.ResponseWriter, r *http.Request) {

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

	req := &ReadSubjectGradeLevelRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	dbResponse, err := database.ReadSubjectGradeLevel(t.TGSCAdb, req.GradeLevel)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, dbResponse)
}

func (t *TGSCAConfiguration) DeleteSubject(w http.ResponseWriter, r *http.Request) {

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

	req := &DeleteSubjectRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	fmt.Println(req)

	response := &DeleteStudentResponse{}

	err = database.DeleteSubject(t.TGSCAdb, req.ID)
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

func (t *TGSCAConfiguration) UpdateSubject(w http.ResponseWriter, r *http.Request) {

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

	req := &UpdateSubjectRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	fmt.Println(req)

	response := &UpdateSubjectResponse{}

	err = database.UpdateSubject(t.TGSCAdb, req.ID, req.Subject, req.GradeLevel, req.Schedule, req.TeachersID)
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
