package model

import "time"

type Entry struct {
	ID            string   `json:"id" firestore:"-"`
	Name          string   `json:"name" firestore:"name"`
	Author        string   `json:"author" firestore:"author"`
	WikiID        string   `json:"wiki_id" firestore:"wiki_id"`
	LatestVersion string   `json:"latest_version" firestore:"latest_version"`
	VersionList   []string `json:"version_list" firestore:"-"`
	Content       string   `json:"content_url" firestore:"-"`

	// Deleted is used for logical deletion
	Deleted bool `json:"-" firestore:"deleted"`

	// Firestore document metadata
	CreationDate     time.Time `json:"creation_date" firestore:"-"`
	ModificationDate time.Time `json:"modification_date" firestore:"-"`
}

type Version struct {
	ID         string `json:"id" firestore:"-"`
	Author     string `json:"author" firestore:"author"`
	ContentURL string `json:"content_url" firestore:"content_url"`
	Latest     bool   `json:"latest" firestore:"latest"`

	// Deleted is used for logical deletion
	Deleted bool `json:"-" firestore:"deleted"`
}
