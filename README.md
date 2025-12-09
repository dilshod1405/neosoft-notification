📡 Real-Time Notification Service

High-performance WebSocket service built with Go, powered by Redis Streams for guaranteed, scalable, fault-tolerant delivery.

This service enables instant real-time notifications across the entire Neosoft education platform, ensuring messages are delivered reliably even under heavy load.

<div align="center">
⚡ Ultra-Fast • 📬 Redis-Backed Delivery • 🔐 Secure • 🐳 Dockerized • 📊 Production-Ready
</div>
🏗️ Architecture Overview
Notification Flow


Django (creates notification)
        │
        ▼
Redis Streams (event queue — durable)
        │  subscriber
        ▼
Go Notification Service (WebSocket broadcaster)
        │
        ▼
Users (mentors / students via WS)



What each part does:
Layer	Responsibility
Django API	Creates & stores notifications in DB
Redis Streams	Guarantees message durability & delivery order
Go WebSocket Service	Pushes real-time notifications to active users
Clients (Web/Flutter)	Receive instant WS messages
🔥 Why Redis Streams?
✔ Messages are never lost, even if Go server restarts
✔ Horizontal scaling becomes trivial (multiple Go WS servers)
✔ Backpressure handled cleanly
✔ Redis groups prevent duplicate delivery
✔ Perfect for microservice event-driven architectures

Redis Streams = reliable event queue → WebSocket server = pure delivery engine.

```bash
📂 Project Structure
internal/
│
├── stream/
│   └── consumer.go      # Redis Streams → Go WS bridge
│
├── websocket/
│   ├── hub.go           # Manages all connected users
│   ├── client.go        # Writes messages to WebSocket
│   └── handler.go       # /ws endpoint (user_id-based session)
│
├── http/
│   └── router.go        # Routes WS & health endpoints
│
└── models/
    └── notification.go  # Notification schema

cmd/server/main.go       # Application entrypoint
Dockerfile               # Distroless production build
docker-compose.yml       # Deployment config
```

🔌 API Flow (NEW — Redis version)
✔ Django → Redis Streams

Instead of calling Go directly, Django writes to Redis:

```python
r.xadd("notifications_stream", {"data": json.dumps(notification)})
```


Example notification written to Redis:

```json
{
  "id": 26,
  "user_id": 3,
  "type": "promo",
  "title": "Holiday Discount",
  "message": "Hurmatli o'rganuvchilar, sizlar uchun maxsus bayram chegirmalari boshlandi!",
  "metadata": {
    "discount_percent": 30,
    "holiday": "New Year 2026"
  },
  "action_url": "/student/discount",
  "created_at": "2025-12-09T03:06:20.784497Z"
}
```

✔ Go Service (Redis Consumer)

The Go service listens on Redis Streams:

```shell
XREADGROUP GROUP notif_group notif_worker STREAMS notifications_stream >
```


Every event becomes a WebSocket push for that specific user.

🌐 WebSocket Endpoint

Connect:
```bash
ws://localhost:8081/ws?user_id={USER_ID}
```


Example:

```bash
ws://localhost:8081/ws?user_id=3
```

Real-time message example:

```json
{
  "id": 26,
  "title": "Holiday Discount",
  "message": "Hurmatli o'rganuvchilar...",
  "type": "promo",
  "metadata": { "discount_percent": 30 }
}
```

🖥️ WebSocket Client Example (React)

```javascript
const socket = new WebSocket(`ws://localhost:8081/ws?user_id=${userId}`);

socket.onmessage = (event) => {
  const notif = JSON.parse(event.data);
  console.log("Real-time:", notif);
};
```


🐳 Dockerfile (Redis-powered Distroless Build)

```Dockerfile
FROM golang:1.23 AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o notification ./cmd/server

FROM gcr.io/distroless/base-debian11

WORKDIR /app
COPY --from=build /app/notification /app/notification

EXPOSE 8081
ENTRYPOINT ["/app/notification"]
```

🐳 Docker Compose

```YAML
services:
  redis:
    image: redis:7
    restart: always
    ports:
      - "6379:6379"

  notification-service:
    build: .
    restart: always
    depends_on:
      - redis
    ports:
      - "8081:8081"
    networks:
      - neosoft-net

networks:
  neosoft-net:
    driver: bridge
```

🧩 Key Advantages
🟩 Blazing Fast

Go routines + Redis Streams → thousands of WS connections with minimal CPU.

🟩 Reliable Delivery

Messages survive:

Go crashes

Network failures

High load

Guaranteed by Redis Streams.

🟩 Scalable

Add multiple Go instances — Redis handles load balancing.

🟩 Secure

Distroless → zero shell, minimal attack surface.

🟩 Enterprise Architecture

Event-driven, microservice-friendly, horizontally scalable.

👨‍💻 Author

Dilshodjon Normurodov
Real-time Systems • Microservices • Go • Redis • Django • DevOps