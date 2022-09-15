# lidl-connect-exporter

> Exporter for billing and usage metrics of LIDL Connect SIM card

## Getting started

There are pre-built docker images, so you can run it

either with docker:

```
docker run --env LIDL_CONNECT_USERNAME=015123456789 --env LIDL_CONNECT_PASSWORD=password ghcr.io/avakarev/lidl-connect-exporter:v1.1.0
```

or with docker-compose:

```
version: "3.7"

services:
  lidl_connect_exporter:
    image: ghcr.io/avakarev/lidl-connect-exporter:v1.1.0
    container_name: lidl-connect-exporter
    expose:
      - 9100
    environment:
      LIDL_CONNECT_USERNAME: "015123456789"
      LIDL_CONNECT_PASSWORD: "password"
```

## Configuration

lidl-connect-exporter is configured via environment variables.

| Environment variable       | Description                | Required? | Example                    |
| -------------------------- | -------------------------- | --------- | -------------------------- |
| TZ                         | System timezone            | no        | Europe/Berlin              |
| LOG_LEVEL                  | Logging level              | no        | info                       |
| LIDL_CONNECT_USERNAME      | MSISDN / Rufnummer         | yes       | 015123456789               |
| LIDL_CONNECT_PASSWORD      | Account's password         | yes       | password                   |
| LIDL_CONNECT_HOST          | API Host                   | no        | api.lidl-connect.de        |
| HTTP_PORT                  | Server's http port         | no        | 9100                       |
| METRICS_PATH               | Server's metrics path      | no        | /metrics                   |

## Prometheus Configuration

Example config:

```
scrape_configs:
  - job_name: lidl_connect
    static_configs:
      - targets: ['lidl_connect_exporter:9100']
```

## Metrics

### Balance

| Name                              | Type             | Labels         | Description                 |
| --------------------------------- | ---------------- | -------------- | --------------------------- |
| lidl_connect_balance              | gauge            | []             | The state of the balance    |

Example:

```
# HELP lidl_connect_balance The state of the balance
# TYPE lidl_connect_balance gauge
lidl_connect_balance 7.01
```

### Booked Tariff Fee

| Name                              | Type             | Labels         | Description                 |
| --------------------------------- | ---------------- | -------------- | --------------------------- |
| booked_tariff_fee                 | gauge            | [name]         | Booked tariff fee           |

Example:

```
# HELP lidl_connect_booked_tariff_fee Booked tariff fee
# TYPE lidl_connect_booked_tariff_fee gauge
lidl_connect_booked_tariff_fee{name="Data S"} 2.99
```

### Consumptions

| Name                                    | Type             | Labels         | Description                            |
| --------------------------------------- | ---------------- | -------------- | -------------------------------------- |
| lidl_connect_consumption_consumed       | gauge            | [unit, type]   | Consumption volume consumed            |
| lidl_connect_consumption_left           | gauge            | [unit, type]   | Consumption volume left                |
| lidl_connect_consumption_max            | gauge            | [unit, type]   | Consumption volume max                 |
| lidl_connect_consumption_expires_in_sec | gauge            | [unit, type]   | Consumption expiration time in seconds |

Example:

```
# HELP lidl_connect_consumption_consumed Consumption consumed
# TYPE lidl_connect_consumption_consumed gauge
lidl_connect_consumption_consumed{type="DATA",unit="GB"} 1.1
# HELP lidl_connect_consumption_expires_in_sec Consumption expires in seconds
# TYPE lidl_connect_consumption_expires_in_sec gauge
lidl_connect_consumption_expires_in_sec{type="DATA",unit="GB"} 1.247851107573e+06
# HELP lidl_connect_consumption_left Consumption left
# TYPE lidl_connect_consumption_left gauge
lidl_connect_consumption_left{type="DATA",unit="GB"} 6.63
# HELP lidl_connect_consumption_max Consumption max
# TYPE lidl_connect_consumption_max gauge
lidl_connect_consumption_max{type="DATA",unit="GB"} 7.73
```


## License

`lidl-connect-exporter` is licensed under MIT license. (see [LICENSE](./LICENSE))
