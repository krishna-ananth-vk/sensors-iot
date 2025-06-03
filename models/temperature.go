package models

type Temperature struct {
	ID          string  `json:"id"`
	Temperature float64 `json:"temperature"`
	Date        string  `json:"date"` // Format: YYYY-MM-DD
	Time        string  `json:"time"` // Format: HH:MM:SS
	Location    string  `json:"location"`
}
