package model

type User struct {
	ID    string `json:"id" firestore:"-"`
	Name  string `json:"name" firestore:"name" binding:"required" updateAllowed:"true"`
	Email string `json:"email" firestore:"email" binding:"required"`
	Role  string `json:"role" firestore:"role" binding:"oneof=admin "`

	Reputation  float64 `json:"reputation" firestore:"reputation" binding:"required"`
	RatingCount int     `json:"rating_count" firestore:"rating_count" binding:"required"`

	// Deleted is used for logical deletion
	Deleted bool `json:"-" firestore:"deleted"`
}
