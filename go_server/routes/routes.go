package routes

import (
	"log"
	"net/http"

	"github.com/DeepjyotiSarmah/go_server/controllers"
	"github.com/DeepjyotiSarmah/go_server/middlewares"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func SetupRoutes() {
	r := mux.NewRouter()

	// Create a new CORS handler with desired options
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://127.0.0.1:5173"}, // Replace with your frontend domain
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	// Apply the CORS middleware to your router
	r.Use(corsHandler.Handler)

	// Create routes...
	r.HandleFunc("/", controllers.GetHome).Methods("GET")
	r.HandleFunc("/problems", controllers.GetProblems).Methods("GET")
	r.HandleFunc("/problems/{id}", controllers.GetProblemById).Methods("GET")
	r.HandleFunc("/me", middlewares.Auth(controllers.GetMe)).Methods("GET")
	r.HandleFunc("/submissions/{probelmId}", middlewares.Auth(controllers.GetSubmissionById)).Methods("GET")
	r.HandleFunc("/submission", middlewares.Auth(controllers.PostSubmission)).Methods("POST")
	r.HandleFunc("/signup", controllers.SignUp).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":3000", r))
}
