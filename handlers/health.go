package handlers

import (
	"net/http"
	"sync"

	"github.com/sl1pm4t/sems-healthz/sems"
)

const (
	healthyStatus   = http.StatusOK
	unhealthyStatus = http.StatusServiceUnavailable
)

var (
	readinessStatus = healthyStatus
	mu              sync.RWMutex
)

func HealthzStatus() int {
	// fmt.Println("HealthzStatus()")
	mu.RLock()
	defer mu.RUnlock()

	_, err := sems.GetActiveCallCount()
	if err != nil {
		return unhealthyStatus
	}
	return healthyStatus
}

func ReadinessStatus() int {
	mu.RLock()
	defer mu.RUnlock()
	// TODO(sl1pm4t) - implement overload condition checking
	return readinessStatus
}

// HealthzHandler responds to health check requests.
func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(HealthzStatus())
}

// ReadinessHandler responds to readiness check requests.
func ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(ReadinessStatus())
}
