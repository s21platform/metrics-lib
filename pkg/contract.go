//go:generate mockgen -destination=mock_contract.go -package=${GOPACKAGE} -source=contract.go

package pkg

type MetricInterface interface {
	Increment(name string)
	Gauge(name string, value float64)
	Count(name string, value int64)
	Duration(timestamp int64, name string)
}
