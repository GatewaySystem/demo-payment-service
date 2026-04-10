package store

import (
	"sync"

	"github.com/GatewaySystem/demo-payment-service/models"
)

type PaymentStore struct {
	mu       sync.RWMutex
	payments map[string]*models.Payment
}

var instance *PaymentStore
var once sync.Once

func Get() *PaymentStore {
	once.Do(func() {
		instance = &PaymentStore{
			payments: make(map[string]*models.Payment),
		}
	})
	return instance
}

func (s *PaymentStore) Save(p *models.Payment) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.payments[p.ID] = p
}

func (s *PaymentStore) Find(id string) *models.Payment {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.payments[id]
}

func (s *PaymentStore) List() []*models.Payment {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make([]*models.Payment, 0, len(s.payments))
	for _, p := range s.payments {
		result = append(result, p)
	}
	return result
}
