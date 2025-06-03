package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"sensors/db"
	"sensors/models"

	"github.com/google/uuid"
)

func CreateTemperature(w http.ResponseWriter, r *http.Request) {
	var t models.Temperature
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	t.ID = uuid.NewString()

	_, err := db.DB.Exec("INSERT INTO temperature (id, temperature, date, time, location) VALUES (?, ?, ?, ?, ?)",
		t.ID, t.Temperature, t.Date, t.Time, t.Location)

	if err != nil {
		http.Error(w, "DB insert error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(t)
}

func GetTemperature(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}
	limit := 10
	offset := (page - 1) * limit

	// Count total records
	var totalRecords int
	err := db.DB.QueryRow("SELECT COUNT(*) FROM temperature").Scan(&totalRecords)
	if err != nil {
		http.Error(w, "DB count error", http.StatusInternalServerError)
		return
	}

	// Query paginated data
	rows, err := db.DB.Query("SELECT id, temperature, date, time, location FROM temperature ORDER BY date DESC, time DESC LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		http.Error(w, "DB read error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var results []models.Temperature
	for rows.Next() {
		var t models.Temperature
		if err := rows.Scan(&t.ID, &t.Temperature, &t.Date, &t.Time, &t.Location); err != nil {
			http.Error(w, "DB scan error", http.StatusInternalServerError)
			return
		}
		results = append(results, t)
	}

	totalPages := (totalRecords + limit - 1) / limit // ceil

	// Build response object
	response := map[string]interface{}{
		"data":        results,
		"page":        page,
		"total_pages": totalPages,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
