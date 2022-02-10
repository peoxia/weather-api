<div align="center">
    <h1>user-api</h1>
    <p>The API for accessing and managing data about users.</p>
</div>

# Description
User-api is an example of a microservice for accessing and managing data about users.
It uses mocked MySQL client as storage and mocked PubSub client for notifying other apps
about the events that modify data.

# API
<p>GET /_healthz</p> - health check

<p>GET /api/v1/users</p> - fetches the list of users by pages. Available query parameters:
pageId (required)
country (optional, use UK or Italy values for testing)

<p>GET /api/v1/users/{user-id}</p> - fetches data about a user

<p>POST /api/v1/users</p> - creates a new user or updates an existing one

<p>DELETE /api/v1/users/{user-id}</p> - deletes a user

# Run
sudo docker build -t user-api .
sudo docker run -p 8000:8000 user-api

The API will be accessible at localhost:8000.