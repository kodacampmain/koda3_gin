package models

type Student struct {
	Id       int    `db:"id" json:"id"`
	Name     string `db:"name" json:"nama_siswa"`
	Role     string `db:"role" json:"peran"`
	Password string `db:"password" json:"password,omitempty"`
}

type StudentAuth struct {
	Name     string `json:"nama_siswa" binding:"required"`
	Password string `json:"password" binding:"required,min=4"`
}
