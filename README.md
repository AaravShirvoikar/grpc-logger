# grpc-logger

A simple logging service using gRPC. This service logs events to a MongoDB database and provides a basic client endpoint for testing.

## Features
- Accepts events with a name and associated data.
- Client communicates with gRPC server to log the data.
- Stores logs in a MongoDB collection.
- Dockerized for easy deployment.

---

## Getting Started

### Prerequisites
- [Docker](https://www.docker.com/) (docker-compose)

---

### Running the Project

1. **Clone the repository**
   ```bash
   git clone https://github.com/AaravShirvoiakr/grpc-logger.git
   cd grpc-logger
   ```

2. **Build and run the service using Docker**
   ```bash
   make up_build
   ```

---

### API Usage

Send a `POST` request to log an event.

- **Endpoint:** `http://localhost:8080/log`
- **Payload:**
  ```json
  {
    "name": "test_name",
    "data": "test_data"
  }
  ```
