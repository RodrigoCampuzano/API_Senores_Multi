package ds_entities



type DS18B20 struct {
    ID          int       `json:"id"`
    DeviceID    int       `json:"device_id"`
    Temperatura float64   `json:"temperatura"`
    Timestamp   string `json:"timestamp"`
}