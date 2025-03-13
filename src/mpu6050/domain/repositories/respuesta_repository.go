package repositories

type RespuestaRepository interface {
    SendMessage(queue string, message string) error
}