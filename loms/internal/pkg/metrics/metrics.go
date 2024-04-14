package metrics

import (
	_ "net/http/pprof"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"route256.ozon.ru/project/loms/internal/config"
)

var requestsCounter = promauto.NewCounterVec(prometheus.CounterOpts{
	Subsystem: strings.Replace(config.AppName, "-", "_", -1),
	Name:      "requests_total",
	Help:      "Total amount of requests made to handler. Example: rate(post_requests_total[1m])",
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

func UpdateRequestsCounter(handler string) {
	requestsCounter.WithLabelValues(handler).Inc()
}

func UpdateResponseCode(handler string, code string) {
	responseCode.WithLabelValues(handler, code).Inc()
}

func UpdateResponseTime(start time.Time) {
	responseTime.Observe(time.Since(start).Seconds())
}

func UpdateOrdersCreated() {
	ordersCreated.Inc()
}

func UpdateOrderStatusChanged(from string, to string) {
	orderStatusChanged.WithLabelValues(from, to).Inc()
}
