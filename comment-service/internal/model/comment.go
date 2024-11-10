package model

type Comment struct {
	ID      string `json:"id" firestore:"-"`
	Author  string `json:"author" firestore:"author"`
	EntryID string `json:"entry_id" firestore:"entry_id"`
	Comment string `json:"comment" firestore:"comment"`

	// Deleted is used for logical deletion
	Deleted bool `json:"-" firestore:"deleted"`
}
