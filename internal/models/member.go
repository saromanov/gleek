package models

// Member defines part of the worker
// on company. It can be Software engineer, Devops, etc
type Member struct {
	ID         int64
	Rating     float64
	FirstName  string
	LastName   string
	Competence string
}
