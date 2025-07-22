package main

import (
	"net/http"
	"github.com/russross/blackfriday/v2"
)

func main() {
	http.ListenAndServe(":8585", http.FileServer(http.Dir(".")))
}
