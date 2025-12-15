package gmg

import (
	"time"
)

type User struct {
	ID        string
	Name      string
	Email     string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Address struct {
	ID          string
	HouseNumber string
	Street      string
	City        string
	State       string
	Country     string
	Zip         string
}
