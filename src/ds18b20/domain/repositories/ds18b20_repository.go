package repositories

import "APIs/src/ds18b20/domain/entities"

type DS18B20Repository interface {
    Save(data *entities.DS18B20) error
}