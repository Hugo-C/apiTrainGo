package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

const KmPriceRatio = 0.223

type responseKmToPrice struct {
	Devise   string  `json:"devise"`
	Price    float64 `json:"price"`
	Distance float64 `json:"distance"`
}

// Return a price in the wanted devise based on the distance
func KmToPrice(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params["devise"])  // TODO handle several devise
	distance, err := strconv.ParseFloat(params["km"], 64)
	if err != nil {
		fmt.Println("bad parameters : " + params["km"])
	}
	fmt.Println(distance)

	price := KmPriceRatio * distance
	res := &responseKmToPrice{
		Devise:   params["devise"],
		Price:    price,
		Distance: distance,
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		fmt.Println(err)
	}
}

// main function to boot up everything
func main() {
    port := "8000"
	router := mux.NewRouter()
	fmt.Println("Server running on port " + port)
	router.HandleFunc("/kmToPrice/{devise}/{km:[0-9]+.[0-9]+}", KmToPrice).Methods("GET")
	log.Fatal(http.ListenAndServe(":" + port, router))
}
