package main

import (
  	"encoding/json"
  	"net/http"
  	"time"
)

type TimeResponse struct {
	Time string `json:"time"`
}

func main() {
	println("Hello world!")
}