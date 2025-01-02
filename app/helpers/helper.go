package helpers

import (
	"log"
	"time"
)

func GenerateTimeLocation() *time.Location {
	// Load lokasi untuk zona waktu Jakarta
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Fatal(err)
	}
	return location
}
