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

type ReadEnrolledStudentRequest struct {
	UserID int64 `json:"UserID"`
}
type CreateEnrolledStudentRequest struct {
	StudentNumber int64 `json:"StudentNumber"`
	SubjectID     int64 `json:"SubjectID"`
}

type CreateEnrolledStudentResponse struct {
	Successful bool   `json:"Successful"`
	Message    string `json:"Message"`
}

type DeleteEnrolledStudentRequest struct {
	ID int64 `json:"ID"`
}

type DeleteEnrolledStudentResponse struct {
	Successful bool   `json:"Successful"`
	Message    string `json:"Message"`
}

type UpdateEnrolledStudentRequest struct {
	ID            int64 `json:"ID"`
	StudentNumber int64 `json:"StudentNumber"`
	SubjectID     int64 `json:"SubjectID"`
}

type UpdateEnrolledStudentResponse struct {
	Successful bool   `json:"Successful"`
	Message    string `json:"Message"`
}

func (t *TGSCAConfiguration) CreateEnrolledForStudent(w http.ResponseWriter, r *http.Request) {

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

	req := &CreateEnrolledRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	response := &CreateEnrolledResponse{}

	err = database.CreateEnrolledPending(t.TGSCAdb, req.StudentNumber, req.SubjectID)
	if err != nil {
		response.Message = fmt.Sprintf(err.Error())
		respondJSON(w, 400, response)
		return
	}

	response.Successful = true

	respondJSON(w, 200, response)
}

func (t *TGSCAConfiguration) ReadEnrolledForStudent(w http.ResponseWriter, r *http.Request) {

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

	req := &ReadEnrolledStudentRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	dbResponse, err := database.ReadEnrolledPending(t.TGSCAdb, req.UserID)
	if err != nil {
		fmt.Println(err)
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, dbResponse)
}

func (t *TGSCAConfiguration) DeleteEnrolledForStudent(w http.ResponseWriter, r *http.Request) {

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

	req := &DeleteEnrolledStudentRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	fmt.Println(req)

	response := &DeleteEnrolledStudentResponse{}

	err = database.DeleteEnrolledPending(t.TGSCAdb, req.ID)
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

func (t *TGSCAConfiguration) UpdateEnrolledForStudent(w http.ResponseWriter, r *http.Request) {

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

	req := &UpdateEnrolledStudentRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	fmt.Println(req)

	response := &UpdateEnrolledStudentResponse{}

	err = database.UpdateEnrolledPending(t.TGSCAdb, req.ID, req.StudentNumber, req.SubjectID)
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
