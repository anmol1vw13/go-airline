package dao

import "database/sql"

type Flight struct {
	ID int
	Name string
}

type Trip struct {
	ID int
	FlightID int
	Name string
}

type User struct {
	ID int
	Name string
}

type Seat struct {
	ID int
	TripID int
	UserID sql.NullInt64
	Name *string
}