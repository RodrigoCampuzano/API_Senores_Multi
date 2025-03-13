package entities

import "time"

type Max30102 struct {
    ID        int       `json:"id"`
    BPM       int       `json:"bpm"`
    SpO2      float64   `json:"spo2"`
    Timestamp time.Time `json:"timestamp"`
}