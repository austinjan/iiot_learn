# iiot_learn
Learning Grafana &amp; Prometheus

## Target
Use Prometheus and Grafana to build a industrial metrix dashboard

- A app `modbus_exporter` to export modbus data to prometheus
  -  


# Quick Start
## 需求
- Docker
- Docker Compose

1. 下載專案
```bash
git clone https://github.com/kanyuanfeng/iiot_learn.git
```
2. 修改 `/config_volumn` 目錄下的設定檔
   - `prometheus.yml` 這是用於 Prometheus 的設定檔
   - `config.yaml` 這是用於 `modbus_exporter` 的設定檔，可以依據需求增減 data_points 的設定

3. 啟動服務
```bash
docker-compose up -d
```

> 啟動服務後，可以透過 Docker 查看服務是否啟動成功
```bash
docker ps
```
4. 瀏覽器查看 Grafana 服務，並設定 Dashboard。預設網址為 `http://localhost:3000`


 # modbus_poller
 讀取指定 Modbus 伺服器上的數據，並將數據轉換為 Prometheus 的數據格式


### Flags
- `--config`: Path to the configuration file (default: `/config/config.yaml`)
- `--simulate`: Simulate the Modbus server with random data

## Configuration
The configuration file is a YAML file with the following structure:

```yaml
modbus:
  address: "localhost:502" # Modbus server address and port
  slaveID: 1 # Slave ID (Unit ID)

polling_interval_seconds: 3 # Polling interval in seconds

data_points:
  - register: 001 # Modbus register address
    name: "room_temperature" # Label for the data point and prometheus metric name
    type: "gauge" # prometheus metric type
    format: "float32" # Display format (e.g., int, float32, float64)
```

The configuration file (`config.yaml`) contains several key sections:

1. `modbus`: This section defines the Modbus server connection details.
   - `address`: The IP address and port of the Modbus server (e.g., "localhost:502").
   - `slaveID`: The Slave ID (also known as Unit ID) of the Modbus device.

2. `polling_interval_seconds`: This specifies how often the application should poll the Modbus server for data, in seconds.

3. `data_points`: This is an array of data points to be polled from the Modbus server. Each data point has the following properties:
   - `register`: The Modbus register address to read from.
   - `name`: A label for the data point, which will also be used as the Prometheus metric name.
   - `type`: The Prometheus metric type (e.g., "gauge", "histogram").
   - `format`: The data format of the register value (e.g., "int", "float32", "float64").

4. `http_server`: This section contains settings for the HTTP server that exposes the Prometheus metrics.
   - `port`: The port number on which the HTTP server will listen.

This configuration allows you to specify which Modbus registers to read, how to interpret the data, and how to expose it as Prometheus metrics. The application will use this configuration to connect to the Modbus server, read the specified registers at the given interval, and make the data available for Prometheus to scrape.



