package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var metrics map[string]prometheus.Collector

func init() {
	metrics = make(map[string]prometheus.Collector)
}

// Define Metrics with config
func DefineMetrics(config *Config) {
	for _, dataPoint := range config.DataPoints {
		switch dataPoint.Type {
		case "gauge":
			gauge := prometheus.NewGauge(prometheus.GaugeOpts{
				Name: dataPoint.Name,
			})
			prometheus.MustRegister(gauge)
			metrics[dataPoint.Name] = gauge
		case "histogram":
			histogram := prometheus.NewHistogram(prometheus.HistogramOpts{
				Name: dataPoint.Name,
			})
			prometheus.MustRegister(histogram)
			metrics[dataPoint.Name] = histogram
		case "counter":
			counter := prometheus.NewCounter(prometheus.CounterOpts{
				Name: dataPoint.Name,
			})
			prometheus.MustRegister(counter)
			metrics[dataPoint.Name] = counter
		default:
			fmt.Printf("Unknown metric type: %s\n", dataPoint.Type)
		}
	}
}

// UpdateInt32Metrics updates the metrics with the given value
func UpdateInt32Metrics(dataPointName string, value int32) error {
	metric, ok := metrics[dataPointName]
	if !ok {
		return fmt.Errorf("metric %s not found", dataPointName)
	}
	switch metric.(type) {
	case prometheus.Gauge:
		metric.(prometheus.Gauge).Set(float64(value))
	case prometheus.Histogram:
		metric.(prometheus.Histogram).Observe(float64(value))
	case prometheus.Counter:
		metric.(prometheus.Counter).Add(float64(value))
	}
	return nil
}

// ExposeMetrics exposes the metrics to the prometheus server
func ExposeMetrics(config *Config) {

	http.Handle("/metrics", promhttp.Handler())
	fmt.Printf("Starting server on port %d\n", config.HttpServer.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", config.HttpServer.Port), nil)

}
