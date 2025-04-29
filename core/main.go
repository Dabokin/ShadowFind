package main

import (
	"fmt"
	"time"

	"github.com/Dabokin/ShadowFind/core/scanner"
)

func main() {
	host := "example.com"
	ports := []int{22, 80, 443, 8080, 5432}

	// Использование с настройками по умолчанию
	fmt.Println("🔍 Быстрое сканирование (таймаут 2s):")
	openPorts := scanner.ScanPorts(host, ports)
	printResults(openPorts)

	// Использование с кастомной конфигурацией
	fmt.Println("\n🔍 Тщательное сканирование (таймаут 5s, 10 workers):")
	cfg := scanner.Config{
		Timeout:    5 * time.Second,
		MaxWorkers: 10,
	}
	openPorts = scanner.ScanPortsWithConfig(host, ports, cfg)
	printResults(openPorts)
}

func printResults(ports []int) {
	if len(ports) == 0 {
		fmt.Println("❌ Открытых портов не найдено")
		return
	}

	fmt.Println("✅ Открытые порты:")
	for _, p := range ports {
		fmt.Printf("- %d\n", p)
	}
}
