# Golim

Rate-limited proxy server in Go. This can be used in two ways

1. **When you are hosting a web server**. This can be used to protect your web server from huge traffic or DDOS attacks.
2. **When you are using someone's web server**. This can be used to make sure you don't cross the free API limit and get billed for more $$.

# Usage

## Pull from DockerHub

The easiest way to use this is via the docker hub image. To start running golim on port 8080:

```sh
docker run --rm -p 8080:80 shubham1172/golim:latest -web-server-addr=http://192.168.0.12:8000
```

## Build from source

Make sure to edit the environment variables in the docker-compose.yml file.

```sh
git clone https://github.com/shubham1172/golim
cd golim
docker-compose up
```

# Configuration

Golim can be configured via the command line arguments or the environment variables. If both are specified, the command line arguments takes precedence.

| Command Line Argument | Environment Variable | Description |
|-|-|-|
| -web-server-addr | GOLIM_WEB_SERVER_ADDR | Address of the web server. |
| -rate-limiter-burst | GOLIM_RATE_LIMITER_BURST | Number of requests that can be made in a given time window. |
| -rate-limiter-window-seconds | GOLIM_RATE_LIMITER_WINDOW_SECONDS | Time window in seconds for which the burst is allowed. |


