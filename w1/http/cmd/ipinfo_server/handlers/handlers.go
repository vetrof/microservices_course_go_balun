package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"http/cmd/ipinfo_server/db"
	"io"
	"log"
	"net/http"
)

type Response = db.IPInfo

func SelfIpHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://ipinfo.io/json")
	if err != nil {
		http.Error(w, "Failed to get IP info", http.StatusInternalServerError)
		log.Println("Error fetching IP info:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		log.Println("Error reading response:", err)
		return
	}

	var ipInfo Response
	if err := json.Unmarshal(body, &ipInfo); err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusInternalServerError)
		log.Println("Error parsing JSON:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(ipInfo); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

func ExtIpHandler(w http.ResponseWriter, r *http.Request) {

	ip := chi.URLParam(r, "ip")

	resp, err := http.Get("https://ipinfo.io/" + ip + "/json")
	if err != nil {
		http.Error(w, "Failed to get IP info", http.StatusInternalServerError)
		log.Println("Error fetching IP info:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		log.Println("Error reading response:", err)
		return
	}

	var ipInfo Response
	if err := json.Unmarshal(body, &ipInfo); err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusInternalServerError)
		log.Println("Error parsing JSON:", err)
		return
	}

	// Сохраняем в БД
	if err := db.SaveIPInfo(ipInfo); err != nil {
		log.Println("DB save error:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(ipInfo); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

func HistoryHandler(w http.ResponseWriter, r *http.Request) {
	records, err := db.HistoryIPInfo()
	if err != nil {
		http.Error(w, "Failed to fetch history", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(records)
}
