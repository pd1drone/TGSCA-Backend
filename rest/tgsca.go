package rest

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (c *TGSCAConfiguration) check(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func Routes() {
	// Check if the log file exists
	if _, err := os.Stat("/root/TGSCA-Backend/http_logs.log"); os.IsNotExist(err) {
		// Create a log file
		logFile, err := os.Create("/root/TGSCA-Backend/http_logs.log")
		if err != nil {
			log.Fatal(err)
		}
		defer logFile.Close()
	}

	// Open the log file
	logFile, err := os.OpenFile("/root/TGSCA-Backend/http_logs.log", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	// // Create a new logger using the log file
	logger := middleware.RequestLogger(&middleware.DefaultLogFormatter{
		Logger:  log.New(logFile, "", log.LstdFlags), // Use the log file as the output
		NoColor: true,
	})

	log.Print("Starting TGSCA Backend Service.....")
	r := chi.NewRouter()

	newTsgsa, err := New()
	if err != nil {
		log.Fatal(err)
	}
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(logger)
	r.Use(func(next http.Handler) http.Handler {
		return logPayloads(logFile, next)
	})

	// Check route if service is running
	r.Get("/check", newTsgsa.check)

	// Login routes
	r.Post("/loginAdmin", newTsgsa.LoginAdmin)
	r.Post("/loginStudent", newTsgsa.LoginStudent)

	// Change Password route
	r.Post("/changePassword", newTsgsa.ChangePassword)

	// CRUD Student routes
	r.Post("/createStudent", newTsgsa.CreateStudent)
	r.Post("/readStudent", newTsgsa.ReadStudent)
	r.Post("/updateStudent", newTsgsa.UpdateStudent)
	r.Delete("/deleteStudent", newTsgsa.DeleteStudent)

	//CRUD requirements
	// File Upload route CREATE requirement
	r.Post("/upload", newTsgsa.UploadRequirements)
	// File Serving route View requirement
	r.Get("/files/{userID}/{filename}", newTsgsa.ServeFile)
	r.Post("/readRequirements", newTsgsa.ReadRequirements)
	r.Delete("/deleteRequirements", newTsgsa.DeleteRequirement)

	// CRUD Teachers routes
	r.Post("/createTeacher", newTsgsa.CreateTeachers)
	r.Post("/readTeacher", newTsgsa.ReadTeachers)
	r.Post("/updateTeacher", newTsgsa.UpdateTeachers)
	r.Delete("/deleteTeacher", newTsgsa.DeleteTeachers)

	//Create READ Appointments
	r.Post("/createAppointment", newTsgsa.CreateAppointment)
	r.Post("/readAppointment", newTsgsa.ReadAppointment)

	//CRUD Subjects
	r.Post("/createSubject", newTsgsa.CreateSubject)
	r.Post("/readSubject", newTsgsa.ReadSubject)
	r.Post("/readSubjectStudent", newTsgsa.ReadSubjectForStudent)
	r.Post("/readSubjectSchedule", newTsgsa.ReadSubjectSchedule)
	r.Post("/readSubjectStudentDropdown", newTsgsa.ReadSubjectForStudentDropDown)
	r.Post("/readSubjectGradeLevel", newTsgsa.ReadSubjectGradeLevel)
	r.Post("/updateSubject", newTsgsa.UpdateSubject)
	r.Delete("/deleteSubject", newTsgsa.DeleteSubject)

	// CRUD ENROLLED ADMIN
	r.Post("/createEnrolled", newTsgsa.CreateEnrolled)
	r.Post("/readEnrolled", newTsgsa.ReadEnrolled)
	r.Post("/updateEnrolled", newTsgsa.UpdateEnrolled)
	r.Delete("/deleteEnrolled", newTsgsa.DeleteEnrolled)

	// CRUD ENROLLMENT STUDENT
	r.Post("/createEnrolledStudent", newTsgsa.CreateEnrolledForStudent)
	r.Post("/readEnrolledStudent", newTsgsa.ReadEnrolledForStudent)
	r.Post("/updateEnrolledStudent", newTsgsa.UpdateEnrolledForStudent)
	r.Delete("/deleteEnrolledStudent", newTsgsa.DeleteEnrolledForStudent)

	// READ ENROLLED SUBJECTS FOR STUDENT
	r.Post("/readEnrolledSubjects", newTsgsa.ReadEnrolledSubjects)

	// CRUD Enrollment Status
	r.Post("/getStudentEnrollmentStatus", newTsgsa.GetStudentEnrollmentStatus)
	r.Post("/approve", newTsgsa.ApproveStudentEnrollment)

	log.Fatal(http.ListenAndServe("0.0.0.0:8082", r))
}

func logPayloads(logFile *os.File, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "*")
		w.Header().Add("Access-Control-Allow-Headers", "*")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Record start time of request processing
		startTime := time.Now()

		// Call the next middleware/handler in the chain and capture response
		responseRecorder := httptest.NewRecorder()
		next.ServeHTTP(responseRecorder, r)

		// Record end time of request processing
		endTime := time.Now()

		// Read request payload
		requestBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Error reading request body:", err)
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}

		// Restore request body after reading
		r.Body = io.NopCloser(bytes.NewBuffer(requestBody))

		// Log the request and response payloads
		log.SetOutput(logFile)
		log.Printf("Request: %s %s\n", r.Method, r.URL.Path)

		if len(requestBody) > 0 {
			log.Printf("Request Payload: %s\n", string(requestBody))
		} else {
			log.Println("Request Payload: [Empty]")
		}

		// Read response payload
		responseBody := responseRecorder.Body.Bytes()

		if len(responseBody) > 0 {
			log.Printf("Request Response: %d %s\nPayload: %s\n", responseRecorder.Code, http.StatusText(responseRecorder.Code), string(responseBody))
		} else {
			log.Printf("Request Response: %d %s\nPayload: [Empty]\n", responseRecorder.Code, http.StatusText(responseRecorder.Code))
		}
		log.Printf("Duration: %v\n", endTime.Sub(startTime))

		// Write response back to original response writer
		for k, v := range responseRecorder.Header() {
			w.Header()[k] = v
		}
		w.WriteHeader(responseRecorder.Code)
		w.Write(responseRecorder.Body.Bytes())

		log.Printf("\n")
	})
}
