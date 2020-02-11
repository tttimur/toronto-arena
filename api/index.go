package handler

import (
	"fmt"
	"math/rand"
	"net/http"
)

type date struct {
	Response string
}

func getLeapYear() string {
	var response string
	// get random number between 0-3
	r := rand.Intn(4)

	if r%4 == 0 {
		response = "Leap Day"
	} else {
		response = "29"
	}
	return response
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, getLeapYear())
}
