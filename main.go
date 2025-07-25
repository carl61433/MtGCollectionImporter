package main

import (
	"fmt"
	"net/http"
)

type card struct {
	set  string
	num  string
	name string
	foil bool
}

func newCard(set string, num string) *card {
	c := card{}
	c.set = set
	c.num = num

	//path := 1
	//fmt.Println("MtG Set: ", set, "Card number: ", num)
	return (&c)
}

func main() {
	fmt.Println("Listening...")
	http.HandleFunc("/submit", cardRequest)
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8585", nil)
}

func cardRequest(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
	setCode := r.FormValue("set-code")
	cardNo := r.FormValue("card-number")
	newCard(setCode, cardNo)
}
