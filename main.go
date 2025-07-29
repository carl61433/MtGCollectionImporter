package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Card struct {
	Set   string `json:"set"`
	Num   string `json:"num"`
	Foil  bool   `json:"foil"`
	Name  string `json:"name"`
	Count int
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
		fmt.Fprintf(rw, "Decode error! please check your JSON formating.\n")
		return
	}
	//fmt.Printf("Output: " + card.Set + "," + card.Num + "," + card.Foil)
	//fmt.Println(card) //Example: {FIN 1 false  0}
	url := "https://api.scryfall.com/cards/" + card.Set + "/" + card.Num
	getCardName(url)
}

func getCardName(url string) *Card {
	c := Card{}
	// Create a new HTTP client
	client := &http.Client{
		Timeout: time.Second * 10, // Timeout each requests
	}
	method := "GET"
	payload := strings.NewReader(`{"key1":"value1", "key2":"value2"}`)

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return &c
	}
	// Set headers
	req.Header.Add("Content-Type", "application/json")

	// Execute the request using the custom HTTP client
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return &c
	}
	defer resp.Body.Close()

	// Read and print response
	err = json.NewDecoder(resp.Body).Decode(&c)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return &c
	}
	fmt.Println(c)
	c.Count += 1

	//path := 1
	//fmt.Println("MtG Set: ", c.Set, "Card number: ", c.Num)
	return &c
}
