package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"portscanner/internal/ip"
	"portscanner/internal/scanner"
)

func run() error {
	hostPtr := flag.String("host", "", "IP address of the host to scan")
	flag.Parse()
	host := *hostPtr

	if !ip.IsValid(host) {
		return fmt.Errorf("invalid IP address: %s", host)
	}

	ports := scanner.CommonPorts()

	slog.Info("starting scan", "host", host)
	openPorts := scanner.Scan(host, ports)
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
