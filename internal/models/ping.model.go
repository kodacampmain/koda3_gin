package models

type Body struct {
	Id      int    `json:"id" binding:"required,min=0"`
	Message string `json:"msg"`
	Gender  string `json:"gender"`
}
