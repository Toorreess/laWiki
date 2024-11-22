package models

type Comment struct {
	ID      string `json:"id" firestore:"-"`
	Author  string `json:"author" firestore:"author"`
	EntryID string `json:"entry_id" firestore:"entry_id"`
	Content string `json:"content" firestore:"content"`
}
