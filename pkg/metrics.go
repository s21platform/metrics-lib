package pkg

import (
	"context"
	"fmt"
	"time"

	"github.com/marpaia/graphite-golang"
)

type Metrics struct {
	g      *graphite.Graphite
	prefix string
}

func NewMetrics(host string, port int, service string, env string) (*Metrics, error) {
	g, err := graphite.NewGraphite(host, port)
	if err != nil {
		return nil, err
	}
	return &Metrics{
		g:      g,
		prefix: fmt.Sprintf("%s.%s.", service, env),
	}, nil
}

func (m *Metrics) Increment(name string) {
	_ = m.g.SendMetric(graphite.Metric{
		Name:      m.prefix + name + ".count",
		Value:     "1",
		Timestamp: time.Now().Unix(),
	})
}

func (m *Metrics) Duration(timestamp int64, name string) {
	_ = m.g.SendMetric(graphite.Metric{
		Name:      name + ".duration",
		Value:     fmt.Sprintf("%d", timestamp),
		Timestamp: time.Now().Unix(),
	})
}

func FromContext(ctx context.Context) *Metrics {
	value := ctx.Value("metrics")
	if value == nil {
		// Обрабатываем ситуацию, когда значение отсутствует в контексте
		return nil
	}

	metrics, ok := value.(*Metrics)
	if !ok {
		// Обрабатываем ситуацию, когда значение есть, но неверного типа
		return nil
	}

	return metrics
}

func (m *Metrics) Disconnect() {
	_ = m.g.Disconnect()
}
