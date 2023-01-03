package main

import (
	"sync"

	"github.com/paveltyukin/practicum-go-service-devops/internal"
)

type mxMetrics struct {
	mx      sync.Mutex
	metrics internal.Metrics
}

func newMxMetrics() *mxMetrics {
	return &mxMetrics{
		metrics: internal.GetCurrentMetrics(internal.Metrics{}),
	}
}

func (m *mxMetrics) Get() internal.Metrics {
	m.mx.Lock()
	defer m.mx.Unlock()
	return m.metrics
}

func (m *mxMetrics) Update() {
	m.mx.Lock()
	defer m.mx.Unlock()
	m.metrics = internal.GetCurrentMetrics(m.metrics)
}
