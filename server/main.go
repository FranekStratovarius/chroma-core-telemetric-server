package main

import (
	"errors"
	"log"
	"net/http"
	"os"
)

// func check(e error) {
// 	if e != nil {
// 		panic(e)
// 	}
// }

func main() {
	var db = connect()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print("> /")
		log.Print(r)
	})

	http.HandleFunc("/ragequit", func(w http.ResponseWriter, r *http.Request) {
		log.Print("> /ragequit")
		log.Print(r)
		log.Print(r.Body)
		upload_ragequit(w, r, db)
	})
	http.HandleFunc("/win", func(w http.ResponseWriter, r *http.Request) {
		log.Print("> /win")
		log.Print(r)
		log.Print(r.Body)
		upload_win(w, r, db)
	})
	http.HandleFunc("/death", func(w http.ResponseWriter, r *http.Request) {
		log.Print("> /death")
		log.Print(r)
		log.Print(r.Body)
		upload_death(w, r, db)
	})

	log.Print("Server running...")
	err := http.ListenAndServe(":1234", nil)

	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("server close")
	} else if err != nil {
		log.Printf("error starting server: %s", err)
		os.Exit(1)
	}

}
