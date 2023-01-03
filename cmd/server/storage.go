package main

import (
	"github.com/paveltyukin/practicum-go-service-devops/internal"
	"github.com/paveltyukin/practicum-go-service-devops/pkg"
)

type MemStorage struct {
	metrics internal.Metrics
}

func (s *MemStorage) Set(mType, mValue string) {
	switch mType {
	case "Alloc":
		s.metrics.Alloc = pkg.ConvToGauge(mValue)
	case "BuckHashSys":
		s.metrics.BuckHashSys = pkg.ConvToGauge(mValue)
	case "Frees":
		s.metrics.Frees = pkg.ConvToGauge(mValue)
	case "GCCPUFraction":
		s.metrics.GCCPUFraction = pkg.ConvToGauge(mValue)
	case "GCSys":
		s.metrics.GCSys = pkg.ConvToGauge(mValue)
	case "HeapAlloc":
		s.metrics.HeapAlloc = pkg.ConvToGauge(mValue)
	case "HeapIdle":
		s.metrics.HeapIdle = pkg.ConvToGauge(mValue)
	case "HeapInuse":
		s.metrics.HeapInuse = pkg.ConvToGauge(mValue)
	case "HeapObjects":
		s.metrics.HeapObjects = pkg.ConvToGauge(mValue)
	case "HeapReleased":
		s.metrics.HeapReleased = pkg.ConvToGauge(mValue)
	case "HeapSys":
		s.metrics.HeapSys = pkg.ConvToGauge(mValue)
	case "LastGC":
		s.metrics.LastGC = pkg.ConvToGauge(mValue)
	case "Lookups":
		s.metrics.Lookups = pkg.ConvToGauge(mValue)
	case "MCacheInuse":
		s.metrics.MCacheInuse = pkg.ConvToGauge(mValue)
	case "MCacheSys":
		s.metrics.MCacheSys = pkg.ConvToGauge(mValue)
	case "MSpanInuse":
		s.metrics.MSpanInuse = pkg.ConvToGauge(mValue)
	case "MSpanSys":
		s.metrics.MSpanSys = pkg.ConvToGauge(mValue)
	case "Mallocs":
		s.metrics.Mallocs = pkg.ConvToGauge(mValue)
	case "NextGC":
		s.metrics.NextGC = pkg.ConvToGauge(mValue)
	case "NumForcedGC":
		s.metrics.NumForcedGC = pkg.ConvToGauge(mValue)
	case "NumGC":
		s.metrics.NumGC = pkg.ConvToGauge(mValue)
	case "OtherSys":
		s.metrics.OtherSys = pkg.ConvToGauge(mValue)
	case "PauseTotalNs":
		s.metrics.PauseTotalNs = pkg.ConvToGauge(mValue)
	case "StackInuse":
		s.metrics.StackInuse = pkg.ConvToGauge(mValue)
	case "StackSys":
		s.metrics.StackSys = pkg.ConvToGauge(mValue)
	case "Sys":
		s.metrics.Sys = pkg.ConvToGauge(mValue)
	case "TotalAlloc":
		s.metrics.TotalAlloc = pkg.ConvToGauge(mValue)
	case "PollCount":
		s.metrics.PollCount = pkg.ConvToCounter(mValue)
	case "RandomValue":
		s.metrics.RandomValue = pkg.ConvToGauge(mValue)
	default:
		panic("not found type metrics")
	}
}
