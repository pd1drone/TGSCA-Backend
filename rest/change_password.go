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

type ChangePasswordRequest struct {
	Username    int64  `json:"Username"`
	Password    string `json:"Password"`
	NewPassword string `json:"NewPassword"`
	UserID      int64  `json:"UserID"`
}

type ChangePasswordResponse struct {
	Successful bool   `json:"Successful"`
	Message    string `json:"Message"`
}

func (t *TGSCAConfiguration) ChangePassword(w http.ResponseWriter, r *http.Request) {

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

	req := &ChangePasswordRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	response := &ChangePasswordResponse{}

	err = database.ChangePassword(t.TGSCAdb, req.UserID, req.Username, req.Password, req.NewPassword, MD5HashPassword(req.NewPassword))
	if err != nil {
		response.Successful = false
		response.Message = fmt.Sprintf(err.Error())
		respondJSON(w, 400, response)
		return
	}

	response.Successful = true

	respondJSON(w, 200, response)
}
