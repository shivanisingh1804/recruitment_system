// package controllers

// import (
// 	"net/http"
// )

// func CreateJob(w http.ResponseWriter, r *http.Request) {
// 	// Placeholder implementation
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Job created successfully"))
// }

//	func GetJob(w http.ResponseWriter, r *http.Request) {
//		// Placeholder implementation
//		w.WriteHeader(http.StatusOK)
//		w.Write([]byte("Job details"))
//	}
package controllers

import (
	"encoding/json"
	"net/http"
	"recruitment-system/config"
	"recruitment-system/models"
)

func CreateJob(w http.ResponseWriter, r *http.Request) {
	var job models.Job
	if err := json.NewDecoder(r.Body).Decode(&job); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Extract user ID from JWT token
	// Validate user type is Admin

	// Save job to the database
	if err := config.DB.Create(&job).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Job created successfully"))
}

func GetJob(w http.ResponseWriter, r *http.Request) {
	// Extract job_id from URL
	// Fetch job details and applicants from the database
}
