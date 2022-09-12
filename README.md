# lidl-connect-exporter

> Exporter for billing and usage metrics of LIDL Connect SIM card

## Configuration

| Environment variable       | Description                | Required? | Example                    |
| -------------------------- | -------------------------- | --------- | -------------------------- |
| TZ                         | System timezone            | no        | Europe/Berlin              |
| LOG_LEVEL                  | Logging level              | no        | info                       |
| LIDL_CONNECT_USERNAME      | MSISDN / Rufnummer         | yes       | 015123456789               |
| LIDL_CONNECT_PASSWORD      | Account's password         | yes       | password                   |
| LIDL_CONNECT_HOST          | API Host                   | no        | api.lidl-connect.de        |
| HTTP_PORT                  | Server's http port         | no        | 9100                       |
| METRICS_PATH               | Server's metrics path      | no        | /metrics                   |
