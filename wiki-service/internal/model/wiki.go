package model

type Wiki struct {
	ID     string `json:"id" firestore:"-"`
	Name   string `json:"name" firestore:"name"`
	Author string `json:"author" firestore:"author"`
}
