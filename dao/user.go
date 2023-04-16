package dao

import (
	"database/sql"
	"log"
)


func GetUsers (db *sql.DB) []User {
	rows, err := db.Query("select id, name from users")
	defer rows.Close()

	if err != nil {
		log.Fatalln("Error", err)
		return nil
	}

	var users []User

	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.Name)
		users = append(users, user)
	}
	
	return users
}