package models

type User struct {
	ID    string `json:"id" firestore:"-"`
	Name  string `json:"name" firestore:"name" binding:"required" updateAllowed:"true"`
	Email string `json:"email" firestore:"email" binding:"required"`
	Role  string `json:"role" firestore:"role" binding:"oneof=admin "`

	Reputation  float64 `json:"reputation" firestore:"reputation" binding:"required"`
	RatingCount int     `json:"rating_count" firestore:"rating_count" binding:"required"`

	Notifications []Notification `json:"notifications" firestore:"notifications"`

	// Deleted is used for logical deletion
	Deleted bool `json:"-" firestore:"deleted"`
}

type Notification struct {
	ID      string `json:"id" firestore:"-"`
	Message string `json:"message" firestore:"message"`
	Read    bool   `json:"read" firestore:"read"`
}
