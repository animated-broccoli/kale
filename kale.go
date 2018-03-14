package kale


import (
	_ "encoding/json"
	_"github.com/gorilla/mux"
	"log"
	_"net/http"
)


type Kale struct {

}

func (k Kale) plant(){

	log.Print("Planted A New Kale")


}