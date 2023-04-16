package main

import (
	"airline/dao"
	"airline/db"
	"database/sql"
	"fmt"
	"sync"
	"time"
)

func main() {
	db := db.Initiate()
	users := dao.GetUsers(db)
	trip := dao.GetLatestTrip(db)
	dao.PrintSeats(db, trip)
	var wg sync.WaitGroup
	start := time.Now()
	for _, user := range(users) {
		wg.Add(1)
		go func (db *sql.DB, tripId int, user dao.User, wg *sync.WaitGroup) {
			dao.AllocateSeat(db, tripId, user)
			wg.Done()
		}(db, trip.ID, user, &wg)
	}
	wg.Wait()
	end := time.Now()
	fmt.Printf("Time taken %d ms \n", (end.UnixMilli() - start.UnixMilli()))
	dao.PrintSeats(db, trip)
}