package models

type User struct {
	Id          int
	First_name  string `json:"first_name" db:"first_name"`
	Second_name string `json:"second_name" db:"second_name"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password"`
	Card_number int    `json:"card_number" db:"card_number"`
	Card_cvv    int    `json:"card_cvv" db:"card_cvv"`
	Image       byte   `json:"image" db:"image"`
}
