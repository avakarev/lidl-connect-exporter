# lidl-connect-exporter

> Exporter for billing and usage metrics of LIDL Connect SIM card

## Getting started

There are pre-built docker images, so you can run it

either with docker:

```
docker run --env LIDL_CONNECT_USERNAME=015123456789 --env LIDL_CONNECT_PASSWORD=password ghcr.io/avakarev/lidl-connect-exporter:v1.0.0
```

or with docker-compose:

```
version: "3.7"

services:
  lidl_connect_exporter:
    image: ghcr.io/avakarev/lidl-connect-exporter:v1.0.0
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

## License

`lidl-connect-exporter` is licensed under MIT license. (see [LICENSE](./LICENSE))
