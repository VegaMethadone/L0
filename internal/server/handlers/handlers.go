package handlers

import (
	"L0/internal/server/cache"
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	imagePath := "img/home.jpg"

	imageFile, err := os.Open(imagePath)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer imageFile.Close()

	w.Header().Set("Content-Type", "image/jpeg")

	http.ServeContent(w, r, filepath.Base(imagePath), time.Time{}, imageFile)
}

func GetByIdOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	orderId := vars["id"]

	value, err := cache.Cache.Get(orderId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonData, err := json.Marshal(value)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}
