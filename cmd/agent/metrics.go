package main

import (
	"math/rand"
	"runtime"
	"sync"
	"time"
)

type gauge float64
type counter int64

type metrics struct {
	Alloc         gauge
	BuckHashSys   gauge
	Frees         gauge
	GCCPUFraction gauge
	GCSys         gauge
	HeapAlloc     gauge
	HeapIdle      gauge
	HeapInuse     gauge
	HeapObjects   gauge
	HeapReleased  gauge
	HeapSys       gauge
	LastGC        gauge
	Lookups       gauge
	MCacheInuse   gauge
	MCacheSys     gauge
	MSpanInuse    gauge
	MSpanSys      gauge
	Mallocs       gauge
	NextGC        gauge
	NumForcedGC   gauge
	NumGC         gauge
	OtherSys      gauge
	PauseTotalNs  gauge
	StackInuse    gauge
	StackSys      gauge
	Sys           gauge
	TotalAlloc    gauge
	PollCount     counter
	RandomValue   gauge
}

func getCurrentMetrics(metrics metrics) metrics {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	metrics.Alloc = gauge(memStats.Alloc)
	metrics.BuckHashSys = gauge(memStats.BuckHashSys)
	metrics.Frees = gauge(memStats.Frees)
	metrics.GCCPUFraction = gauge(memStats.GCCPUFraction)
	metrics.GCSys = gauge(memStats.GCSys)
	metrics.HeapAlloc = gauge(memStats.HeapAlloc)
	metrics.HeapIdle = gauge(memStats.HeapIdle)
	metrics.HeapInuse = gauge(memStats.HeapInuse)
	metrics.HeapObjects = gauge(memStats.HeapObjects)
	metrics.HeapReleased = gauge(memStats.HeapReleased)
	metrics.HeapSys = gauge(memStats.HeapSys)
	metrics.LastGC = gauge(memStats.LastGC)
	metrics.Lookups = gauge(memStats.Lookups)
	metrics.MCacheInuse = gauge(memStats.MCacheInuse)
	metrics.MCacheSys = gauge(memStats.MCacheSys)
	metrics.MSpanInuse = gauge(memStats.MSpanInuse)
	metrics.MSpanSys = gauge(memStats.MSpanSys)
	metrics.Mallocs = gauge(memStats.Mallocs)
	metrics.NextGC = gauge(memStats.NextGC)
	metrics.NumForcedGC = gauge(memStats.NumForcedGC)
	metrics.NumGC = gauge(memStats.NumGC)
	metrics.OtherSys = gauge(memStats.OtherSys)
	metrics.PauseTotalNs = gauge(memStats.PauseTotalNs)
	metrics.StackInuse = gauge(memStats.StackInuse)
	metrics.StackSys = gauge(memStats.StackSys)
	metrics.Sys = gauge(memStats.Sys)
	metrics.TotalAlloc = gauge(memStats.TotalAlloc)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	metrics.RandomValue = gauge(r1.Float64())
	metrics.PollCount = metrics.PollCount + 1

	return metrics
}

type mxMetrics struct {
	mx      sync.Mutex
	metrics metrics
}

func newMxMetrics() *mxMetrics {
	return &mxMetrics{
		metrics: getCurrentMetrics(metrics{}),
	}
}

func (m *mxMetrics) Get() metrics {
	m.mx.Lock()
	defer m.mx.Unlock()
	return m.metrics
}

func (m *mxMetrics) Update() {
	m.mx.Lock()
	defer m.mx.Unlock()
	m.metrics = getCurrentMetrics(m.metrics)
}
