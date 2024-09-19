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


 



