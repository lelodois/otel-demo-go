# otel-demo

Sample golang application showcasing how to instrumenting distributed tracing with OpenTelemetry and Jaeger


## Prerequisites

- go
- docker
- docker-compose

## Start demo application

```sh
$ make demo
```

## Using api

````
curl --location 'http://localhost:3000/leads' --header 'Content-Type: application/json' --data-raw 
'{
    "name": "Lead Tasty Plastic Pizza",
    "email": "lead@Practical Frozen Table.com",
    "phone_number": "11 - 323-490-9912",
    "address": "46438 Schroeder Island"
}'

curl --location 'http://localhost:3000/leads/:id'
````

## Stop

```sh
$ make stop
```

## Ports

 - Jaeger: 16686 (https://www.jaegertracing.io/docs/1.56/getting-started)
 - Database: 5432
 - App: 3000


## Doc

Jaeger:
![jaeger ui](docs%2Fjaeger_ui.png)
---

Docker containers
![docker](docs%2Fdocker.png)
---

Postgres
![database](docs%2Fdatabase.png)
