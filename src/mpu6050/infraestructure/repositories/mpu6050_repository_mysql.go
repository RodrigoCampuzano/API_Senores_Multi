package repositories

import (
    "database/sql"
    "APIs/src/mpu6050/domain/entities"
)

type MPU6050RepositoryMySQL struct {
    db *sql.DB
}

func NewMPU6050RepositoryMySQL(db *sql.DB) *MPU6050RepositoryMySQL {
    return &MPU6050RepositoryMySQL{db: db}
}

func (r *MPU6050RepositoryMySQL) Save(data *entities.MPU6050) error {
    query := `INSERT INTO mpu6050 (ax, ay, az, gx, gy, gz) VALUES (?, ?, ?, ?, ?, ?)`
    _, err := r.db.Exec(query, data.Ax, data.Ay, data.Az, data.Gx, data.Gy, data.Gz)
    return err
}