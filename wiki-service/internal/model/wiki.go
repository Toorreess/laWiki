package model

import "time"

type Wiki struct {
	ID          string `json:"id" firestore:"-"`
	Name        string `json:"name" firestore:"name"`
	Author      string `json:"author" firestore:"author"`
	Description string `json:"description" firestore:"description"`

	// Deleted is used for logical deletion
	Deleted bool `json:"-" firestore:"deleted"`
	// Firestore document metadata
	CreationDate     time.Time `json:"creation_date" firestore:"-"`
	ModificationDate time.Time `json:"modification_date" firestore:"-"`
}
