package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Ragequit struct {
	Level int
	Time  int
}

type Win struct {
	Level int
	Time  int
}

type Death struct {
	Level     int
	Time      int
	PositionX int
	PositionY int
}

// upload handlers
func upload_ragequit(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var ragequit Ragequit
		err := json.NewDecoder(r.Body).Decode(&ragequit)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Printf("time for ragequit: %d in level %d\n", ragequit.Time, ragequit.Level)
	default:
		fmt.Fprintf(w, "Sorry, only POST methods are supported.\n")
	}
}

func upload_win(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var win Win
		err := json.NewDecoder(r.Body).Decode(&win)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Printf("time for win: %d in level %d\n", win.Time, win.Level)
	default:
		fmt.Fprintf(w, "Sorry, only POST methods are supported.\n")
	}
}

func upload_death(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var death Death
		err := json.NewDecoder(r.Body).Decode(&death)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Printf("time for death: %d in level %d at Position %d,%d\n", death.Time, death.Level, death.PositionX, death.PositionY)
	default:
		fmt.Fprintf(w, "Sorry, only POST methods are supported.\n")
	}
}
