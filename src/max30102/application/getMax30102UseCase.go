package application

import (
    "APIs/src/max30102/domain/entities"
    "APIs/src/max30102/domain/repositories"
)

type GetAll30102UseCase struct {
    max30102Repo repositories.Max30102Repository
}

func NewGetAllMax30102UseCase ( max30102Repo repositories.Max30102Repository ) *GetAll30102UseCase {
    return &GetAll30102UseCase {
        max30102Repo: max30102Repo,
    }
}

func (s *GetAll30102UseCase ) GetMax30102Data() ([]*entities.Max30102, error) {
    return s.max30102Repo.GetAll()
}