package kale

import (
	"net/http"
	"log"
	"time"
)

type HB struct {
    ID      string
    Timestamp       time.Time
}

func heartbeat(w http.ResponseWriter, r *http.Request){
	log.Print("Heartbeat received!")
}	