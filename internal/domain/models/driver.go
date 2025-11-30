package models

type Driver struct {
	Id     int    `json:"driver_id" db:"id"`
	Name   string `json:"driver_name" db:"name"`
	IsBusy bool   `json:"is_busy" db:"is_busy"`
}

type DriverState int

const (
	IDLE DriverState = iota
	BOOKING
	BUSY
)
