package repositories

import "APIs/src/max30102/domain/entities"

type Max30102Repository interface {
    Save(data *entities.Max30102) error
    GetAll() ([]entities.Max30102, error) // Agregar esta línea
}
