package repositories

import (
    "database/sql"
    "APIs/src/max30102/domain/entities"
)

type Max30102RepositoryMySQL struct {
    db *sql.DB
}

func NewMax30102RepositoryMySQL(db *sql.DB) *Max30102RepositoryMySQL {
    return &Max30102RepositoryMySQL{db: db}
}

func (r *Max30102RepositoryMySQL) Save(data *entities.Max30102) error {
    query := `INSERT INTO max30102 (bpm, spo2) VALUES (?, ?)`
    _, err := r.db.Exec(query, data.BPM, data.SpO2)
    return err
}
