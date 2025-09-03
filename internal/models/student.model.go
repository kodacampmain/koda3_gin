package models

type Student struct {
	Id   int    `db:"id" json:"id"`
	Name string `db:"name" json:"nama_siswa"`
	// Role sql.NullString `db:"role"`
	Role *string `db:"role" json:"peran"`
}
