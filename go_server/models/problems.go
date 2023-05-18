package models

type Problems struct {
	ProblemId   string `json:"problemId"`
	Title       string `json:"tittle"`
	Difficulty  string `json:"difficulty"`
	Acceptance  string `json:"acceptance"`
	Description string `json:"description"`
	ExampleIn   string `json:"exampleIn"`
	ExampleOut  string `json:"exampleOut"`
}
