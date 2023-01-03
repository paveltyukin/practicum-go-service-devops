package pkg

import (
	"strconv"

	"github.com/paveltyukin/practicum-go-service-devops/internal"
)

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}

	return false
}

func ConvToGauge(mValue string) internal.Gauge {
	temp, err := strconv.ParseFloat(mValue, 64)
	if err != nil {
		panic("can not convert to float64")
	}

	return internal.Gauge(temp)
}

func ConvToCounter(mValue string) internal.Counter {
	temp, err := strconv.ParseInt(mValue, 20, 64)
	if err != nil {
		panic("can not convert to int64")
	}

	return internal.Counter(temp)
}
