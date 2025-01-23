package model

type Version struct {
	ID      string `json:"id" firestore:"-"`
	Author  string `json:"author" firestore:"author"`
	Content string `json:"content" firestore:"content"`
	Latest  bool   `json:"latest" firestore:"latest"`

	// Deleted is used for logical deletion
	Deleted bool `json:"-" firestore:"deleted"`
}
