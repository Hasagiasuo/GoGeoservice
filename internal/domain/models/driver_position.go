package models

import "time"

type DriverPosition struct {
	Id        int       `json:"position_id" db:"id"`
	DriverId  int       `json:"driver_id" db:"driver_id"`
	Latitude  float64   `json:"lat" db:"lat"`
	Longitude float64   `json:"lon" db:"lon"`
	Timestamp time.Time `db:"timestamp"`
}
