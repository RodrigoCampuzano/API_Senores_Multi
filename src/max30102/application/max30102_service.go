package application

import (
    "APIs/src/max30102/domain/entities"
    "APIs/src/max30102/domain/repositories"
)

type Max3102Service struct {
    max30102Repo repositories.Max30102Repository
}

func NewMax30102Service(max30102Repo repositories.Max30102Repository) *Max3102Service {
    return &Max3102Service{
        max30102Repo: max30102Repo,
    }
}

func (s *Max3102Service) SaveMax30102Data(data *entities.Max30102) error {
    return s.max30102Repo.Save(data)
}