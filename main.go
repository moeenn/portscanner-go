package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"portscanner/internal/host"
	"portscanner/internal/scanner"
	"time"
)

func run() error {
	hostPtr := flag.String("host", "", "URL or IP address of the host to scan")
	maxParallerPtr := flag.Int("parallel", 4, "Maximum number of parallel workers")
	flag.Parse()

	host, err := host.Validate(hostPtr)
	if err != nil {
		return fmt.Errorf("invalid host: %s", err)
	}

	ports := scanner.CommonPorts()

	slog.Info("starting scan", "host", host)
	openPorts := scanner.Scan(scanner.ScanArgs{
		Host:        host,
		Ports:       ports,
		Timeout:     time.Second,
		MaxParallel: *maxParallerPtr,
	})
	slog.Info("scan complete")

	fmt.Println()

	for _, p := range openPorts {
		fmt.Printf("%-15s: %d\n", p.Name, p.Port)
	}

	return nil
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		os.Exit(1)
	}
}
