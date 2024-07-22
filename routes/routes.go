package routes

import (
	"net/http"

	"recruitment-system/controllers"
	middleware "recruitment-system/middlewares"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/signup", controllers.SignUp).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.Handle("/uploadResume", middleware.AuthMiddleware(http.HandlerFunc(controllers.UploadResume))).Methods("POST")
	r.Handle("/admin/job", middleware.AuthMiddleware(http.HandlerFunc(controllers.CreateJob))).Methods("POST")
	r.Handle("/admin/job/{job_id}", middleware.AuthMiddleware(http.HandlerFunc(controllers.GetJob))).Methods("GET")
	r.Handle("/admin/applicants", middleware.AuthMiddleware(http.HandlerFunc(controllers.GetAllApplicants))).Methods("GET")
	r.Handle("/admin/applicant/{applicant_id}", middleware.AuthMiddleware(http.HandlerFunc(controllers.GetApplicant))).Methods("GET")
	r.Handle("/jobs", middleware.AuthMiddleware(http.HandlerFunc(controllers.GetJobs))).Methods("GET")
	r.Handle("/jobs/apply", middleware.AuthMiddleware(http.HandlerFunc(controllers.ApplyJob))).Methods("GET")
}
