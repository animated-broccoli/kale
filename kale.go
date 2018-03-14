package kale


import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	_"fmt"
	"github.com/google/uuid"
	
)


type Kale struct {
	//nodes []int
	ID uuid.UUID
}


func Handle_Heartbeat(w http.ResponseWriter, r *http.Request) {
	go heartbeat(w,r)
	
}

func Handle_Poll(w http.ResponseWriter, r *http.Request) {
	
	var td = HB{ID: ID.String()}
	json.NewEncoder(w).Encode(td)
	//go poll(&w,r)
	
}

var ID uuid.UUID



func (k Kale) Plant(){
	k.ID = uuid.New()
	ID = k.ID
	log.Print("Planted A New Kale: ",k.ID)
	router := mux.NewRouter()
	router.HandleFunc("/heartbeat", Handle_Heartbeat).Methods("GET")
	router.HandleFunc("/poll", Handle_Poll).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
	//fmt.Println("Planted a New Kale")

}