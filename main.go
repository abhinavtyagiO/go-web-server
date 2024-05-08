package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollars float64

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

// Add Item
func (db database) add(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if _, ok := db[item]; ok {
		message := fmt.Sprintf("Duplicate item found: %q", item)
		http.Error(w, message, http.StatusNotFound)
		return
	}

	p, err := strconv.ParseFloat(price, 32)

	if err != nil {
		message := fmt.Sprintf("Invalid Price: %q", price)
		http.Error(w, message, http.StatusBadRequest)
		return
	}

	db[item] = dollars(p)
	fmt.Fprintf(w, "Successfully added new item: %s for %s\n", item, db[item])

}

// Edit Item
func (db database) edit(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if _, ok := db[item]; !ok {
		message := fmt.Sprintf("Item not found: %q", item)
		http.Error(w, message, http.StatusNotFound)
		return
	}

	p, err := strconv.ParseFloat(price, 32)

	if err != nil {
		message := fmt.Sprintf("Invalid Price: %q", price)
		http.Error(w, message, http.StatusBadRequest)
		return
	}

	db[item] = dollars(p)
	fmt.Fprintf(w, "Successfully edited item: %s for %s\n", item, db[item])

}

// List Item
func (db database) list(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	if _, ok := db[item]; !ok {
		message := fmt.Sprintf("Item not found: %q", item)
		http.Error(w, message, http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Item: %s for %s\n", item, db[item])

}

// Delete Item
func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	if _, ok := db[item]; !ok {
		message := fmt.Sprintf("Item not found: %q", item)
		http.Error(w, message, http.StatusNotFound)
		return
	}

	delete(db, item)
	fmt.Fprintf(w, "Deleted Item: %s\n", item)

}

type database map[string]dollars

func main() {
	db := database{
		"shoes": 50,
		"socks": 5,
	}

	http.HandleFunc("/add", db.add)
	http.HandleFunc("/edit", db.edit)
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/delete", db.delete)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
