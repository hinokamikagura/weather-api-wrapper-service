# Weather API Wrapper Service

[![Go Version](https://img.shields.io/badge/Go-1.25.1-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![Gin Framework](https://img.shields.io/badge/Gin-v1.11.0-00ADD8.svg)](https://gin-gonic.com/)

A high-performance REST API wrapper service for WeatherAPI.com built with Go and Gin framework. This service provides weather information with Redis caching, rate limiting, and optimized response handling.

## 🚀 Features

- **Weather Data Fetching**: Retrieve current weather information for any city via WeatherAPI.com
- **Redis Caching**: Intelligent caching mechanism to reduce API calls and improve response times
- **Rate Limiting**: IP-based rate limiting (1 request/minute with burst of 50) to prevent abuse
- **RESTful API**: Clean and intuitive REST API design
- **Error Handling**: Comprehensive error handling with proper HTTP status codes
- **Logging**: Built-in request logging and error tracking
- **High Performance**: Built on Gin framework for optimal performance

## 📋 Prerequisites

Before you begin, ensure you have the following installed:

- **Go** 1.25.1 or higher ([Download](https://golang.org/dl/))
- **Redis** server (local or remote instance)
- **WeatherAPI.com** API key ([Get one here](https://www.weatherapi.com/signup.aspx))

## 🛠️ Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/hinokamikagura/weather-api-wrapper-service.git
   cd weather-api-wrapper-service
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Configure environment variables**
   
   Create a `.env` file in the root directory (or set environment variables):
   ```env
   WEATHER_API_KEY=your_weather_api_key_here
   REDIS_ADDR=localhost:6379
   REDIS_USERNAME=default
   REDIS_PASSWORD=your_redis_password
   REDIS_DB=0
   PORT=8080
   ```

4. **Update configuration**
   
   ⚠️ **Important**: Update the hardcoded credentials in the following files:
   - `handler/handler.go` - Replace hardcoded API key with environment variable
   - `connection/redis.go` - Replace hardcoded Redis credentials with environment variables

5. **Run the application**
   ```bash
   go run main.go
   ```

   The service will start on `http://0.0.0.0:8080`

## 📖 API Documentation

### Base URL
```
http://localhost:8080/api/v1
```

### Endpoints

#### Get Weather by City

Retrieve current weather information for a specific city.

**Endpoint:** `GET /weather`

**Query Parameters:**
| Parameter | Type   | Required | Description           |
|-----------|--------|----------|-----------------------|
| city      | string | Yes      | Name of the city      |

**Example Request:**
```bash
curl "http://localhost:8080/api/v1/weather?city=London"
```

**Example Response (Success):**
```json
{
  "status": "success",
  "data": {
    "city": "London",
    "data": {
      "location": {
        "name": "London",
        "region": "City of London, Greater London",
        "country": "United Kingdom",
        "lat": 51.52,
        "lon": -0.11,
        "tz_id": "Europe/London",
        "localtime_epoch": 1699123456,
        "localtime": "2023-11-04 12:30"
      },
      "current": {
        "temp_c": 15.0,
        "temp_f": 59.0,
        "condition": {
          "text": "Partly cloudy",
          "icon": "//cdn.weatherapi.com/weather/64x64/day/116.png",
          "code": 1003
        },
        "humidity": 65,
        "cloud": 50,
        "feelslike_c": 14.5,
        "feelslike_f": 58.1
      }
    }
  }
}
```

**Example Response (Error):**
```json
{
  "status": "error",
  "message": "City parameter is required"
}
```

**Status Codes:**
- `200 OK` - Request successful
- `400 Bad Request` - Missing or invalid parameters
- `429 Too Many Requests` - Rate limit exceeded
- `500 Internal Server Error` - Server error or external API failure

#### Forecast Endpoint (Placeholder)

**Endpoint:** `GET /forecast`

Currently returns a placeholder response. Implementation pending.

## 🔧 Configuration

### Rate Limiting

The service implements IP-based rate limiting with the following configuration:
- **Limit**: 1 request per minute per IP address
- **Burst**: 50 requests allowed in burst mode
- **Window**: 1 minute

### Caching

Weather data is cached in Redis using the key format: `weather_{city_name}` (lowercase).

⚠️ **Note**: Currently, cache TTL is set to 0 (no expiration). Consider implementing a TTL (e.g., 15-30 minutes) for better cache management.

## 🏗️ Project Structure

```
weather-api-wrapper-service/
├── connection/
│   └── redis.go          # Redis connection management
├── handler/
│   ├── handler.go        # Handler initialization and configuration
│   ├── getCityWeather.go # Weather data fetching logic
│   └── response.go       # Response formatting utilities
├── router/
│   ├── router.go         # Router initialization
│   └── routes.go         # Route definitions and middleware
├── main.go               # Application entry point
├── go.mod                # Go module dependencies
└── README.md             # Project documentation
```

## 🔐 Security Considerations

⚠️ **Important Security Notes:**

1. **Never commit API keys or credentials** to version control
2. **Use environment variables** for all sensitive configuration
3. **Implement proper secret management** (e.g., AWS Secrets Manager, HashiCorp Vault)
4. **Enable HTTPS** in production environments
5. **Review and update rate limiting** based on your use case
6. **Implement authentication/authorization** if exposing publicly

## 🧪 Testing

```bash
# Run tests (when implemented)
go test ./...

# Run with coverage
go test -cover ./...
```

## 🚀 Deployment

### Docker (Recommended)

Create a `Dockerfile`:

```dockerfile
FROM golang:1.25.1-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o weather-service main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/weather-service .
CMD ["./weather-service"]
```

Build and run:
```bash
docker build -t weather-api-wrapper-service .
docker run -p 8080:8080 --env-file .env weather-api-wrapper-service
```

### Production Considerations

- Use a process manager (systemd, supervisor, or PM2)
- Set up proper logging (file rotation, log aggregation)
- Implement health check endpoints
- Configure reverse proxy (Nginx, Caddy)
- Set up monitoring and alerting
- Use connection pooling for Redis
- Implement graceful shutdown

## 📝 Development Roadmap

- [ ] Environment variable configuration
- [ ] Cache TTL implementation
- [ ] Health check endpoint
- [ ] Graceful shutdown
- [ ] Unit and integration tests
- [ ] Docker containerization
- [ ] CI/CD pipeline
- [ ] API documentation (Swagger/OpenAPI)
- [ ] Forecast endpoint implementation
- [ ] Multiple weather provider support
- [ ] Request/response logging middleware
- [ ] Metrics and monitoring

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Gin Web Framework](https://gin-gonic.com/)
- [WeatherAPI.com](https://www.weatherapi.com/)
- [Redis](https://redis.io/)
- [Go Redis Client](https://github.com/redis/go-redis)

## 📧 Contact

For questions or support, please open an issue on GitHub.

---

**Note**: This is a wrapper service. Make sure you comply with WeatherAPI.com's terms of service and rate limits when using this service.
