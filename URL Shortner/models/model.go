package models

import "time"

type URL struct {
	URL          string    `json:"url"`
	Short_URL    string    `json:"short_url"`
	Created_Date time.Time `json:"created_date"`
}
