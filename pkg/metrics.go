package pkg

import (
	"log"
	"time"

	"github.com/marpaia/graphite-golang"
)

type Metrics struct {
	g *graphite.Graphite
}

func NewMetrics(host string, port int) (*Metrics, error) {
	g, err := graphite.NewGraphite(host, port)
	if err != nil {
		return nil, err
	}
	return &Metrics{g: g}, nil
}

func (m *Metrics) Test() {
	log.Println("Test. Metric works")
}

func (m *Metrics) Increment(name string) {
	_ = m.g.SendMetric(graphite.Metric{
		Name:      name,
		Value:     "1",
		Timestamp: time.Now().Unix(),
	})
}

func (m *Metrics) Disconnect() {
	_ = m.g.Disconnect()
}
