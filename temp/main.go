package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// Support for SQLite3
	_ "github.com/mattn/go-sqlite3"
)

// Coupon Structure
type Coupon struct {
	Name    string  `json:"name"`
	Brand   string  `json:"brand"`
	Value   float64 `json:"value"`
	Created string  `json:"createdAt"`
	Expiry  string  `json:"expiry"`
}

// Coupons can be also exported and used outside the Package main
type Coupons []Coupon

func databaseQuery() [][]byte {
	var fetchedCoupons [][]byte

	conn, err := sql.Open("sqlite3", "coupons.db")
	if err != nil {
		log.Println("Failed to connect to database:")
		log.Println(err)
	}

	rows, err := conn.Query("SELECT coupon FROM coupons")
	if err != nil {
		log.Println("Failed to query settings database:")
		log.Println(err)
	}

	var res []byte

	for rows.Next() {
		err = rows.Scan(&res)
		fetchedCoupons = append(fetchedCoupons, res)
	}

	return fetchedCoupons
}

func allCoupons(w http.ResponseWriter, r *http.Request) {
	fetchedCoupons := databaseQuery()
	var coupons Coupon

	for _, val := range fetchedCoupons {
		json.Unmarshal(val, &Coupon)
	}
	coupons := Coupons{
		Coupon{Name: "Save £20 at Tesco", Brand: "Tesco", Value: 20.0, Created: "2018-03-01 10:15:53", Expiry: "2019-03-01 10:15:53"},
		Coupon{Name: "Save £20 at Tesco", Brand: "Tesco", Value: 20.0, Created: "2018-03-01 10:15:53", Expiry: "2019-03-01 10:15:53"},
	}

	fmt.Println("Endpoint Hit: All Coupons Endpoint")
	json.NewEncoder(w).Encode(coupons)
}

// The homepage endpoint
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/coupons", allCoupons)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
