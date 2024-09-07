package pkg

import (
	"github.com/marpaia/graphite-golang"
	"log"
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

func (m *Metrics) Close() {
	_ = m.g.Disconnect()
}
