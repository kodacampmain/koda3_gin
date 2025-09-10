package models

import "mime/multipart"

type StudentData struct {
	Id       int     `db:"id" json:"-"`
	Name     string  `db:"name" form:"name" json:"name"`
	Role     string  `db:"role" json:"-"`
	Password string  `db:"password" form:"pwd" json:"pwd,omitempty"`
	Image    *string `db:"images" json:"image"`
}

type StudentAuth struct {
	Name     string `json:"nama_siswa" binding:"required"`
	Password string `json:"password" binding:"required,min=4"`
}

type StudentBody struct {
	StudentData
	Image *multipart.FileHeader `form:"image"`
}

type AuthData struct {
	Token string `json:"token" example:"jwt token"`
}

type AuthResponse struct {
	SuccessResponse
	Data AuthData `json:"data"`
}

type StudentResponse struct {
	SuccessResponse
	Data StudentData `json:"data"`
}
