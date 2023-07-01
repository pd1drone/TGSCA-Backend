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

type CreateAppointmentRequest struct {
	Name                   string `json:"Name"`
	Email                  string `json:"Email"`
	ContactNumber          string `json:"ContactNumber"`
	StudentNumber          int64  `json:"StudentNumber"`
	AppointmentType        string `json:"AppointmentType"`
	AppointmentDescription string `json:"AppointmentDescription"`
	AppointmentDate        string `json:"AppointmentDate"`
}

type CreateAppointmentResponse struct {
	Successful bool   `json:"Successful"`
	Message    string `json:"Message"`
}

func (t *TGSCAConfiguration) CreateAppointment(w http.ResponseWriter, r *http.Request) {

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

	req := &CreateAppointmentRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	response := &CreateAppointmentResponse{}

	err = database.CreateAppointment(t.TGSCAdb, req.Name, req.Email, req.ContactNumber, req.StudentNumber, req.AppointmentType, req.AppointmentDescription, req.AppointmentDate)
	if err != nil {
		response.Message = fmt.Sprintf(err.Error())
		respondJSON(w, 400, response)
		return
	}

	response.Successful = true

	respondJSON(w, 200, response)
}

func (t *TGSCAConfiguration) ReadAppointment(w http.ResponseWriter, r *http.Request) {

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

	dbResponse, err := database.ReadAppointment(t.TGSCAdb)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, dbResponse)
}
