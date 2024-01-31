package model

type User struct {
	UserName string
	UserId   string
	Password string
}

type UserDetails struct {
	UserDetails []User
}
