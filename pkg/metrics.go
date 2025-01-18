package pkg

import (
	"context"
	"fmt"

	"github.com/alexcesaro/statsd"
)

type Metrics struct {
	s      *statsd.Client
	prefix string
}

func NewMetrics(host string, port int, service string, env string) (*Metrics, error) {
	s, err := statsd.New(statsd.Address(fmt.Sprintf("%s:%d", host, port)))
	if err != nil {
		return nil, err
	}
	return &Metrics{
		s:      s,
		prefix: fmt.Sprintf("%s.%s.", service, env),
	}, nil
}

func (m *Metrics) Increment(name string) {
	m.s.Increment(m.prefix + name)
}

func (m *Metrics) Gauge(name string, value float64) {
	m.s.Gauge(m.prefix+name, value)
}

func (m *Metrics) Count(name string, value int64) {
	m.s.Count(m.prefix+name, value)
}

func (m *Metrics) Duration(timestamp int64, name string) {
	m.s.Timing(m.prefix+name+".duration", timestamp)
}

func FromContext(ctx context.Context, name interface{}) MetricInterface {
	value := ctx.Value(name)
	if value == nil {
		// Обрабатываем ситуацию, когда значение отсутствует в контексте
		return nil
	}

	metrics, ok := value.(MetricInterface)
	if !ok {
		// Обрабатываем ситуацию, когда значение есть, но неверного типа
		return nil
	}

	return metrics
}

func (m *Metrics) Disconnect() {
	m.s.Close()
}
