package entities

import "time"

type MPU6050 struct {
    ID        int       `json:"id"`
    Ax        float64   `json:"ax"`
    Ay        float64   `json:"ay"`
    Az        float64   `json:"az"`
    Gx        float64   `json:"gx"`
    Gy        float64   `json:"gy"`
    Gz        float64   `json:"gz"`
    Timestamp time.Time `json:"timestamp"`
}
