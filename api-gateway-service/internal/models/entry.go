package models

type Entry struct {
	ID      string `json:"id" firestore:"-"`
	Name    string `json:"name" firestore:"name"`
	Author  string `json:"author" firestore:"author"`
	WikiID  string `json:"wiki_id" firestore:"wiki_id"`
	Content string `json:"content" firestore:"-"`
}
