package model

type User struct {
    Id int `db:"id" json:"id"`
    Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
    Password string `db:"password" json:"password"`
    Salt string `db:"salt" json:"salt"`
    JWT string `db:"token" json:"token"`
    Role string `db:"role" json:"role"`
}
