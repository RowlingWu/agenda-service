package entity

// User is used as a format in sqlite's table.
type User struct {
	// ID should BE a primary key, auto-increment, not null
	Id int
	// NAME is not null
	Name string
	// PASSWD is not null
	Passwd string
	Email  string
	Tel    string
}
