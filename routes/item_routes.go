package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Item struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Value int    `json:"value"`
}

var items = []Item{}

func CreateItemHandler(w http.ResponseWriter, r *http.Request) {
	var item Item

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	item.ID = fmt.Sprintf("%d", len(items)+1)
	items = append(items, item)

	fmt.Printf("%v\n\n", items)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func GetItemsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.MarshalIndent(items, "", "  ")
	json.NewEncoder(w).Encode(items)
}
