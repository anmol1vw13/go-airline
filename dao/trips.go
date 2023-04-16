package dao

import "database/sql"

func GetLatestTrip(db *sql.DB) Trip {
	var trip Trip
	db.QueryRow("select id, name from Trips order by id desc limit 1").Scan(&trip.ID, &trip.Name)
	return trip
}