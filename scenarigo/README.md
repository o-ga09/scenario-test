# Scenarigo 使い方

## Run on local

## Run on Docker

```bash
docker build -t go-dev-env ./docker/
docker run -it --rm -v $(pwd):/go/src/app -p 8080:8080 go-dev-env
```

## How to Build

```bash
scenarigo plugin build
```