package metrics

import (
	_ "net/http/pprof"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"route256.ozon.ru/project/cart/internal/config"
)

var requestsCounter = promauto.NewCounterVec(prometheus.CounterOpts{
	Subsystem: strings.Replace(config.AppName, "-", "_", -1),
	Name:      "requests_total",
	Help:      "Total amount of requests made to handler. Example: rate(post_requests_total[1m])",
}, []string{"method", "handler"})

var responseCode = promauto.NewCounterVec(prometheus.CounterOpts{
	Subsystem: strings.Replace(config.AppName, "-", "_", -1),
	Name:      "response_code",
	Help:      "Response code of handlers. Example: rate(response_code[1m])",
}, []string{"method", "handler", "code"})

var responseTime = promauto.NewHistogram(prometheus.HistogramOpts{
	Subsystem: strings.Replace(config.AppName, "-", "_", -1),
	Name:      "response_time",
	Buckets:   prometheus.DefBuckets,
	Help:      "Response time of handlers. Example: rate(response_time[1m])",
})

func UpdateRequestsCounter(method, handler string) {
	requestsCounter.WithLabelValues(method, handler).Inc()
}

func UpdateResponseCode(method, handler, code string) {
	responseCode.WithLabelValues(method, handler, code).Inc()
}

func UpdateResponseTime(start time.Time) {
	responseTime.Observe(time.Since(start).Seconds())
}
