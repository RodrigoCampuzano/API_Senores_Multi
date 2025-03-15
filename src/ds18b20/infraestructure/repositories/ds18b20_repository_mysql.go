package repositories

import (
    "database/sql"
    "APIs/src/ds18b20/domain/entities"
)

type DS18B20RepositoryMySQL struct {
    db *sql.DB
}

func NewDS18B20RepositoryMySQL(db *sql.DB) *DS18B20RepositoryMySQL {
    return &DS18B20RepositoryMySQL{db: db}
}

func (r *DS18B20RepositoryMySQL) Save(data *ds_entities.DS18B20) error {
    query := `INSERT INTO ds18b20 (device_id, temperatura) VALUES (?, ?)`
    _, err := r.db.Exec(query, data.DeviceID, data.Temperatura)
    return err
}
