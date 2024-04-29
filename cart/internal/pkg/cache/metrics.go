package cache

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"route256.ozon.ru/project/cart/internal/config"
)

var cacheMissTotal = promauto.NewCounterVec(prometheus.CounterOpts{
	Subsystem: strings.Replace(config.AppName, "-", "_", -1),
	Name:      "cache_miss_total",
}, []string{"key"})

var cacheHitTotal = promauto.NewCounterVec(prometheus.CounterOpts{
	Subsystem: strings.Replace(config.AppName, "-", "_", -1),
	Name:      "cache_hit_total",
}, []string{"key"})

func UpdateCacheMissTotal(key string) {
	cacheMissTotal.WithLabelValues(key).Inc()
}

func UpdateCacheHitTotal(key string) {
	cacheHitTotal.WithLabelValues(key).Inc()
}
