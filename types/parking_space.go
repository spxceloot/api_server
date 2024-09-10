package types

import "time"

type ParkingSpace struct {
	ID            string      `json:"id"`
	GeoLocation   GeoLocation `json:"location"`
	Address       Address     `json:"address"`
	ImageUrls     []string    `json:"image_urls"`
	CoverImageUrl string      `json:"cover_image_url"`
	CreatedAt     time.Time   `json:"created_at"`
}
