package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/DeepjyotiSarmah/go_server/models"
	"github.com/gorilla/mux"
)

var problems = []models.Problems{
	{
		ProblemId:   "1",
		Title:       "401. Bitwise AND of Numbers Range",
		Difficulty:  "Medium",
		Acceptance:  "42%",
		Description: "Given two integers left and right that represent the range [left, right], return the bitwise AND of all numbers in this range, inclusive.",
		ExampleIn:   "left = 5, right = 7",
		ExampleOut:  "4",
	},
	// Add other problem objects
}

func GetProblems(w http.ResponseWriter, r *http.Request) {
	filterProblems := make([]models.Problems, len(problems))

	for i, problem := range problems {
		filterProblems[i] = models.Problems{
			ProblemId:  problem.ProblemId,
			Title:      problem.Title,
			Difficulty: problem.Difficulty,
			Acceptance: problem.Acceptance,
		}
	}

	response := map[string][]models.Problems{"problems": filterProblems}

	responseData, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}

func GetProblemById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id := vars["id"]

	var problem *models.Problems

	for _, prb := range problems {
		if prb.ProblemId == Id {
			problem = &prb
			break
		}
	}

	if problem == nil {
		w.WriteHeader(http.StatusLengthRequired)
		json.NewEncoder(w).Encode(map[string]interface{}{})
		return
	}

	response := map[string]interface{}{"problem": problem}
	responseData, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}
