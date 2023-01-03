package main

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/paveltyukin/practicum-go-service-devops/internal"
)

type reportParams struct {
	mType  string
	mName  string
	mValue string
}

func newHTTPClient() *http.Client {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100

	return &http.Client{
		Timeout:   10 * time.Second,
		Transport: t,
	}
}

func updateMetrics(ctx context.Context, mxMetrics *mxMetrics) {
	pollInterval := time.NewTicker(2 * time.Second)
	defer pollInterval.Stop()

	for {
		select {
		case <-pollInterval.C:
			fmt.Println(mxMetrics)
			mxMetrics.Update()
		case <-ctx.Done():
			return
		}
	}
}

func sendMetrics(ctx context.Context, m *mxMetrics, client *http.Client) {
	reportInterval := time.NewTicker(10 * time.Second)
	defer reportInterval.Stop()

	for {
		select {
		case <-reportInterval.C:
			send(m, client)
		case <-ctx.Done():
			return
		}
	}
}

func send(m *mxMetrics, client *http.Client) {
	params := reportParams{}
	curMetrics := m.Get()

	v := reflect.ValueOf(curMetrics)
	metricTypes := v.Type()

	for i := 0; i < v.NumField(); i++ {
		switch v.Field(i).Interface().(type) {
		case internal.Gauge:
			params.mType = "gauge"
		case internal.Counter:
			params.mType = "counter"
		default:
			panic("Metrics error types")
		}

		params.mValue = fmt.Sprintf("%v", v.Field(i).Interface())
		params.mName = fmt.Sprintf("%v", metricTypes.Field(i).Name)

		err := sendMetricsToServer(params, client)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func sendMetricsToServer(params reportParams, client *http.Client) error {
	var res *http.Response

	url := fmt.Sprintf("http://127.0.0.1:8080/update/%v/%v/%v", params.mType, params.mName, params.mValue)
	request, err := http.NewRequest(http.MethodPost, url, nil)
	request.Header.Set("Content-Type", "text/plain")
	if err != nil {
		return err
	}

	res, err = client.Do(request)
	if err != nil {
		return err
	}

	err = res.Body.Close()
	if err != nil {
		return err
	}

	fmt.Println(request.URL, request.Method, params)

	return nil
}
