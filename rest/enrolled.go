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

type CreateEnrolledRequest struct {
	StudentNumber int64 `json:"StudentNumber"`
	SubjectID     int64 `json:"SubjectID"`
}

type CreateEnrolledResponse struct {
	Successful bool   `json:"Successful"`
	Message    string `json:"Message"`
}

type DeleteEnrolledRequest struct {
	ID int64 `json:"ID"`
}

type DeleteEnrolledResponse struct {
	Successful bool   `json:"Successful"`
	Message    string `json:"Message"`
}

type UpdateEnrolledRequest struct {
	ID            int64 `json:"ID"`
	StudentNumber int64 `json:"StudentNumber"`
	SubjectID     int64 `json:"SubjectID"`
}

type UpdateEnrolledResponse struct {
	Successful bool   `json:"Successful"`
	Message    string `json:"Message"`
}

func (t *TGSCAConfiguration) CreateEnrolled(w http.ResponseWriter, r *http.Request) {

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

	err = database.CreateEnrolled(t.TGSCAdb, req.StudentNumber, req.SubjectID)
	if err != nil {
		response.Message = fmt.Sprintf(err.Error())
		respondJSON(w, 400, response)
		return
	}

	response.Successful = true

	respondJSON(w, 200, response)
}

func (t *TGSCAConfiguration) ReadEnrolled(w http.ResponseWriter, r *http.Request) {

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

	dbResponse, err := database.ReadEnrolled(t.TGSCAdb)
	if err != nil {
		fmt.Println(err)
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, dbResponse)
}

func (t *TGSCAConfiguration) DeleteEnrolled(w http.ResponseWriter, r *http.Request) {

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

	req := &DeleteEnrolledRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	fmt.Println(req)

	response := &DeleteEnrolledResponse{}

	err = database.DeleteEnrolled(t.TGSCAdb, req.ID)
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

func (t *TGSCAConfiguration) UpdateEnrolled(w http.ResponseWriter, r *http.Request) {

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

	req := &UpdateEnrolledRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	fmt.Println(req)

	response := &UpdateEnrolledResponse{}

	err = database.UpdateEnrolled(t.TGSCAdb, req.ID, req.StudentNumber, req.SubjectID)
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

type ReadEnrolledSubjects struct {
	UserID int64 `json:"UserID"`
}

func (t *TGSCAConfiguration) ReadEnrolledSubjects(w http.ResponseWriter, r *http.Request) {

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

	req := &ReadEnrolledSubjects{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	dbResponse, err := database.ReadEnrolledSubjects(t.TGSCAdb, req.UserID)
	if err != nil {
		fmt.Println(err)
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, dbResponse)
}
