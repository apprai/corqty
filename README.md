# Corqty

Corqty is the new way to define infrastructure

## Quick Start

Check out your development [configuration](config/development.yaml)

```sh
docker-compose up -d

docker run -d -p 5432:5432 -e POSTGRES_USER=loco -e POSTGRES_DB=[app name]_development -e POSTGRES_PASSWORD="loco" postgres:15.3-alpine

cargo run -- start
```
