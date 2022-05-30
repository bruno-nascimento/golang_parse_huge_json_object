//go:generate go run github.com/golang/mock/mockgen@v1.6.0 -source=./ports.go -destination ./ports_mocks.go -package=service
package service

import (
	"test/internal/models"
	"test/internal/repository"
)

type Port struct {
	repo *repository.Port
}

func NewPort() *Port {
	return &Port{repo: repository.NewPorts()}
}

func (p Port) Add(port *models.Port) error {
	return p.repo.Upsert(port)
}

func (p Port) PortsBufferFlush() error {
	return p.repo.PortsBufferFlush()
}

type PortService interface {
	Add(port *models.Port) error
	PortsBufferFlush() error
}
