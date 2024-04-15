package metrics

import (
	_ "net/http/pprof"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"route256.ozon.ru/project/cart/internal/config"
)

var requestsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
	Subsystem: strings.Replace(config.AppName, "-", "_", -1),
	Name:      "requests_total",
	Help:      "Total amount of requests made to handler. Example: rate(requests_total[1m])",
}, []string{"method", "url"})

var responseCode = promauto.NewCounterVec(prometheus.CounterOpts{
	Subsystem: strings.Replace(config.AppName, "-", "_", -1),
	Name:      "response_code",
	Help:      "Response code of handlers. Example: rate(response_code[1m])",
}, []string{"method", "url", "code"})

var responseTime = promauto.NewHistogram(prometheus.HistogramOpts{
	Subsystem: strings.Replace(config.AppName, "-", "_", -1),
	Name:      "response_time",
	Buckets:   prometheus.DefBuckets,
	Help:      "Response time of handlers. Example: rate(response_time[1m])",
})

var externalRequestsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
	Subsystem: strings.Replace(config.AppName, "-", "_", -1),
	Name:      "external_requests_total",
	Help:      "Total amount of requests to external services. Example: rate(external_requests_total[1m])",
}, []string{"service", "handler"})

var externalResponseCode = promauto.NewCounterVec(prometheus.CounterOpts{
	Subsystem: strings.Replace(config.AppName, "-", "_", -1),
	Name:      "external_response_code",
	Help:      "Response code of requests to external services. Example: rate(external_response_code[1m])",
}, []string{"service", "handler", "code"})

var externalResponseTime = promauto.NewHistogram(prometheus.HistogramOpts{
	Subsystem: strings.Replace(config.AppName, "-", "_", -1),
	Name:      "external_response_time",
	Buckets:   prometheus.DefBuckets,
	Help:      "Response time of requests to external services. Example: rate(external_response_time[1m])",
})

var databaseRequestsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
	Subsystem: strings.Replace(config.AppName, "-", "_", -1),
	Name:      "database_requests_total",
	Help:      "Total amount of requests made to database. Example: rate(database_requests_total[1m])",
}, []string{"repository", "method", "operation"})

var databaseResponseCode = promauto.NewCounterVec(prometheus.CounterOpts{
	Subsystem: strings.Replace(config.AppName, "-", "_", -1),
	Name:      "database_response_code",
	Help:      "Response code of database request. Example: rate(external_response_code[1m])",
}, []string{"repository", "method", "operation", "code"})

var databaseResponseTime = promauto.NewHistogram(prometheus.HistogramOpts{
	Subsystem: strings.Replace(config.AppName, "-", "_", -1),
	Name:      "database_requests_time",
	Buckets:   prometheus.DefBuckets,
	Help:      "Response time of database request. Example: rate(database_response_time[1m])",
})

func UpdateRequestsTotal(method, url string) {
	requestsTotal.WithLabelValues(method, url).Inc()
}

func UpdateResponseCode(method, url, code string) {
	responseCode.WithLabelValues(method, url, code).Inc()
}

func UpdateResponseTime(start time.Time) {
	responseTime.Observe(time.Since(start).Seconds())
}

func UpdateExternalRequestsTotal(service, handler string) {
	externalRequestsTotal.WithLabelValues(service, handler).Inc()
}

func UpdateExternalResponseCode(service, handler, code string) {
	externalResponseCode.WithLabelValues(service, handler, code).Inc()
}

func UpdateExternalResponseTime(start time.Time) {
	externalResponseTime.Observe(time.Since(start).Seconds())
}

func UpdateDatabaseRequestsTotal(repository, method, operation string) {
	databaseRequestsTotal.WithLabelValues(repository, method, operation).Inc()
}

func UpdateDatabaseResponseCode(repository, method, operation, code string) {
	databaseResponseCode.WithLabelValues(repository, method, operation, code).Inc()
}

func UpdateDatabaseResponseTime(start time.Time) {
	databaseResponseTime.Observe(time.Since(start).Seconds())
}
