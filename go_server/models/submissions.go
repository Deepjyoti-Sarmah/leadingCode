package models

type Submissions struct {
	Submission string `json:"submission"`
	ProblemId  string `json:"problemId"`
	UserId     string `json:"userId"`
	Status     string `json:"status"`
}
