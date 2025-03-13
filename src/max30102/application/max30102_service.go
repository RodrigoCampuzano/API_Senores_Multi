package application

import (
    "APIs/src/max30102/domain/entities"
    "APIs/src/max30102/domain/repositories"
)

type Max30102Service struct {
    max30102Repo repositories.Max30102Repository
}

func NewMax30102Service(max30102Repo repositories.Max30102Repository) *Max30102Service {
    return &Max30102Service{
        max30102Repo: max30102Repo,
    }
}

func (s *Max30102Service) SaveMax30102Data(data *entities.Max30102) error {
    return s.max30102Repo.Save(data)
}

func (s *Max30102Service) GetMax30102Data() ([]entities.Max30102, error) {
    return s.max30102Repo.GetAll()
}