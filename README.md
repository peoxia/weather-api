<div align="center">
    <h1>weather-api</h1>
    <p>The API for requesting weather information.</p>
</div>

# Requirements
* Go >=1.17
* Docker >= 20.10.7

# API
**GET /_healthz** - health check

**GET /v1/weather** - returns current weather by the location. Available query parameters:
* city (optional)


# Run
**Make**
```
make server
```

**Docker**
```
sudo docker build -t weather-api .
sudo docker run -p 8000:8000 weather-api
```

The API will be accessible at localhost:8000.
