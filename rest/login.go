package rest

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"tgsca/database"
)

type AdminLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AdminLoginResponse struct {
	Authenticated bool   `json:"Authenticated"`
	UserType      string `json:"UserType"`
}

type StudentLoginRequest struct {
	StudentNumber string `json:"studentNumber"`
	DateOfBirth   string `json:"dateOfBirth"`
	Password      string `json:"password"`
}

type StudentLoginResponse struct {
	Authenticated bool   `json:"Authenticated"`
	UserType      string `json:"UserType"`
	UserID        int64  `json:"UserID"`
}

func (t *TGSCAConfiguration) LoginAdmin(w http.ResponseWriter, r *http.Request) {

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

	req := &AdminLoginRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	response := &AdminLoginResponse{}

	LoginAuthorized, err := database.LoginAdmin(t.TGSCAdb, req.Username, req.Password)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	response.Authenticated = LoginAuthorized
	if response.Authenticated {
		response.UserType = "Admin"
	}

	respondJSON(w, 200, response)
}

func (t *TGSCAConfiguration) LoginStudent(w http.ResponseWriter, r *http.Request) {

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

	req := &StudentLoginRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	response := &StudentLoginResponse{}

	LoginAuthorized, userID, err := database.LoginStudent(t.TGSCAdb, req.StudentNumber, req.Password, req.DateOfBirth)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	response.Authenticated = LoginAuthorized
	response.UserID = userID
	if response.Authenticated {
		response.UserType = "Student"
	}

	respondJSON(w, 200, response)
}
