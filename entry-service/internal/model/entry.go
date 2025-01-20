package model

import "time"

type Entry struct {
	ID            string `json:"id" firestore:"-"`
	Name          string `json:"name" firestore:"name" binding:"required"`
	Author        string `json:"author" firestore:"author" binding:"required"`
	WikiID        string `json:"wiki_id" firestore:"wiki_id" binding:"required"`
	LatestVersion string `json:"latest_version" firestore:"latest_version" binding:"required"`
	Content       string `json:"content" firestore:"-"`
	
	ImageURL string `json:"image_url" firestore:"image_url"`

	Location struct{
		Latitude float64 `json:"latitude" firestore:"latitude"`
		Longitude float64 `json:"longitude" firestore:"longitude"` 
	} `json:"location" firestore:"location"`

	// Deleted is used for logical deletion
	Deleted bool `json:"-" firestore:"deleted"`

	// Firestore document metadata
	CreationDate     time.Time `json:"creation_date" firestore:"-"`
	ModificationDate time.Time `json:"modification_date" firestore:"-"`
}
