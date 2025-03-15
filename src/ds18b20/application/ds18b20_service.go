package application

import (
    "APIs/src/ds18b20/domain/entities"
    "APIs/src/ds18b20/domain/repositories"
)

type Ds18b20Service struct {
    ds18b20Repo  repositories.DS18B20Repository
}

func NewDS18B20Service(ds18b20Repo repositories.DS18B20Repository) *Ds18b20Service {
    return &Ds18b20Service{
        ds18b20Repo:  ds18b20Repo,
    }
}

func (s *Ds18b20Service) SaveDS18B20Data(data *ds_entities.DS18B20) error {
    return s.ds18b20Repo.Save(data)
}