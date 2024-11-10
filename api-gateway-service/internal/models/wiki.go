package models

type Wiki struct {
	Name        string `json:"name" firestore:"name"`
	Description string `json:"description" firestore:"description"`
}
