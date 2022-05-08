# go-kriging-service

Golang Service for [go-kriging](https://github.com/lvisei/go-kriging) (Golang Multi-Goroutine spatial interpolation algorithm library for geospatial prediction and mapping via ordinary kriging)

## How to get

### Download

You can download from GitHub [releases](https://github.com/lvisei/go-kriging-service/releases).

For example download file:

- windows: `**_windows_x86_64.zip`
- maxOS x86: `**_darwin_x86_64.tar.gz`
- maxOS M1: `**_darwin_arm64.tar.gz`

### Build from source

```
git clone https://github.com/lvisei/go-kriging-service
cd cmd/go-kriging-service && go install
```

## Usage

```bash
go-kriging-service web -c config.toml
```

Options flags:

```
go-kriging-service

Usage: go-kriging-service web -c <configFile>

Options:
  --conf value, -c value  config file(.json,.yaml,.toml)
```
