package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"path/filepath"
)

func (t *TGSCAConfiguration) UploadRequirements(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	// Parse the JSON request body
	var formData struct {
		File            []byte `json:"file"`
		RequirementType string `json:"requirementType"`
		FileName        string `json:"fileName"`
		UserID          int64  `json:"userID"`
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondJSON(w, 500, nil)
		return
	}

	// Restore request body after reading
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	err = json.Unmarshal(body, &formData)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	// Construct the file path
	filePath := fmt.Sprintf("/root/TGSCA-Backend/files/%d/%s", formData.UserID, formData.FileName)

	// Save the file to disk with the provided file path
	err = ioutil.WriteFile(filePath, formData.File, 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("File uploaded successfully"))
}

func (t *TGSCAConfiguration) ServeFile(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")
	// Get the UserID and filename from the URL parameters
	userID := 8
	filename := "Arduino-Mega-Pinout.jpg"

	// Construct the file path
	filePath := fmt.Sprintf("/root/TGSCA-Backend/files/%s/%s", userID, filename)

	// Read the file content
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the appropriate content type based on the file extension
	contentType := mime.TypeByExtension(filepath.Ext(filePath))
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	w.Header().Set("Content-Type", contentType)

	// Write the file content as the response body
	w.Write(fileContent)
}
