package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/DeepjyotiSarmah/go_server/models"
	"github.com/gorilla/mux"
)

var submissions = []models.Submissions{}

func GetSubmissionById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	ProbelmId := vars["probelmId"]
	UserId := vars["userId"]

	var filterSubmissions []models.Submissions

	for _, submission := range submissions {
		if submission.ProblemId == ProbelmId && submission.UserId == UserId {
			filterSubmissions = append(filterSubmissions, submission)
		}
	}

	response := map[string][]models.Submissions{
		"submission": filterSubmissions,
	}

	responseData, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}

func PostSubmission(w http.ResponseWriter, r *http.Request) {

	var requestBody struct {
		ProblemId  string `json:"problemId"`
		Submission string `json:"submission"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	isCorrect := rand.Float64() < 0.5

	submission := models.Submissions{
		ProblemId: requestBody.ProblemId,
		UserId:    r.Context().Value("userId").(string),
		Status:    "WA",
	}

	if isCorrect {
		submission.Status = "AC"
	}

	submissions = append(submissions, submission)

	response := map[string]string{"status": submission.Status}

	responseData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}
