package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Card struct {
	Set  string `json:"set"`
	Num  string `json:"collector_number"`
	Name string `json:"name"`
	Foil bool   `json:"foil"`
}

func main() {
	fmt.Println("Starting webserver on port 8585...")
	http.HandleFunc("/submit", formHandler)
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8585", nil)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	var card Card
	// decode input or return error
	err := json.NewDecoder(r.Body).Decode(&card)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Decode error! please check your JSON formating.\n")
		return
	}
	url := "https://api.scryfall.com/cards/" + card.Set + "/" + card.Num

	cardData := getCardName(url, card.Foil)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "/")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.WriteHeader(http.StatusOK)
	resBody, _ := json.Marshal(cardData)
	_, _ = w.Write(resBody)
}

func getCardName(url string, foil bool) *Card {
	fmt.Println("URL: " + url)
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
	c.Foil = foil

	return &c
}
