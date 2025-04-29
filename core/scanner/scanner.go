package scanner

import (
	"fmt"
	"net"
	"sync"
	"time"
)

// Config содержит настройки сканирования
type Config struct {
	Timeout    time.Duration // Таймаут для каждого соединения
	MaxWorkers int           // Максимальное количество одновременных сканирований
}

// formatTarget правильно форматирует host:port для IPv4/IPv6
func formatTarget(host string, port int) string {
	if ip := net.ParseIP(host); ip != nil && ip.To4() == nil {
		return fmt.Sprintf("[%s]:%d", host, port) // IPv6
	}
	return fmt.Sprintf("%s:%d", host, port) // IPv4 или домен
}

// ScanPort сканирует один порт с учетом конфигурации
func ScanPort(host string, port int, wg *sync.WaitGroup, results chan<- int, timeout time.Duration) {
	defer wg.Done()

	conn, err := net.DialTimeout("tcp", formatTarget(host, port), timeout)
	if err == nil {
		conn.Close()
		results <- port
	}
}

// ScanPortsWithConfig сканирует порты с использованием конфигурации
func ScanPortsWithConfig(host string, ports []int, cfg Config) []int {
	var openPorts []int
	var wg sync.WaitGroup
	results := make(chan int, len(ports))

	// Ограничиваем количество одновременных workers
	semaphore := make(chan struct{}, cfg.MaxWorkers)
	if cfg.MaxWorkers <= 0 {
		semaphore = make(chan struct{}, 10) // Значение по умолчанию
	}

	for _, port := range ports {
		wg.Add(1)
		semaphore <- struct{}{} // Занимаем слот

		go func(p int) {
			defer func() { <-semaphore }()
			ScanPort(host, p, &wg, results, cfg.Timeout)
		}(port)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for port := range results {
		openPorts = append(openPorts, port)
	}

	return openPorts
}

// ScanPorts (удобная обертка с настройками по умолчанию)
func ScanPorts(host string, ports []int) []int {
	return ScanPortsWithConfig(host, ports, Config{
		Timeout:    2 * time.Second,
		MaxWorkers: 100,
	})
}
