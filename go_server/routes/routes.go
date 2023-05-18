package routes

import (
	"log"
	"net/http"

	"github.com/DeepjyotiSarmah/go_server/controllers"
	"github.com/DeepjyotiSarmah/go_server/middlewares"

	"github.com/gorilla/mux"
)

func SetupRoutes() {
	r := mux.NewRouter()

	//create routes
	r.HandleFunc("/", controllers.GetHome).Methods("GET")
	r.HandleFunc("/problems", controllers.GetProblems).Methods("GET")
	r.HandleFunc("/problems/{id}", controllers.GetProblemById).Methods("GET")
	r.HandleFunc("/me", middlewares.Auth(controllers.Me)).Methods("GET")
	r.HandleFunc("/submissions/{id}", middlewares.Auth(controllers.GetSubmissionById)).Methods("GET")
	r.HandleFunc("/submission", middlewares.Auth(controllers.PostSubmission)).Methods("POST")
	r.HandleFunc("/sigin", controllers.SignIn).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	//start server
	log.Fatal(http.ListenAndServe(":3000", r))

}
