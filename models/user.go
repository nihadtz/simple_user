package models

type User struct {
	ID          int64  `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	YearOfBirth int    `db:"year_of_birth" json:"year_of_birth"`
	Updated     int64  `db:"updated" json:"updated"`
}
