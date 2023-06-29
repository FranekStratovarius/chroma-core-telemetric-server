package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Ragequit struct {
	ID        int64
	Level     int64
	Time      float64
	PositionX float64
	PositionY float64
}

type Win struct {
	ID    int64
	Level int64
	Time  float64
}

type Death struct {
	ID        int64
	Level     int64
	Time      float64
	PositionX float64
	PositionY float64
}

// upload handlers
func upload_ragequit(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	switch r.Method {
	case "POST":
		var ragequit Ragequit
		err := json.NewDecoder(r.Body).Decode(&ragequit)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Printf("time for ragequit: %f in level %d at Position %f,%f\n", ragequit.Time, ragequit.Level, ragequit.PositionX, ragequit.PositionY)
		_, err = db.Exec("INSERT INTO ragequits (level, time, positionx, positiony) VALUES ($1, $2, $3, $4)", ragequit.Level, ragequit.Time, ragequit.PositionX, ragequit.PositionY)
		if err != nil {
			log.Print("can't insert into ragequits db: ", err)
		}
	case "GET":
		rows, err := db.Query("SELECT * FROM ragequits WHERE level = 1 ORDER BY time ASC LIMIT 10")
		if err != nil {
			log.Print("can'r read from ragequits db: ", err)
		}
		fmt.Fprintf(w, "Level 1:\n")
		defer rows.Close()
		var ctr int = 0
		for rows.Next() {
			var ragequit Ragequit
			err := rows.Scan(&ragequit.ID, &ragequit.Level, &ragequit.Time, &ragequit.PositionX, &ragequit.PositionY)
			if err != nil {
				log.Print("can't parse ragequit from db: ", err)
			}
			ctr++
			fmt.Fprintf(w, "#%d: %d:%d:%d [%f|%f]\n", ctr, int64(ragequit.Time/3600), int64(ragequit.Time/60)%60, int64(ragequit.Time)%60, ragequit.PositionX, ragequit.PositionY)
		}
		rows, err = db.Query("SELECT * FROM ragequits WHERE level = 2 ORDER BY time ASC LIMIT 10")
		if err != nil {
			log.Print("can'r read from ragequits db: ", err)
		}
		fmt.Fprintf(w, "Level 2:\n")
		defer rows.Close()
		ctr = 0
		for rows.Next() {
			var ragequit Ragequit
			err := rows.Scan(&ragequit.ID, &ragequit.Level, &ragequit.Time, &ragequit.PositionX, &ragequit.PositionY)
			if err != nil {
				log.Print("can't parse ragequit from db: ", err)
			}
			ctr++
			fmt.Fprintf(w, "#%d: %d:%d:%d [%f|%f]\n", ctr, int64(ragequit.Time/3600), int64(ragequit.Time/60)%60, int64(ragequit.Time)%60, ragequit.PositionX, ragequit.PositionY)
		}
	default:
		fmt.Fprintf(w, "Sorry, only POST and GET methods are supported.\n")
	}
}

func upload_win(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	switch r.Method {
	case "POST":
		var win Win
		err := json.NewDecoder(r.Body).Decode(&win)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Printf("time for win: %f in level %d\n", win.Time, win.Level)
		_, err = db.Exec("INSERT INTO wins (level, time) VALUES ($1, $2)", win.Level, win.Time)
		if err != nil {
			log.Print("can't insert into wins db ", err)
		}
	case "GET":
		rows, err := db.Query("SELECT * FROM wins WHERE level = 1 ORDER BY time ASC LIMIT 10")
		if err != nil {
			log.Print("can'r read from wins db: ", err)
		}
		fmt.Fprintf(w, "Level 1:\n")
		defer rows.Close()
		var ctr int = 0
		for rows.Next() {
			var win Win
			err := rows.Scan(&win.ID, &win.Level, &win.Time)
			if err != nil {
				log.Print("can't parse win from db: ", err)
			}
			ctr++
			fmt.Fprintf(w, "#%d: %d:%d:%d\n", ctr, int64(win.Time/3600), int64(win.Time/60)%60, int64(win.Time)%60)
		}
		rows, err = db.Query("SELECT * FROM wins WHERE level = 2 ORDER BY time ASC LIMIT 10")
		if err != nil {
			log.Print("can'r read from wins db: ", err)
		}
		fmt.Fprintf(w, "Level 2:\n")
		defer rows.Close()
		ctr = 0
		for rows.Next() {
			var win Win
			err := rows.Scan(&win.ID, &win.Level, &win.Time)
			if err != nil {
				log.Print("can't parse win from db: ", err)
			}
			ctr++
			fmt.Fprintf(w, "#%d: %d:%d:%d\n", ctr, int64(win.Time/3600), int64(win.Time/60)%60, int64(win.Time)%60)
		}
	default:
		fmt.Fprintf(w, "Sorry, only POST and GET methods are supported.\n")
	}
}

func upload_death(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	switch r.Method {
	case "POST":
		var death Death
		err := json.NewDecoder(r.Body).Decode(&death)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Printf("time for death: %f in level %d at Position %f,%f\n", death.Time, death.Level, death.PositionX, death.PositionY)
		_, err = db.Exec("INSERT INTO deaths (level, time, positionx, positiony) VALUES ($1, $2, $3 ,$4)", death.Level, death.Time, death.PositionX, death.PositionY)
		if err != nil {
			log.Print("can't insert into deaths db ", err)
		}
	case "GET":
		rows, err := db.Query("SELECT * FROM deaths WHERE level = 1 ORDER BY time ASC LIMIT 10")
		if err != nil {
			log.Print("can'r read from deaths db: ", err)
		}
		fmt.Fprintf(w, "Level 1:\n")
		defer rows.Close()
		var ctr int = 0
		for rows.Next() {
			var death Death
			err := rows.Scan(&death.ID, &death.Level, &death.Time, &death.PositionX, &death.PositionY)
			if err != nil {
				log.Print("can't parse death from db: ", err)
			}
			ctr++
			fmt.Fprintf(w, "#%d: %d:%d:%d [%f|%f]\n", ctr, int64(death.Time/3600), int64(death.Time/60)%60, int64(death.Time)%60, death.PositionX, death.PositionY)
		}
		rows, err = db.Query("SELECT * FROM deaths WHERE level = 2 ORDER BY time ASC LIMIT 10")
		if err != nil {
			log.Print("can'r read from deaths db: ", err)
		}
		fmt.Fprintf(w, "Level 2:\n")
		defer rows.Close()
		ctr = 0
		for rows.Next() {
			var death Death
			err := rows.Scan(&death.ID, &death.Level, &death.Time, &death.PositionX, &death.PositionY)
			if err != nil {
				log.Print("can't parse death from db: ", err)
			}
			ctr++
			fmt.Fprintf(w, "#%d: %d:%d:%d [%f|%f]\n", ctr, int64(death.Time/3600), int64(death.Time/60)%60, int64(death.Time)%60, death.PositionX, death.PositionY)
		}
	default:
		fmt.Fprintf(w, "Sorry, only POST and GET methods are supported.\n")
	}
}
