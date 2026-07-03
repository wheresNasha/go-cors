package models

// User represents a user resource (for future CRUD with DB).
type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
