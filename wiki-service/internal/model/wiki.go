package model

type Wiki struct {
	ID          string `json:"id" firestore:"-"`
	Name        string `json:"name" firestore:"name"`
	Author      string `json:"author" firestore:"author"`
	Creation    int    `json:"creation" firestore:"creation"`
	Description string `json:"description" firestore:"description"`

	Deleted bool `json:"-" firestore:"deleted"`
}
