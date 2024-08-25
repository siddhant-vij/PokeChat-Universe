package services

type Service interface {
	Health() map[string]string
	Close() error
}
