package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Coupon Structure
type Coupon struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Brand   string  `json:"brand"`
	Value   float64 `json:"value"`
	Created string  `json:"createdAt"`
	Expiry  string  `json:"expiry"`
}

// Coupons can be also exported and used outside the Package main
var Coupons []Coupon

// Delete All Coupons Endpoint
func getCouponsEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: All Coupons Endpoint")

	for _, item := range Coupons {
		json.NewEncoder(w).Encode(item)
	}
}

// Get Specific Coupon Endpoint
func getCouponEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Simple Coupon Endpoint")
	params := mux.Vars(r)

	for _, item := range Coupons {
		if item.ID == params["id"] { // Can use brand etc.
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Coupon{})
}

// Create Coupon Endpoint
func createCouponEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Create Coupon Endpoint")
	params := mux.Vars(r)
	var coupon Coupon

	_ = json.NewDecoder(r.Body).Decode(&coupon)

	coupon.ID = params["id"]
	tempTime := time.Now()
	coupon.Created = tempTime.Format("2018-03-01 10:15:53")
	coupon.Expiry = tempTime.AddDate(0, 3, 0).Format("2018-03-01 10:15:53") //expire in 3 months from creation
	Coupons = append(Coupons, coupon)

	json.NewEncoder(w).Encode(Coupons)
}

// Delete Coupon Endpoint
func deleteCouponEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Delete Coupon Endpoint")
	params := mux.Vars(r)

	for index, item := range Coupons {
		if item.ID == params["id"] { // Can use brand etc.
			Coupons = append(Coupons[:index], Coupons[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Coupons)
}

// The homepage endpoint
func homePageEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

// Handler for requests
func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePageEndpoint)
	router.HandleFunc("/coupons", getCouponsEndpoint).Methods("GET")
	router.HandleFunc("/coupons/{id}", getCouponEndpoint).Methods("GET")
	router.HandleFunc("/coupons/{id}", createCouponEndpoint).Methods("POST")
	router.HandleFunc("/coupons/{id}", deleteCouponEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", router))
}

// Main body running
func main() {
	Coupons = append(Coupons, Coupon{ID: "1", Name: "Save £20 at Tesco", Brand: "Tesco", Value: 20.0, Created: "2018-03-01 10:15:53", Expiry: "2019-03-01 10:15:53"})
	Coupons = append(Coupons, Coupon{ID: "2", Name: "15% off at Booking.com", Brand: "Booking.com", Value: 0.15, Created: "2018-03-01 10:15:53", Expiry: "2019-03-01 10:15:53"})
	handleRequests()
}
