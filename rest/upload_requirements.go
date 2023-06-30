package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"tgsca/database"

	"github.com/go-chi/chi"
)

type ReadRequirementsRequest struct {
	UserID int64 `json:"UserID,omitempty"`
}

type ReadRequirementsResponse struct {
	ID              int64  `json:"ID"`
	StudentNumber   int64  `json:"StudentNumber"`
	UploadedFile    string `json:"UploadedFile"`
	RequirementType string `json:"RequirementType"`
}

type DeleteRequirementsRequest struct {
	ID              int64  `json:"ID"`
	StudentNumber   int64  `json:"StudentNumber"`
	UploadedFile    string `json:"UploadedFile"`
	RequirementType string `json:"RequirementType"`
}

type DeleteRequirementsResponse struct {
	Successful bool   `json:"Successful"`
	Message    string `json:"Message"`
}

func (t *TGSCAConfiguration) UploadRequirements(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	// Parse our multipart form, 32 << 20 specifies a maximum
	// upload of 32 MB files.
	r.ParseMultipartForm(32 << 20)

	file, handler, err := r.FormFile("uploadRequirement")
	if err != nil {
		w.Write([]byte("Error in uploading file!"))
		respondJSON(w, 400, nil)
		return
	}
	defer file.Close()

	UserID := r.FormValue("UserID")
	fmt.Println(UserID)
	requirementType := r.FormValue("requirementType")

	err = createDirectoryIfNotExist("/root/TGSCA-Backend/files/" + UserID)
	if err != nil {
		w.Write([]byte("Error in uploading file!"))
		respondJSON(w, 500, nil)
		return
	}

	filePath := fmt.Sprintf("/root/TGSCA-Backend/files/%s/%s", UserID, handler.Filename)

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		w.Write([]byte("Error in uploading file!"))
		respondJSON(w, 500, nil)
		return
	}
	// Save the file to disk with the provided file path
	err = os.WriteFile(filePath, fileBytes, 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	intUserID, err := strconv.Atoi(UserID)
	if err != nil {
		w.Write([]byte("Error in uploading file!"))
		respondJSON(w, 500, nil)
		return
	}

	err = database.CreateRequirements(t.TGSCAdb, filePath, int64(intUserID), requirementType)
	if err != nil {
		w.Write([]byte("Error in uploading file!"))
		respondJSON(w, 500, nil)
		return
	}

	// Return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")

}

func (t *TGSCAConfiguration) ServeFile(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	userID := chi.URLParam(r, "userID")
	filename := chi.URLParam(r, "filename")

	// Use the retrieved userID and filename as needed
	fmt.Println("userID:", userID)
	fmt.Println("filename:", filename)

	// Construct the file path
	filePath := fmt.Sprintf("/root/TGSCA-Backend/files/%s/%s", userID, filename)

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Set the appropriate Content-Type header
	contentType := mime.TypeByExtension(filepath.Ext(filePath))
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	w.Header().Set("Content-Type", contentType)

	// Copy the file contents to the response writer
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}
}

func createDirectoryIfNotExist(directoryPath string) error {
	// Check if the directory already exists
	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		// Directory does not exist, create it
		err := os.MkdirAll(directoryPath, 0755)
		if err != nil {
			return fmt.Errorf("error creating directory: %w", err)
		}
		fmt.Println("Directory created:", directoryPath)
	} else if err != nil {
		return fmt.Errorf("error checking directory: %w", err)
	} else {
		fmt.Println("Directory already exists:", directoryPath)
	}

	return nil
}

func (t *TGSCAConfiguration) ReadRequirements(w http.ResponseWriter, r *http.Request) {

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

	req := &ReadRequirementsRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	response := make([]*ReadRequirementsResponse, 0)

	dbResponse, err := database.ReadRequirements(t.TGSCAdb, req.UserID)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	if req.UserID != 0 {
		singleResponse := &ReadRequirementsResponse{
			ID:              dbResponse[0].ID,
			StudentNumber:   dbResponse[0].StudentNumber,
			UploadedFile:    dbResponse[0].UploadedFile,
			RequirementType: dbResponse[0].RequirementType,
		}

		respondJSON(w, 200, singleResponse)
		return
	}

	for _, requirement := range dbResponse {
		response = append(response, &ReadRequirementsResponse{
			ID:              requirement.ID,
			StudentNumber:   requirement.StudentNumber,
			UploadedFile:    requirement.UploadedFile,
			RequirementType: requirement.RequirementType,
		})
	}

	respondJSON(w, 200, response)
}

func (t *TGSCAConfiguration) DeleteRequirement(w http.ResponseWriter, r *http.Request) {

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

	req := &DeleteRequirementsRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, nil)
		return
	}

	fmt.Println(req)

	response := &DeleteRequirementsResponse{}

	err = deleteFileIfExists(req.UploadedFile)
	if err != nil {
		respondJSON(w, 500, nil)
		return
	}

	err = database.DeleteRequirement(t.TGSCAdb, req.ID)
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

func deleteFileIfExists(filePath string) error {
	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// File does not exist, no need to delete
		return nil
	}

	// Delete the file
	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	fmt.Println("File deleted successfully:", filePath)
	return nil
}
