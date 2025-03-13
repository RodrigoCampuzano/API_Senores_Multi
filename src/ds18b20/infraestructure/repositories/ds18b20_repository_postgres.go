package repositories

import (
    "database/sql"
    "APIs/src/ds18b20/domain/entities"
)

type DS18B20RepositoryPostgres struct {
    db *sql.DB
}

func NewDS18B20RepositoryPostgres(db *sql.DB) *DS18B20RepositoryPostgres {
    return &DS18B20RepositoryPostgres{db: db}
}

func (r *DS18B20RepositoryPostgres) Save(data *entities.DS18B20) error {
    query := `INSERT INTO ds18b20 (device_id, temperatura) VALUES ($1, $2)`
    _, err := r.db.Exec(query, data.DeviceID, data.Temperatura)
    return err
}