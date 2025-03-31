package repositories

import (
    "APIs/src/ds18b20/domain/entities"
)

type DS18B20RepositoryStub struct{}

func NewDS18B20RepositoryStub() *DS18B20RepositoryStub {
    return &DS18B20RepositoryStub{}
}

func (r *DS18B20RepositoryStub) Save(data *ds_entities.DS18B20) error {
    // Simula el guardado de datos sin hacer nada
    return nil
}