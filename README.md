# 🚀 Order Management Service

## 📌 Project Overview

This module is a part of the **Martix** application responsible for managing customer orders. It integrates with other
services (e.g., Product, User, Auth) and ensures scalability, maintainability, and security.

This project is a **high-performance, microservices-based order processing system** designed to handle **payments, order
management, and fulfillment** with low latency and scalability.

Built with **GoLang**, it leverages high-performance tools and best practices for **fault tolerance, security, and
scalability**.

## 🎯 Key Features

✅ **Microservices Architecture**: Each component (Payments, Orders, Fulfillment) runs independently for scalability.  
✅ **High-Performance API**: Uses **Gin** for optimized request handling.  
✅ **Structured Logging**: Zap ensures efficient logging with minimal overhead.  
✅ **Event-Driven Order Processing**: Ensures **asynchronous** and **real-time updates** using messaging queues.  
✅ **Security First**: OAuth2/JWT authentication with rate-limiting for protection.  
✅ **Resilient & Fault-Tolerant**: Implements retries, circuit breakers, and distributed transactions.  
✅ **Database Consistency**: Uses PostgreSQL/CockroachDB for ACID compliance and **strong consistency**.

---

## 🛠 Communication & Architecture

The system uses a **hybrid approach** combining **gRPC, messaging queues, and REST APIs** for efficient inter-service
communication.

### **1️⃣ API Gateway & External Communication**

- **REST APIs**: Public-facing APIs for customer interactions (e.g., placing orders, checking payment status).
- **GraphQL** *(Optional)*: Provides a single endpoint for querying multiple services efficiently.

### **2️⃣ Internal Microservice Communication**

- **gRPC**: Used for high-speed, synchronous communication between critical services (e.g., Payment ↔ Order).
- **Message Queue (Kafka/RabbitMQ)**: Asynchronous processing for **event-driven workflows**:
    - **Order Events**: Order creation, status updates, inventory updates.
    - **Payment Events**: Payment success/failure, refunds.
    - **Notification Events**: Triggering emails, SMS updates.

### **3️⃣ Data Management & Consistency**

- **Event Sourcing**: Ensures order state consistency in a **distributed matrix system**.
- **CQRS (Command Query Responsibility Segregation)**:
    - **Write Operations** → Handled by services with event logs.
    - **Read Operations** → Optimized using materialized views.
- **Database**: PostgreSQL/CockroachDB (Ensures consistency, high availability).

---

## 🔥 System Workflow

1️⃣ **Customer places an order** → Order Service validates request.  
2️⃣ **Payment Service processes the transaction** → Communicates via **gRPC**.  
3️⃣ **Order Service updates order status** → Publishes event to **Kafka/RabbitMQ**.  
4️⃣ **Fulfillment Service processes shipping** → Listens to **order events** asynchronously.  
5️⃣ **Notification Service sends order confirmation** via **event-driven messaging**.

---

## 🔒 Security & Protection

- **Authentication & Authorization**: OAuth2/JWT for user authentication.
- **Rate Limiting & DDoS Protection**: API Gateway limits abusive requests.
- **Retry Mechanism & Circuit Breakers**: Ensures system stability under failures.
- **Database Transactions**: ACID compliance to avoid race conditions.

---

## 🚀 Scalability & Performance

- **Horizontal Scaling**: Microservices auto-scale independently.
- **Load Balancing**: Uses **NGINX** or **Envoy Proxy** for traffic management.
- **Caching**: **Redis** used for frequently accessed data (e.g., order history).

---

## 🎯 Why This Architecture?

1️⃣ **Low Latency** - gRPC ensures fast, binary communication.  
2️⃣ **High Availability** - Event-driven messaging reduces bottlenecks.  
3️⃣ **Scalability** - Stateless services scale horizontally.  
4️⃣ **Data Integrity** - Event Sourcing + CQRS maintains order consistency.

---

## 📌 Popular tools used

A list of high-performance tools used in GoLang for efficient logging, routing, server handling, and dependency
injection.

| Tool           | Company | Purpose                                     | Repository                                                                                                                                                                                                     |
|----------------|---------|---------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| **gin**        |         | Fast, easy to use, middleware support       | [gin](github.com/gin-gonic/gin)                                                                                                                                                                                |
| **Zap**        | Uber    | High-performance logging                    | [zap](https://github.com/uber-go/zap)                                                                                                                                                                          |
| **FX**         | Uber    | Dependency injection                        | [fx](https://github.com/uber-go/fx)                                                                                                                                                                            |
| **pgx**        |         | Fastest PostgreSQL driver in Go.            |                                                                                                                                                                                                                |
| **Watermill**  |         | Kafka/NATS support.                         | [watermill](https://watermill.io/docs/getting-started/)                                                                                                                                                        |
| **golang-jwt** |         | OAuth2/JWT Authentication.                  |                                                                                                                                                                                                                |
|                |         | API Gateway & Rate Limiting.                | [toolboth](https://github.com/didip/tollbooth) [The Anatomy of an API Gateway in Golang](https://hackernoon.com/the-anatomy-of-an-api-gateway-in-golang)                                                       |
|                |         | Retries, Circuit Breakers, Fault Tolerance. | [Writing a circuit breaker in Go](https://rednafi.com/go/circuit_breaker/)                                                                                                                                     |
|                |         | Distributed Transactions (SAGA Pattern)     | https://github.com/itimofeev/go-saga, https://www.codingexplorations.com/blog/implementing-the-saga-pattern-in-go-a-practical-guide , https://dev.to/yedf2/how-to-implement-saga-pattern-in-microservices-2gj3 |

---

## Features

1. **Order CRUD**: Create, retrieve, and manage orders.
2. **Pagination**: Efficient handling of large datasets.
3. **Messaging**: Event-driven architecture with RabbitMQ.
4. **Authentication**: Secure APIs with JWT.
5. **Database Integration**: PostgreSQL with connection pooling and migrations.

## Project possible tree view (In progress.)

```
📁 order-service/
├── 📁 cmd/
│   └── 📄 main.go                # Application entry point
├── 📁 config/
│   ├── 📄 config.go              # Configuration structs
│   ├── 📄 mongodb.go             # MongoDB setup
│   └── 📄 rabbitmq.go            # RabbitMQ setup
├── 📁 internal/
│   ├── 📁 bootstrap/
│   │   └── 📄 app.go             # Application bootstrapper
│   ├── 📁 usecase/
│   │   └── 📄 order.go           # Order business logic
│   ├── 📁 entity/
│   │   ├── 📄 order.go           # Core order entity
│   │   ├── 📄 order_item.go      # Order items
│   │   ├── 📄 invoice.go         # Invoice generation
│   │   ├── 📄 shipping.go        # Shipping details
│   │   ├── 📄 payment.go         # Payment information
│   │   └── 📄 notification.go     # Notification templates
│   ├── 📁 repository/
│   │   └── 📁 mongodb/
│   │       ├── 📄 order_repo.go   
│   │       ├── 📄 invoice_repo.go 
│   │       ├── 📄 shipping_repo.go
│   │       └── 📄 payment_repo.go
│   ├── 📁 delivery/              # Communication layers
│   │   ├── 📁 http/              # REST API
│   │   │   ├── 📁 middleware/
│   │   │   │   ├── 📄 auth.go    # JWT authentication
│   │   │   │   └── 📄 logger.go  # Request logging
│   │   │   ├── 📁 controller/
│   │   │   │   └── 📄 order.go   # Order HTTP handlers
│   │   │   └── 📁 route/
│   │   │       └── 📄 route.go   # HTTP routes
│   │   ├── 📁 grpc/              # gRPC service
│   │   └── 📁 message/           # Message queue handlers
│   │       ├── 📁 consumer/      # Incoming messages
│   │       │   ├── 📄 payment_events.go    # Payment processing
│   │       │   ├── 📄 inventory_events.go  # Stock updates
│   │       │   ├── 📄 shipping_events.go   # Delivery updates
│   │       │   └── 📄 notification_events.go
│   │       └── 📁 publisher/     # Outgoing messages
│   │           ├── 📄 order_created.go
│   │           ├── 📄 order_updated.go
│   │           └── 📄 order_cancelled.go
│   └── 📁 service/               # Domain services
│       ├── 📄 pricing.go         # Price calculations
│       ├── 📄 tax.go             # Tax calculations
│       ├── 📄 discount.go        # Discount rules
│       └── 📄 fraud.go           # Fraud detection
├── 📁 pkg/
│   ├── 📁 utils/
│   │   ├── 📄 logger.go          # Logging utility
│   │   └── 📄 validator.go       # Input validation
│   └── 📁 messaging/
│       └── 📄 rabbitmq.go        # RabbitMQ helper functions
├── 📄 Dockerfile                 # Service container definition
├── 📄 docker-compose.yml         # Multi-container setup
├── 📄 .env                       # Environment variables
├── 📄 .env.example              # Environment template
├── 📄 Makefile                  # Build automation
└── 📄 README.md                 # Project documentation

```

---

## 🔥 Installation

### Install all required tools:

```sh
go get -u \
    github.com/jackc/pgx/v5 \
    github.com/gin-gonic/gin \
    go.uber.org/zap \
    go.uber.org/fx \
    github.com/ThreeDotsLabs/watermill \
    github.com/golang-jwt/jwt/v4
```

---

## Commands

```bash
go mod tidy # Cleans up the go.mod and go.sum files by adding missing dependencies and removing unused ones
go get go.mongodb.org/mongo-driver/mongo
```

---

## Useful links
- [A Million WebSockets and Go](https://www.freecodecamp.org/news/million-websockets-and-go-cc58418460bb/)
