package marketData

import "time"

type MarketData struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Expiry    time.Time `json:"expiry"`
}
