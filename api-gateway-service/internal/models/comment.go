package models

type Comment struct {
	Author  string `json:"author" binding:"required"`
	Content string `json:"content" binding:"required"`
}
