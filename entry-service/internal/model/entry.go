package model

type Entry struct {
	ID      string `json:"id" firestore:"-"`
	Name    string `json:"name" firestore:"name"`
	Author  string `json:"author" firestore:"author"`
	Content string `json:"content" firestore:"content"`
	WikiID  string `json:"wiki_id" firestore:"wiki_id"`
	Deleted bool   `json:"-" firestore:"deleted"`
}
