package handlers

import (
	"L0/internal/server/cache"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	imageURL := "https://sun9-30.userapi.com/impg/swfxHxUXbMgeaKSIr-rT45mVMoUdDruzhRXJBA/ieuPQAAVGNk.jpg?size=2560x1440&quality=96&sign=5b6ffe424ce5580cd9564fe47e084ee3&type=album"

	resp, err := http.Get(imageURL)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func OrdersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Printf("Failed to get wd: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	filePath := filepath.Join(wd, "..", "internal", "server", "static", "index.html")
	log.Println("File path:", filePath)

	http.ServeFile(w, r, filePath)
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
