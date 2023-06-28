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
	http.HandleFunc("/ragequit", upload_ragequit)
	http.HandleFunc("/win", upload_win)
	http.HandleFunc("/death", upload_death)

	log.Print("Server running...")
	err := http.ListenAndServe(":1234", nil)

	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("server close")
	} else if err != nil {
		log.Printf("error starting server: %s", err)
		os.Exit(1)
	}

}
