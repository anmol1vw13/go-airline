package dao

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)


func GetSeats (db *sql.DB, tripId int) []Seat {
	
	rows, rowErr := db.Query(fmt.Sprintf("select id, user_id, name from seats where trip_id = %d", tripId) )
	if rowErr != nil {
		log.Println("Error getting seats", rowErr)
		return nil
	}
	
	var seats []Seat

	for rows.Next() {
		seat := new(Seat)
	
		rows.Scan(&seat.ID, &seat.UserID, &seat.Name)
		seats = append(seats, *seat)
	}

	// log.Println("Seats here are")
	// log.Println(seats)
	
	return seats
}

func AllocateSeat(db *sql.DB, tripId int, user User) int{

	var seatId int
	tx, txErr := db.Begin()
	if txErr != nil {
		fmt.Println("Error beginning the transaction", txErr)
	}
	stmt, err := tx.Prepare("select id from seats where trip_id=$1 AND user_id is null order by id limit 1 FOR UPDATE SKIP LOCKED")

	if err!=nil {
		fmt.Println("Error with prepared statement", err)
		panic(err.Error())
	}
	defer stmt.Close()
	
	queryError := stmt.QueryRow(tripId).Scan(&seatId)
	if queryError != nil {
		fmt.Println("Error while allocating seat to user ", user.Name, err)
		return 0
	}

	fmt.Println("Allocating Seat Id: ", seatId, "for User Id: ", user.ID)
	updateStmt, updateErr := tx.Prepare("update seats set user_id = $1 where id = $2")
	defer updateStmt.Close()
	if updateErr != nil {
		fmt.Println("Error performing the update", updateErr)
		panic(updateErr)
	}
	updateStmt.Exec(user.ID, seatId)
	
	tx.Commit()
	return seatId
}


func PrintSeats(db *sql.DB, trip Trip) {

	seats := GetSeats(db, trip.ID)

	chars := [6]string{"A", "B", "C", "D", "E", "F"}
	
	for _, char := range chars {
		for _, seat := range seats {
			if strings.Split(*seat.Name, "") [1] == char {
				if seat.UserID.Valid {
					fmt.Printf("%d  ",seat.UserID.Int64)
				} else {
					fmt.Print("x  ")
				}
			}
		}
		fmt.Println()
		if char == "C" {
			fmt.Println()
		}
	}
}