package metrics

import (
	_ "net/http/pprof"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"route256.ozon.ru/project/loms/internal/config"
)

var requestsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
	Subsystem: strings.Replace(config.AppName, "-", "_", -1),
	Name:      "requests_total",
	Help:      "Total amount of requests made to handler. Example: rate(requests_total[1m])",
}, []string{"handler"})

var responseCode = promauto.NewCounterVec(prometheus.CounterOpts{
	Subsystem: strings.Replace(config.AppName, "-", "_", -1),
	Name:      "response_code",
	Help:      "Response code of handlers. Example: rate(response_code[1m])",
}, []string{"handler", "code"})

var responseTime = promauto.NewHistogram(prometheus.HistogramOpts{
	Subsystem: strings.Replace(config.AppName, "-", "_", -1),
	Name:      "response_time",
	Buckets:   prometheus.DefBuckets,
	Help:      "Response time of handlers. Example: rate(response_time[1m])",
})

var ordersCreated = promauto.NewCounter(prometheus.CounterOpts{
	Subsystem: strings.Replace(config.AppName, "-", "_", -1),
	Name:      "orders_created",
	Help:      "Counter of created orders. Example: rate(orders_created[1m])",
})

var orderStatusChanged = promauto.NewCounterVec(prometheus.CounterOpts{
	Subsystem: strings.Replace(config.AppName, "-", "_", -1),
	Name:      "order_status_changed",
	Help:      "Counter of changed order status. Example: rate(order_status_changed[1m])",
}, []string{"from", "to"})

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

func UpdateRequestsTotal(handler string) {
	requestsTotal.WithLabelValues(handler).Inc()
}

func UpdateResponseCode(handler, code string) {
	responseCode.WithLabelValues(handler, code).Inc()
}

func UpdateResponseTime(start time.Time) {
	responseTime.Observe(time.Since(start).Seconds())
}

func UpdateOrdersCreated() {
	ordersCreated.Inc()
}

func UpdateOrderStatusChanged(from, to string) {
	orderStatusChanged.WithLabelValues(from, to).Inc()
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
