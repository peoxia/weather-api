<div align="center">
    <h1>user-api</h1>
    <p>The API for accessing and managing data about users.</p>
</div>

# Description
User-api is an example of a microservice for accessing and managing data about users.
It uses mocked MySQL client as storage and mocked PubSub client for notifying other apps
about the events that modify data.

# Requirements
* Go >=1.17
* Docker >= 20.10.7

# API
**GET /_healthz** - health check

**GET /api/v1/users** - fetches the list of users by pages. Available query parameters:
* pageId (required)
* country (optional, use UK or Italy values for testing)

**GET /api/v1/users/{user-id}** - fetches data about a user

**POST /api/v1/users** - creates a new user or updates an existing one

**DELETE /api/v1/users/{user-id}** - deletes a user

# Run
```
sudo docker build -t user-api .
sudo docker run -p 8000:8000 user-api
```

The API will be accessible at localhost:8000.
