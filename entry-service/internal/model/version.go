package model

type Version struct {
	ID         string `json:"id" firestore:"-"`
	Author     string `json:"author" firestore:"author"`
	ContentURL string `json:"content_url" firestore:"content_url"`
	Latest     bool   `json:"latest" firestore:"latest"`

	// Deleted is used for logical deletion
	Deleted bool `json:"-" firestore:"deleted"`
}
