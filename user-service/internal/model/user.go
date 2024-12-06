package model

type User struct {
	ID    string `json:"id" firestore:"-"`
	Name  string `json:"name" firestore:"name" binding:"required" updateAllowed:"true"`
	Email string `json:"email" firestore:"email" binding:"required"`

	Rating float32 `json:"rating" firestore:"rating" binding:"required"`

	// Deleted is used for logical deletion
	Deleted bool `json:"-" firestore:"deleted"`
}
