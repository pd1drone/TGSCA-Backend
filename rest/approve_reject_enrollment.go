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

func (t *TGSCAConfiguration) GetStudentEnrollmentStatus(w http.ResponseWriter, r *http.Request) {

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

	dbResponse, err := database.GetStudentEnrollmentStatus(t.TGSCAdb)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, dbResponse)
}

type ApproveStudentEnrollementRequest struct {
	StudentNumber int64 `json:"StudentNumber"`
}
type ApproveStudentEnrollementResponse struct {
	Successful bool   `json:"Successful"`
	Message    string `json:"Message"`
}

func (t *TGSCAConfiguration) ApproveStudentEnrollment(w http.ResponseWriter, r *http.Request) {

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

	req := &ApproveStudentEnrollementRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	response := &ApproveStudentEnrollementResponse{}

	err = database.ApproveStudentEnrollment(t.TGSCAdb, req.StudentNumber)

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
