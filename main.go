package main

import (
  	"encoding/json"
  	"net/http"
  	"time"
)

type TimeResponse struct {
	Time string `json:"time"`
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TimeResponse{Time: time.Now().Format(time.RFC3339)})
}

func main() {
	println("Hello world!")
}