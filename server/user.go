package server

type User struct {
	Email    string
	Password string
}

var user = make(map[string]User)
