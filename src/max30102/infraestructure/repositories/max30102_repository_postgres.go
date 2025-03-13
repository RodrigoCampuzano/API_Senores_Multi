package repositories

import (
    "database/sql"
    "APIs/src/max30102/domain/entities"
)

type Max30102RepositoryPostgres struct {
    db *sql.DB
}

func NewMax30102RepositoryPostgres(db *sql.DB) *Max30102RepositoryPostgres {
    return &Max30102RepositoryPostgres{db: db}
}

func (r *Max30102RepositoryPostgres) Save(data *entities.Max30102) error {
    query := `INSERT INTO max30102 (bpm, spo2) VALUES ($1, $2)`
    _, err := r.db.Exec(query, data.BPM, data.SpO2)
    return err
}