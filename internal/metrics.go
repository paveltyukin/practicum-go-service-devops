package internal

import (
	"math/rand"
	"runtime"
	"time"
)

type Gauge float64
type Counter int64

type Metrics struct {
	Alloc         Gauge
	BuckHashSys   Gauge
	Frees         Gauge
	GCCPUFraction Gauge
	GCSys         Gauge
	HeapAlloc     Gauge
	HeapIdle      Gauge
	HeapInuse     Gauge
	HeapObjects   Gauge
	HeapReleased  Gauge
	HeapSys       Gauge
	LastGC        Gauge
	Lookups       Gauge
	MCacheInuse   Gauge
	MCacheSys     Gauge
	MSpanInuse    Gauge
	MSpanSys      Gauge
	Mallocs       Gauge
	NextGC        Gauge
	NumForcedGC   Gauge
	NumGC         Gauge
	OtherSys      Gauge
	PauseTotalNs  Gauge
	StackInuse    Gauge
	StackSys      Gauge
	Sys           Gauge
	TotalAlloc    Gauge
	PollCount     Counter
	RandomValue   Gauge
}

func GetCurrentMetrics(metrics Metrics) Metrics {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	metrics.Alloc = Gauge(memStats.Alloc)
	metrics.BuckHashSys = Gauge(memStats.BuckHashSys)
	metrics.Frees = Gauge(memStats.Frees)
	metrics.GCCPUFraction = Gauge(memStats.GCCPUFraction)
	metrics.GCSys = Gauge(memStats.GCSys)
	metrics.HeapAlloc = Gauge(memStats.HeapAlloc)
	metrics.HeapIdle = Gauge(memStats.HeapIdle)
	metrics.HeapInuse = Gauge(memStats.HeapInuse)
	metrics.HeapObjects = Gauge(memStats.HeapObjects)
	metrics.HeapReleased = Gauge(memStats.HeapReleased)
	metrics.HeapSys = Gauge(memStats.HeapSys)
	metrics.LastGC = Gauge(memStats.LastGC)
	metrics.Lookups = Gauge(memStats.Lookups)
	metrics.MCacheInuse = Gauge(memStats.MCacheInuse)
	metrics.MCacheSys = Gauge(memStats.MCacheSys)
	metrics.MSpanInuse = Gauge(memStats.MSpanInuse)
	metrics.MSpanSys = Gauge(memStats.MSpanSys)
	metrics.Mallocs = Gauge(memStats.Mallocs)
	metrics.NextGC = Gauge(memStats.NextGC)
	metrics.NumForcedGC = Gauge(memStats.NumForcedGC)
	metrics.NumGC = Gauge(memStats.NumGC)
	metrics.OtherSys = Gauge(memStats.OtherSys)
	metrics.PauseTotalNs = Gauge(memStats.PauseTotalNs)
	metrics.StackInuse = Gauge(memStats.StackInuse)
	metrics.StackSys = Gauge(memStats.StackSys)
	metrics.Sys = Gauge(memStats.Sys)
	metrics.TotalAlloc = Gauge(memStats.TotalAlloc)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	metrics.RandomValue = Gauge(r1.Float64())
	metrics.PollCount = metrics.PollCount + 1

	return metrics
}
