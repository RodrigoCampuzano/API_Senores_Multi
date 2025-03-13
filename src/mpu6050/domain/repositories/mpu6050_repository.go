package repositories

import "APIs/src/mpu6050/domain/entities"

type MPU6050Repository interface {
    Save(data *entities.MPU6050) error
}