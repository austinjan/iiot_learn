package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/goburrow/modbus"
	"gopkg.in/yaml.v2"
)

// Config represents the structure of config.yaml
type Config struct {
	Modbus struct {
		Address string `yaml:"address"`
		SlaveID int    `yaml:"slaveID"`
	} `yaml:"modbus"`
	PollingIntervalSeconds int `yaml:"polling_interval_seconds"`
	DataPoints             []struct {
		Register int    `yaml:"register"`
		Name     string `yaml:"name"`
		Type     string `yaml:"type"`
		Format   string `yaml:"format"`
	} `yaml:"data_points"`
	HttpServer struct {
		Port int `yaml:"port"`
	} `yaml:"http_server"`
}

// DataPoint holds the polled data
type DataPoint struct {
	Register int
	Name     string
	Value    interface{}
}

// loadConfig loads the configuration from a YAML file
func loadConfig(filename string) (*Config, error) {
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading YAML file: %w", err)
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling YAML: %w", err)
	}

	return &config, nil
}

// connectToModbus connects to the Modbus server
func connectToModbus(config *Config) (modbus.Client, func(), error) {
	handler := modbus.NewTCPClientHandler(config.Modbus.Address)
	handler.SlaveId = byte(config.Modbus.SlaveID)
	client := modbus.NewClient(handler)

	cleanup := func() {
		handler.Close()
	}

	return client, cleanup, nil
}

func main() {

	// flags
	// simulate flag
	simulate := flag.Bool("simulate", false, "simulate the modbus server")
	configPath := flag.String("config", "/config/config.yaml", "path to the config file")
	flag.Parse()

	fmt.Println("Modbus Poller")
	config, err := loadConfig(*configPath)
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}
	client, cleanup, err := connectToModbus(config)
	if err != nil {
		fmt.Println("Error connecting to Modbus:", err)
		return
	}
	defer cleanup()

	DefineMetrics(config)

	var wg sync.WaitGroup
	// if simulate flag is set, start the modbus server
	if *simulate {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Simulating Modbus server")
			// write random values to config.data_points (0-40) every interval
			ticker := time.NewTicker(time.Duration(config.PollingIntervalSeconds) * time.Second)
			defer ticker.Stop()
			for {
				<-ticker.C

				for _, dataPoint := range config.DataPoints {
					value := rand.Int31n(40)
					fmt.Printf("Writing random value to data point %s: %v\n", dataPoint.Name, value)
					// write value to data point
					UpdateInt32Metrics(dataPoint.Name, value)
				}
			}
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		ExposeMetrics(config)
	}()

	wg.Add(1)
	// poll data points
	if !*simulate {
		go func() {
			defer wg.Done()
			for {
				for _, dataPoint := range config.DataPoints {
					fmt.Printf("Polling data (%d) point %s \n", dataPoint.Register, dataPoint.Name)
					value, err := client.ReadHoldingRegisters(uint16(dataPoint.Register), 1)
					if err != nil {
						fmt.Printf("Error reading data point %s: %v\n", dataPoint.Name, err)
						continue
					}

					// convert value to int32 (big endian 16 bit)
					int32value := int32(value[0])<<8 | int32(value[1])
					fmt.Printf("Data point %s: %v\n", dataPoint.Name, int32value)
					// update metrics
					UpdateInt32Metrics(dataPoint.Name, int32value)
				}

				// sleep for the polling interval
				time.Sleep(time.Duration(config.PollingIntervalSeconds) * time.Second)
			}
		}()
	}

	wg.Wait()

}
