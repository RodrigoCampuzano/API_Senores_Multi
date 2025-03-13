package application

import "APIs/src/ds18b20/domain/repositories"

type RespuestaService struct {
    respuestaRepo repositories.RespuestaRepository
}

func NewRespuestaService(respuestaRepo repositories.RespuestaRepository) *RespuestaService {
    return &RespuestaService{respuestaRepo: respuestaRepo}
}

func (s *RespuestaService) SendMessage(queue string, message string) error {
    return s.respuestaRepo.SendMessage(queue, message)
}