# subscribe-middleware-go

subscribe-middleware-go is a project that implements a proxying (and more) "subscription" method for xray, v2ray systems, and panels based on them (x-ui, 3x-ui). This is a personal project created to fulfill specific needs.

## Features

- Proxying "subscription" method for xray and v2ray systems
- Compatibility with x-ui, 3x-ui, and other panels based on xray and v2ray
- Customizable and extensible functionality

## Getting Started

### Prerequisites

- Go programming language (version 1.16 or later)
- Go framework GIN
- Git
- Docker(optional)
- Docker compose(optional)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/niazlv/subscribe-middleware-go.git
```
2. Navigate to the project directory:
```bash
cd subscribe-middleware-go
```
3. Build the project and run:
```bash
go mod download
go build -o main ./cmd/subscribe-middleware-go/main.go
./main
```
or using docker:
```bash
docker build . -t subscribe_middleware_backend
docker run -d -p 56792:56792 -v $(pwd)/storage:/app/storage --name subscribe_middleware_backend_container subscribe_middleware_backend
```
or using docker compose:
```bash
docker-compose up -d --build
```

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgments

- [xray](https://github.com/xtls/xray-core)
- [v2ray](https://github.com/v2fly/v2ray-core)
- [x-ui](https://github.com/vaxilu/x-ui)
- [3x-ui](https://github.com/MHSanaei/3x-ui)
