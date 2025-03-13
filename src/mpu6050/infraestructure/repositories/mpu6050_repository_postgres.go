package repositories

import (
    "database/sql"
    "APIs/src/mpu6050/domain/entities"
)

type MPU6050RepositoryPostgres struct {
    db *sql.DB
}

func NewMPU6050RepositoryPostgres(db *sql.DB) *MPU6050RepositoryPostgres {
    return &MPU6050RepositoryPostgres{db: db}
}

func (r *MPU6050RepositoryPostgres) Save(data *entities.MPU6050) error {
    query := `INSERT INTO mpu6050 (, ax, ay, az, gx, gy, gz) VALUES ($1, $2, $3, $4, $5, $6, $7)`
    _, err := r.db.Exec(query, data.Ax, data.Ay, data.Az, data.Gx, data.Gy, data.Gz)
    return err
}