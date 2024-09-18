package main

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	config, err := loadConfig("config.yaml")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	if config.Modbus.Address != "192.168.1.100:502" {
		t.Fatalf("Expected address to be 192.168.1.100:502, got %s", config.Modbus.Address)
	}

	if config.PollingIntervalSeconds != 10 {
		t.Fatalf("Expected polling interval to be 10, got %d", config.PollingIntervalSeconds)
	}

	if len(config.DataPoints) != 2 {
		t.Fatalf("Expected 2 data points, got %d", len(config.DataPoints))
	}

	if config.DataPoints[0].Register != 100 {
		t.Fatalf("Expected register to be 100, got %d", config.DataPoints[0].Register)
	}

}
