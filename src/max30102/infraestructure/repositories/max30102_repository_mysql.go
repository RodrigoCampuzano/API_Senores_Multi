package repositories

import (
    "database/sql"
    "APIs/src/max30102/domain/entities"
)

type Max30102Repository interface {
    Save(data *entities.Max30102) error
    GetAll() ([]entities.Max30102, error)
}

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

func (r *Max30102RepositoryMySQL) GetAll() ([]entities.Max30102, error) {
    query := `SELECT id, bpm, spo2 FROM max30102`
    rows, err := r.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var results []entities.Max30102
    for rows.Next() {
        var data entities.Max30102
        if err := rows.Scan(&data.ID, &data.BPM, &data.SpO2); err != nil {
            return nil, err
        }
        results = append(results, data)
    }
    return results, nil
}