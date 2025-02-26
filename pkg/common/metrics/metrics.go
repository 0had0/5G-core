package metrics

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/0had0/5G-core/pkg/common/logger"
	"go.uber.org/zap"
)

var (
	// RequestCounter counts the number of requests processed
	RequestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "requests_total",
			Help: "Total number of requests processed",
		},
		[]string{"service", "method", "status"},
	)

	// RequestDuration tracks the duration of requests
	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "request_duration_seconds",
			Help:    "Duration of requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"service", "method"},
	)

	// ActiveConnections tracks the number of active connections
	ActiveConnections = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "active_connections",
			Help: "Number of active connections",
		},
		[]string{"service"},
	)

	// ServiceRegistrations tracks the number of service registrations with NRF
	ServiceRegistrations = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "service_registrations_total",
			Help: "Total number of service registrations with NRF",
		},
		[]string{"service", "status"},
	)

	// ServiceDiscoveries tracks the number of service discoveries from NRF
	ServiceDiscoveries = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "service_discoveries_total",
			Help: "Total number of service discoveries from NRF",
		},
		[]string{"service", "target_service", "status"},
	)
)

// Initialize sets up the Prometheus metrics and starts the metrics server
func Initialize(serviceName string, metricsPort int) {
	// Register the metrics
	prometheus.MustRegister(RequestCounter)
	prometheus.MustRegister(RequestDuration)
	prometheus.MustRegister(ActiveConnections)
	prometheus.MustRegister(ServiceRegistrations)
	prometheus.MustRegister(ServiceDiscoveries)

	// Start the metrics server
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		address := fmt.Sprintf(":%d", metricsPort)
		logger.Info("Starting metrics server", zap.String("address", address))
		err := http.ListenAndServe(address, nil)
		if err != nil {
			logger.Error("Metrics server error", zap.Error(err))
		}
	}()

	logger.Info("Metrics initialized", zap.String("service", serviceName))
}
