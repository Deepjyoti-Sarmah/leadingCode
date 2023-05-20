package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/DeepjyotiSarmah/go_server/models"
	// "github.com/dgrijalva/jwt-go"
)

// var problems = []models.Problems{
// 	{
// 		ProblemId:   "1",
// 		Title:       "401. Bitwise AND of Numbers Range",
// 		Difficulty:  "Medium",
// 		Acceptance:  "42%",
// 		Description: "Given two integers left and right that represent the range [left, right], return the bitwise AND of all numbers in this range, inclusive.",
// 		ExampleIn:   "left = 5, right = 7",
// 		ExampleOut:  "4",
// 	},
// 	// Add other problem objects
// }

// var users = []models.User{}

// var userCountId = 1

// const JWT_SECRET = "secret"

// var submissions = []models.Submissions{}

func GetHome(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"msg": "Hello form go Server"}

	responseData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
	// json.NewEncoder(w).Encode(response)
}

// func GetProblems(w http.ResponseWriter, r *http.Request) {
// 	filterProblems := make([]models.Problems, len(problems))

// 	for i, problem := range problems {
// 		filterProblems[i] = models.Problems{
// 			ProblemId:  problem.ProblemId,
// 			Title:      problem.Title,
// 			Difficulty: problem.Difficulty,
// 			Acceptance: problem.Acceptance,
// 		}
// 	}

// 	response := map[string][]models.Problems{"problems": filterProblems}

// 	responseData, err := json.Marshal(response)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(responseData)
// }

// func GetProblemById(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	Id := vars["id"]

// 	var problem *models.Problems

// 	for _, prb := range problems {
// 		if prb.ProblemId == Id {
// 			problem = &prb
// 			break
// 		}
// 	}

// 	if problem == nil {
// 		w.WriteHeader(http.StatusLengthRequired)
// 		json.NewEncoder(w).Encode(map[string]interface{}{})
// 		return
// 	}

// 	response := map[string]interface{}{"problem": problem}
// 	responseData, err := json.Marshal(response)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(responseData)
// }

func GetMe(w http.ResponseWriter, r *http.Request) {

	userId := r.Context().Value("userId")

	var user *models.User

	for _, usr := range users {
		if usr.UserId == userId {
			user = &usr
			break
		}
	}

	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{})
		return
	}

	response := map[string]interface{}{
		"userId": user.UserId,
		"name":   user.Name,
		"email":  user.Email,
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

// func GetSubmissionById(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	ProbelmId := vars["probelmId"]
// 	UserId := vars["userId"]

// 	var filterSubmissions []models.Submissions

// 	for _, submission := range submissions {
// 		if submission.ProblemId == ProbelmId && submission.UserId == UserId {
// 			filterSubmissions = append(filterSubmissions, submission)
// 		}
// 	}

// 	response := map[string][]models.Submissions{
// 		"submission": filterSubmissions,
// 	}

// 	responseData, err := json.Marshal(response)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(responseData)
// }

// func PostSubmission(w http.ResponseWriter, r *http.Request) {

// 	var requestBody struct {
// 		ProbelmId  string `json:"probelmId"`
// 		Submission string `json:"submission"`
// 	}

// 	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	isCorrect := rand.Float64() < 0.5

// 	submission := models.Submissions{
// 		ProblemId: requestBody.ProbelmId,
// 		UserId:    r.Context().Value("userId").(string),
// 		Status:    "WA",
// 	}

// 	if isCorrect {
// 		submission.Status = "AC"
// 	}

// 	submissions = append(submissions, submission)

// 	response := map[string]string{"status": submission.Status}

// 	responseData, err := json.Marshal(response)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(responseData)
// }

// func SignUp(w http.ResponseWriter, r *http.Request) {

// 	var requestSignup struct {
// 		Email    string `json:"email"`
// 		Password string `json:"password"`
// 	}

// 	if err := json.NewDecoder(r.Body).Decode(&requestSignup); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	for _, usr := range users {
// 		if usr.Email == requestSignup.Email {
// 			http.Error(w, "Email already exits", http.StatusForbidden)
// 			return
// 		}
// 	}

// 	user := models.User{
// 		Email:    requestSignup.Email,
// 		Password: requestSignup.Password,
// 		UserId:   string(userCountId),
// 	}

// 	users = append(users, user)
// 	userCountId++

// 	response := map[string]string{"msg": "Success"}

// 	responseData, err := json.Marshal(response)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(responseData)
// }

// func Login(w http.ResponseWriter, r *http.Request) {

// 	var requestLogin struct {
// 		Email    string `json:"email"`
// 		Password string `json:"password"`
// 	}

// 	if err := json.NewDecoder(r.Body).Decode(&requestLogin); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	var user *models.User

// 	for _, usr := range users {
// 		if usr.Email == requestLogin.Email {
// 			user = &usr
// 			break
// 		}
// 	}

// 	if user == nil {
// 		http.Error(w, "User not found", http.StatusForbidden)
// 		return
// 	}

// 	if user.Password != requestLogin.Password {
// 		http.Error(w, "Incorrect Password", http.StatusForbidden)
// 		return
// 	}

// 	//Generate JWT Token
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"id": user.UserId,
// 	})

// 	signedToken, err := token.SignedString([]byte(JWT_SECRET))
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	response := map[string]string{"token": signedToken}

// 	responseData, err := json.Marshal(response)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(responseData)

// }
