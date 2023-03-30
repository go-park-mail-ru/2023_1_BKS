package model

type Post struct {
	ID     int
	User   User
	Header string
	Body   string
	Price  string
}
