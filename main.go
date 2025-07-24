package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Listening...")
	http.HandleFunc("/addToList", newCard)
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8585", nil)
}

func newCard(rw http.ResponseWriter, r *http.Request) {
	setCode := r.FormValue("set-code")
	//cardNo := r.FormValue("card-number")

	fmt.Println("MtG Set: ", setCode)
	//Do more stuff here
}
