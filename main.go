package main

import (
  "net/http"
  "html/template"
  "math/rand"
  "os"
)

type Date struct {
  Response string
}

func getLeapYear () string {
  var response string
  // get random number between 1 - 4
  r := rand.Intn(5)
  if r % 4 == 0 {
    response = "Leap Day"
  } else {
    response = "29"
  }
  return response
}

func s (w http.ResponseWriter, r *http.Request) {
  t, _ := template.ParseFiles("s.html")
  d := Date{ Response: getLeapYear() }
  t.Execute(w, d)
}

func main () {
  port := os.Getenv("PORT")
  http.HandleFunc("/", s)
  http.ListenAndServe(":"+port, nil)
}

