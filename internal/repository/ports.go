package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"sync"
	"test/internal/models"
	"test/pkg/db"
)

type Port struct {
	db        *gorm.DB
	batchSize int
	bufferLen int
	buffer    *sync.Map
	lock      sync.Mutex
}

func NewPorts() *Port {
	return &Port{db: db.Connect(), batchSize: 200, buffer: new(sync.Map)}
}

func (p *Port) Upsert(port *models.Port) error {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.bufferStore(port)
	if p.bufferLen >= p.batchSize {
		return p.PortsBufferFlush()
	}
	return nil
}

func (p *Port) PortsBufferFlush() error {
	ports := p.bufferRead()
	result := p.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(ports, len(ports))
	if result.Error != nil {
		log.Printf("error flushing ports buffer: %s", result.Error.Error())
		return result.Error
	}
	return nil
}

func (p *Port) bufferStore(port *models.Port) {
	p.buffer.Store(port.ID, port)
	p.bufferLen++
}

func (p *Port) bufferRead() []*models.Port {
	var ports []*models.Port
	p.buffer.Range(func(key, value any) bool {
		ports = append(ports, value.(*models.Port))
		return true
	})
	p.bufferLen = 0
	p.buffer = new(sync.Map)
	return ports
}
