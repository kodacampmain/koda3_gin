package models

import (
	"mime/multipart"
	"time"
)

type Movies struct {
	Id           int           `db:"id" json:"id"`
	Title        string        `db:"title" json:"title" form:"title"`
	Poster       string        `db:"poster" json:"poster"`
	Backdrop     string        `db:"backdrop" json:"backdrop"`
	ReleaseDate  time.Time     `db:"release_date" json:"release_date" form:"release_date"`
	Duration     time.Duration `db:"duration" json:"duration" form:"duration"`
	Synopsis     string        `db:"synopsis" json:"synopsis" form:"synopsis"`
	Rating       float32       `db:"rating" json:"rating" form:"rating"`
	DirectorName string        `db:"director_name" json:"director_name"`
	Genres       []Genre       `db:"genres" json:"genres"`
	Actors       []Actor       `db:"actors" json:"actors"`
}

type Genre struct {
	Id   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}
type Actor struct {
	Id   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type MovieBody struct {
	Movies
	Poster   *multipart.FileHeader `form:"poster"`
	Backdrop *multipart.FileHeader `form:"backdrop"`
	// Title       string                `form:"title"`
	// ReleaseDate time.Time             `form:"release_date"`
	// Duration    time.Duration         `form:"duration"`
	// Synopsis    string                `form:"synopsis"`
	// Rating      float32               `form:"rating"`
}
