// package controllers

// import (
// 	"net/http"
// )

//	func UploadResume(w http.ResponseWriter, r *http.Request) {
//		// Placeholder implementation
//		w.WriteHeader(http.StatusOK)
//		w.Write([]byte("Resume uploaded successfully"))
//	}
package controllers

import (
	"io"
	"net/http"
	"os"
	"recruitment-system/config"
	"recruitment-system/models"
	"recruitment-system/utils"
)

func UploadResume(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from JWT token
	// Validate user type is Applicant

	// Parse the file from request
	file, _, err := r.FormFile("resume")
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Save the file to a temporary location
	tempFile, err := os.CreateTemp("", "*.pdf")
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, file)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}

	// Process the resume using third-party API
	resumeData, err := utils.ParseResume(tempFile.Name())
	if err != nil {
		http.Error(w, "Failed to process resume", http.StatusInternalServerError)
		return
	}

	// Save the resume data to the database
	var user models.User
	// Get user ID from JWT token and find the user

	profile := models.Profile{
		UserID:     user.ID,
		ResumeFile: tempFile.Name(),
		Skills:     resumeData.Skills,
		Education:  resumeData.Education,
		Experience: resumeData.Experience,
	}

	if err := config.DB.Create(&profile).Error; err != nil {
		http.Error(w, "Failed to save resume data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Resume uploaded successfully"))
}
