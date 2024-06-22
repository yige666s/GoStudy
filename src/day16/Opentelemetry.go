package main

import (
	"log"
	"net/http"
)

func Dice() {
	http.HandleFunc("/roll", roll)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
