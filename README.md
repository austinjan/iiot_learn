# iiot_learn
Learning Grafana &amp; Prometheus

## Target
Use Prometheus and Grafana to build a industrial metrix dashboard

- A app `modbus_exporter` to export modbus data to prometheus
  -  


# Quick Start
## Prerequisites
- Docker
- Docker Compose

## Run
```bash
docker-compose up -d
```

## Grafana
Browse to http://localhost:3000
- Add Prometheus as data source
  - Type: Prometheus
  - URL: http://prometheus:9090
- Create new dashboard
- Import dashboard
 



