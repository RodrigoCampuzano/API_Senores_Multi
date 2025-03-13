package entities

import "time"

type DS18B20 struct {
    ID          int       `json:"id"`
    DeviceID    int       `json:"device_id"`
    Temperatura float64   `json:"temperatura"`
    Timestamp   time.Time `json:"timestamp"`
}