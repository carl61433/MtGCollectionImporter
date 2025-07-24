package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Listening...")
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("addToList", newCard)
	http.ListenAndServe(":8585", nil)
}

func newCard(rw http.ResponseWriter, r *http.Request) {
	setCode := r.PostFormValue("set-code")
	//cardNo := r.FormValue("card-number")

	fmt.Println("MtG Set: ", setCode)
}
