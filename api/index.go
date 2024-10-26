// main.go
package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// Struct untuk menyimpan data
type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Slice untuk menyimpan data dan mutex untuk akses bersamaan
var items []Item
var nextID = 1
var mu sync.Mutex

// Handler untuk mendapatkan semua item
func GetItemsHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// Handler untuk mendapatkan item berdasarkan ID
func GetItemHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	id := r.URL.Query().Get("id")
	for _, item := range items {
		if fmt.Sprintf("%d", item.ID) == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.NotFound(w, r)
}

// Handler untuk membuat item baru
func CreateItemHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var newItem Item
	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newItem.ID = nextID
	nextID++
	items = append(items, newItem)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)
}

// Handler untuk memperbarui item berdasarkan ID
func UpdateItemHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	id := r.URL.Query().Get("id")
	for i, item := range items {
		if fmt.Sprintf("%d", item.ID) == id {
			var updatedItem Item
			if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			updatedItem.ID = item.ID
			items[i] = updatedItem
			json.NewEncoder(w).Encode(updatedItem)
			return
		}
	}
	http.NotFound(w, r)
}

// Handler untuk menghapus item berdasarkan ID
func DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	id := r.URL.Query().Get("id")
	for i, item := range items {
		if fmt.Sprintf("%d", item.ID) == id {
			items = append(items[:i], items[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.NotFound(w, r)
}

// Fungsi utama tidak diperlukan di sini untuk Vercel