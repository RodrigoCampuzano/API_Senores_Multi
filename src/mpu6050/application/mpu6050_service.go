package application

import (
    "APIs/src/mpu6050/domain/repositories"
    "APIs/src/mpu6050/domain/entities"
)

type Mpu6050Service struct {
    mpu6050Repo  repositories.MPU6050Repository
}

func NewMpu6050Service(mpu6050Repo repositories.MPU6050Repository) *Mpu6050Service {
    return &Mpu6050Service{
        mpu6050Repo:  mpu6050Repo,
    }
}

func (s *Mpu6050Service) SaveMPU6050Data(data *entities.MPU6050) error {
    return s.mpu6050Repo.Save(data)
}