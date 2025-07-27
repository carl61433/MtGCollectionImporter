package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Card struct {
	set   string `json:"code"`
	num   string `json:"num"`
	name  string `json:"name"`
	foil  string `json:"foil"`
	count int    `json:"count"`
}

func newCard(set string, num string, foil string) *Card {
	c := Card{}
	c.set = set
	c.num = num
	c.foil = foil
	c.count = 0

	//path := 1
	fmt.Println("MtG Set: ", c.set, "Card number: ", c.num)
	return (&c)
}

func main() {

	fmt.Println("Listening...")
	http.HandleFunc("/submit", cardRequest)
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8585", nil)
}

func cardRequest(rw http.ResponseWriter, r *http.Request) {
	var card Card

	// decode input or return error
	err := json.NewDecoder(r.Body).Decode(&card)
	if err != nil {
		rw.WriteHeader(400)
		fmt.Fprintf(rw, "Decode error! please check your JSON formating.")
		return
	}
	fmt.Println(card.set, card.num, card.foil)
	newCard(card.set, card.num, card.foil)
}
