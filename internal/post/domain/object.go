package domain

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	IdPost uuid.UUID
	Title  Title
}

type PostDeliveryInterface struct {
	UserID     uuid.UUID
	Close      bool
	Title      string
	Desciption string
	Price      string
	Tags       []string
	Images     []string
	Time       time.Time
}

type Images []string

func (l Images) String() []string {
	var newString []string
	for i := 0; i < len(l); i++ {
		newString = append(newString, l[i])
	}
	return newString
}

type Desciption string

func (l Desciption) String() string {
	return string(l)
}

type Title string

func (l Title) String() string {
	return string(l)
}

type Tags []string

func (l Tags) String() []string {
	var newString []string
	for i := 0; i < len(l); i++ {
		newString = append(newString, l[i])
	}
	return newString
}

type Price string

func (l Price) String() string {
	return string(l)
}

type Close bool

func (l Close) Bool() bool {
	return bool(l)
}

type TimeStamp time.Time

func (l TimeStamp) Time() time.Time {
	return time.Time(l)
}
