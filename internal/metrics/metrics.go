package metrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/timickb/transport-sound/internal/interfaces"
	"net/http"
	"strconv"
)

type Metrics struct {
	port   int
	logger interfaces.Logger

	httpReqTotal  *prometheus.CounterVec
	newUsersTotal prometheus.Counter
}

func New(logger interfaces.Logger, port int) *Metrics {
	return &Metrics{
		port:   port,
		logger: logger,
		httpReqTotal: promauto.NewCounterVec(prometheus.CounterOpts{
			Namespace: "soundp_main",
			Subsystem: "http",
			Name:      "requests_total",
		}, []string{"status"}),
		newUsersTotal: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: "soundp_main",
			Subsystem: "http",
			Name:      "new_accounts_total",
		}),
	}
}

func (m *Metrics) handler() http.HandlerFunc {
	h := promhttp.Handler()

	return func(w http.ResponseWriter, r *http.Request) {
		m.logger.Info("Metrics fetching")
		h.ServeHTTP(w, r)
	}
}

func (m *Metrics) Listen() error {
	mux := http.NewServeMux()
	mux.Handle("/metrics", m.handler())

	return http.ListenAndServe(fmt.Sprintf(":%d", m.port), mux)
}

func (m *Metrics) AddHttpRequest(status int) {
	m.httpReqTotal.WithLabelValues(strconv.Itoa(status)).Inc()
}

func (m *Metrics) AddNewUser() {
	m.newUsersTotal.Inc()
}
