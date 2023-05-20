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
	{
		ProblemId:   "2",
		Title:       "205. Add two numbers",
		Difficulty:  "Medium",
		Acceptance:  "41%",
		Description: "Given two numbers, add them and return them in integer range. use MOD=1e9+7",
		ExampleIn:   "a = 100 , b = 200",
		ExampleOut:  "300",
	},
	{
		ProblemId:   "3",
		Title:       "202. Happy Number",
		Difficulty:  "Easy",
		Acceptance:  "54.9%",
		Description: "Write an algorithm to determine if a number n is happy.",
		ExampleIn:   "n = 19",
		ExampleOut:  "true",
	},
	{
		ProblemId:   "4",
		Title:       "203. Remove Linked List Elements",
		Difficulty:  "Hard",
		Acceptance:  "42%",
		Description: "Given number k , removed kth element",
		ExampleIn:   "list: 1->2->3 , k=2",
		ExampleOut:  "1->3",
	},
	{
		ProblemId:   "5",
		Title:       "201. Bitwise AND of Numbers Range",
		Difficulty:  "Medium",
		Acceptance:  "42%",
		Description: "Given two integers left and right that represent the range [left, right], return the bitwise AND of all numbers in this range, inclusive.",
		ExampleIn:   "left = 5, right = 7",
		ExampleOut:  "4",
	},
	{
		ProblemId:   "6",
		Title:       "205. Add two numbers",
		Difficulty:  "Medium",
		Acceptance:  "41%",
		Description: "Given two numbers, add them and return them in integer range. use MOD=1e9+7",
		ExampleIn:   "a = 100 , b = 200",
		ExampleOut:  "300",
	},
	{
		ProblemId:   "7",
		Title:       "202. Happy Number",
		Difficulty:  "Easy",
		Acceptance:  "54.9%",
		Description: "Write an algorithm to determine if a number n is happy.",
		ExampleIn:   "n = 19",
		ExampleOut:  "true",
	},
	{
		ProblemId:   "8",
		Title:       "203. Remove Linked List Elements",
		Difficulty:  "Hard",
		Acceptance:  "42%",
		Description: "Given number k , removed kth element",
		ExampleIn:   "list: 1->2->3 , k=2",
		ExampleOut:  "1->3",
	},
}

func GetProblems(w http.ResponseWriter, r *http.Request) {
	filterProblems := make([]models.Problems, len(problems))

	for i, problem := range problems {
		filterProblems[i] = models.Problems{
			ProblemId:   problem.ProblemId,
			Title:       problem.Title,
			Difficulty:  problem.Difficulty,
			Acceptance:  problem.Acceptance,
			Description: problem.Description,
			ExampleIn:   problem.ExampleIn,
			ExampleOut:  problem.ExampleOut,
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
