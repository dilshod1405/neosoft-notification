# 📡 Real-Time Notification Service  
**High-performance WebSocket service built with Go**, designed for instant, reliable delivery of notifications across the Neosoft education platform.

This service powers **real-time communication** between backend (Django) and users (mentors & students).  
Perfect for distributed systems, microservice environments, and modern SaaS dashboards.

---

<div align="center">
  
### ⚡ Ultra-Fast • 🧽 Clean Architecture • 🔐 Secure • 🐳 Dockerized • 📊 Production-Ready

</div>

---

## 🏗️ Architecture Overview

**Django** handles business logic + Notification storage.  
**Go Notification Service** handles WebSocket delivery instantly.  
Clients receive messages in **real-time**, even under high load.




                   ┌──────────────────────────┐
                   │        Django API        │
                   │  Notification created    │
                   └─────────────┬────────────┘
                                 │ POST /publish
                                 ▼
                   ┌──────────────────────────┐
                   │   Go Notification Svc    │
                   │  (WebSocket Broadcaster) │
                   └─────────────┬────────────┘
                                 │ WS Push
                                 ▼
                   ┌──────────────────────────┐
                   │   Web / Mobile Clients   │
                   │ React / Flutter / etc.   │
                   └──────────────────────────┘





---

## 🧠 Core Concepts

### ✔ **Real-Time Delivery**
Each connected user holds a dedicated WebSocket channel.  
Notifications are pushed instantly using lightweight Go routines.

### ✔ **User-Specific Channels**
Connections are mapped as:




user_id -> active WebSocket clients



A user with multiple tabs/devices receives notifications everywhere instantly.

### ✔ **Super-Light Runtime**
Final production image is built using **Distroless**:
- No shell  
- No package manager  
- Minimal attack surface  
- Ultra-fast & secure

### ✔ **Offline-Safe**
Even if a user is offline:
- Django stores the notification  
- User reads it later via REST API  
- WebSocket is only for real-time delivery

---

## 🛠️ Technology Stack

| Layer | Technology |
|-------|------------|
| Language | **Go 1.23** |
| Framework | **Gin** (HTTP router) |
| WebSockets | **Gorilla/WebSocket** |
| Runtime | **Distroless (Debian 11)** |
| Deployment | Docker & Docker Compose |
| Architecture | Microservice, Event-driven push |
| Backend paired with | Django REST Framework |

---

## 📂 Project Structure

internal/
│
├── http/
│ ├── handler.go # POST /publish (Django → Go)
│ └── router.go # Service routing
│
├── websocket/
│ ├── hub.go # Manages all users & connections
│ ├── client.go # Single WS client logic
│ └── upgrader.go # HTTP → WebSocket upgrader
│
└── models/
└── notification.go # Notification payload model

cmd/server/main.go # Application entrypoint
Dockerfile # Distroless production image
docker-compose.yml # Easy deployment



---

## 🔌 API Endpoints

### 1️⃣ Publish Notification  
**POST /publish**

Django backend uses this to broadcast a new notification.

```json
{
  "id": 14,
  "user_id": 52,
  "type": "promo",
  "title": "Holiday Discount",
  "message": "Chegirmalar boshlandi!",
  "metadata": { "discount": 30 },
  "action_url": "/courses",
  "created_at": "2025-01-15T10:00:00Z"
}



2️⃣ WebSocket Endpoint

GET /ws?user_id={id}

Example:

ws://localhost:8081/ws?user_id=52


Real-time messages arrive as JSON:

{
  "title": "New Message",
  "message": "Your lesson has been updated",
  "type": "info"
}

🖥️ Client Example (React/Next.js)
const socket = new WebSocket(`ws://localhost:8081/ws?user_id=${userId}`);

socket.onmessage = (event) => {
  const notif = JSON.parse(event.data);
  console.log("Real-time:", notif);
};

🐳 Docker (Production Ready)
Dockerfile (Distroless)
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

Docker Compose
version: "3.9"

services:
  notification-service:
    build: .
    ports:
      - "8081:8081"
    restart: always
    networks:
      - neosoft-net

networks:
  neosoft-net:
    driver: bridge

🧩 Key Advantages
🟩 Fast

Handles thousands of concurrent connections with minimal memory.

🟩 Reliable

Even if client disconnects, Django stores the notification.

🟩 Clean

Stateless architecture makes scaling trivial.

🟩 Secure

Distroless → Zero shell, zero bloat.

🟩 Professional

Ideal for microservice-based platforms.

🧑‍💻 Author

Dilshodjon Normurodov
Real-time systems • Microservices • Go • Django • DevOps
