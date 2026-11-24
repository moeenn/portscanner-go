package scanner

import (
	"log/slog"
	"net"
	"portscanner/internal/semaphore"
	"strconv"
	"sync"
	"time"
)

const (
	checkTimeout time.Duration = time.Second
)

type PortGroup struct {
	Name  string
	Ports []int
}

type PortWithName struct {
	Name string
	Port int
}

func CommonPorts() []PortGroup {
	return []PortGroup{
		{"http/s", []int{80, 443}},
		{"ssh", []int{22, 2222, 22222}},
		{"postgresql", []int{5432}},
		{"mysql", []int{3306, 33060, 33062}},
		{"mongodb", []int{27017, 27018, 27019, 27020}},
		{"redis", []int{6379}},
		{"influx-db", []int{8086}},
		{"elastic-search", []int{9200}},
		{"rabbit-mq", []int{5672, 15672}},
		{"web-app", []int{3000, 5000, 8000, 8080, 9000}},
		{"prometheus", []int{9090}},
		{"loki", []int{3100, 9095}},
		{"tempo", []int{3200, 9096}},
		{"mimir", []int{9009}},
		{"grpc", []int{50051}},
		{"jupyter", []int{8888}},
	}
}

func isPortOpen(host string, port int) bool {
	address := net.JoinHostPort(host, strconv.Itoa(port))
	conn, err := net.DialTimeout("tcp", address, checkTimeout)
	if err != nil {
		return false
	}
	//nolint:errcheck
	defer conn.Close()
	return true
}

func Scan(host string, ports []PortGroup) []PortWithName {
	result := make([]PortWithName, 0, 5)

	sem := semaphore.New(4)
	var wg sync.WaitGroup

	for _, pg := range ports {
		for _, p := range pg.Ports {
			wg.Go(func() {
				sem.Acquire()
				defer sem.Release()

				slog.Info("checking port", "port", p)
				if isPortOpen(host, p) {
					result = append(result, PortWithName{Port: p, Name: pg.Name})
				}
			})
		}
	}

	wg.Wait()
	return result
}
