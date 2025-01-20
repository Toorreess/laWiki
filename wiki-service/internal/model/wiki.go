package model

import "time"

type Wiki struct {
	ID          string `json:"id" firestore:"-"`
	Name        string `json:"name" firestore:"name" binding:"required" updateAllowed:"true"`
	Creator     string `json:"creator" firestore:"creator" binding:"required"`
	Description string `json:"description" firestore:"description" updateAllowed:"true"`

	ImageURL string `json:"image_url" firestore:"image_url" updateAllowed:"true"`
	// Deleted is used for logical deletion
	Deleted bool `json:"-" firestore:"deleted"`

	// Firestore Document metadata
	CreationDate     time.Time `json:"creation_date" firestore:"-"`
	ModificationDate time.Time `json:"modification_date" firestore:"-"`
}
