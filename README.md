
# Order Management Service

This module is a part of the **Martix** application responsible for managing customer orders. It integrates with other services (e.g., Product, User, Auth) and ensures scalability, maintainability, and security.

## Commands
```bash
go mod tidy # Cleans up the go.mod and go.sum files by adding missing dependencies and removing unused ones
go get go.mongodb.org/mongo-driver/mongo
```

## Features
1. **Order CRUD**: Create, retrieve, and manage orders.
2. **Pagination**: Efficient handling of large datasets.
3. **Messaging**: Event-driven architecture with RabbitMQ.
4. **Authentication**: Secure APIs with JWT.
5. **Database Integration**: PostgreSQL with connection pooling and migrations.

## Project possible tree view (NEW)
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

## Project possible tree view (OLD)
```
order-management/
├── cmd/
│   └── main.go                          # Application entry point
├── config/
│   ├── config.go                        # Configuration management
│   ├── app.env                          # Environment variables
├── internal/
│   ├── app/                             # Application logic
│   │   ├── order/
│   │   │   ├── dto.go                   # Data Transfer Objects
│   │   │   ├── handler.go               # HTTP or gRPC Handlers
│   │   │   ├── repository.go            # Repository interface and implementation
│   │   │   ├── service.go               # Business logic
│   │   │   ├── validator.go             # Input validation logic
│   │   │   └── events.go                # Event definitions (e.g., OrderCreated)
│   │   └── messaging/
│   │       ├── producer.go              # Publish messages to RabbitMQ/Kafka
│   │       ├── consumer.go              # Consume messages from RabbitMQ/Kafka
│   │       ├── handlers.go              # Handlers for incoming events
│   │       └── message_processor.go     # Event processing and dispatch logic
│   ├── core/                            # Core utilities and extensions
│   │   ├── auth.go                      # Authentication middleware
│   │   ├── error.go                     # Centralized error handling
│   │   ├── pagination.go                # Pagination utility
│   │   ├── logger.go                    # Logger setup
│   │   └── extensions.go                # Generic extensions
│   ├── data/                            # Database layer
│   │   ├── db.go                        # PostgreSQL connection setup
│   │   ├── migrations/                  # SQL migrations
│   │   │   ├── 001_create_orders.sql    # Initial schema for orders
│   │   │   └── 002_create_indexes.sql   # Index creation
│   │   ├── entities.go                  # Database entities
│   │   ├── repository.go                # Generic DB repository utilities
│   │   └── cache.go                     # Redis caching setup
│   ├── proto/                           # gRPC Protobuf definitions
│   │   ├── order.proto                  # gRPC API for orders
│   │   └── generated/                   # Compiled Protobuf files
│   │       ├── order.pb.go
│   │       └── order_grpc.pb.go
│   ├── routes/
│   │   └── routes.go                    # RESTful API route setup
│   └── events/
│       ├── dispatcher.go                # Dispatch events to appropriate handlers
│       ├── event_bus.go                 # Event bus for subscribing and publishing
│       └── handlers/                    # Event handler implementations
│           ├── order_created.go         # OrderCreated event handler
│           └── order_updated.go         # OrderUpdated event handler
├── pkg/                                 # Shared utility packages
│   ├── utils/
│   │   ├── hash.go                      # Password hashing utility
│   │   └── time.go                      # Time utilities
│   ├── logger/
│   │   ├── logger.go                    # Application-wide logger
│   │   └── log_formatter.go             # Custom log formatting
├── .gitignore                           # Git ignore rules
├── Dockerfile                           # Docker configuration for deployment
├── docker-compose.yml                   # Docker Compose for dependencies
├── go.mod                               # Go module file
├── go.sum                               # Dependencies checksum
└── README.md                            # Documentation
```

## Installation
- [Please install protobuff compiler](https://github.com/protocolbuffers/protobuf/releases)

## gRPC
```bash
protoc --go_out=. --go-grpc_out=. internal/app/order/order.proto
```

### 1. Reuse .proto File in Node.js
The .proto file remains the same, serving as the single source of truth. You will use it to generate client and server code for Node.js with NestJS and React.

### 2. Setting Up for NestJS
- Install Dependencies
- Install required gRPC libraries for NestJS:
```bash
npm install @grpc/grpc-js @grpc/proto-loader
```

**NestJS gRPC Server**

Create a grpc-server.module.ts to handle your gRPC server:

```js
// grpc-server.module.ts
import { Module } from '@nestjs/common';
import { GrpcOptions, Transport } from '@nestjs/microservices';
import { join } from 'path';

@Module({})
export class GrpcServerModule {
  static getGrpcConfig(): GrpcOptions {
    return {
      transport: Transport.GRPC,
      options: {
        package: 'order', // Name of the package defined in .proto
        protoPath: join(__dirname, '../proto/order.proto'), // Path to the .proto file
        url: 'localhost:50051',
      },
    };
  }
}

```

**NestJS Service Implementation**
Implement the OrderService defined in the .proto file:
```js
// order.service.ts

import { Injectable } from '@nestjs/common';
import { Observable } from 'rxjs';
import { OrderServiceController, CreateOrderRequest, OrderResponse } from './proto/order';

@Injectable()
export class OrderService implements OrderServiceController {
  createOrder(data: CreateOrderRequest): Observable<OrderResponse> {
    const response: OrderResponse = {
      id: 1,
      customerId: data.customerId,
      status: 'Pending',
      total: data.totalPrice,
      createdAt: new Date().toISOString(),
    };
    return new Observable((subscriber) => {
      subscriber.next(response);
      subscriber.complete();
    });
  }
}


```

### 3. Setting Up for React
Install gRPC-Web
Install the required libraries:
```bash
npm install @grpc/grpc-web google-protobuf
```

**Generate Client Code**
Use the protoc-gen-grpc-web plugin to generate JavaScript/TypeScript client code:
```bash
protoc --proto_path=internal/app/order \
       --js_out=import_style=commonjs:./src/proto \
       --grpc-web_out=import_style=typescript,mode=grpcwebtext:./src/proto \
       internal/app/order/order.proto

```
**This generates:**
- `order_pb.js:` Protobuf message types.
- `order_grpc_web_pb.js`: gRPC-Web client code.

**React Code to Consume gRPC**
Use the generated client to make gRPC calls in React:

```js
import React, { useEffect } from "react";
import { OrderServiceClient } from "./proto/order_grpc_web_pb";
import { CreateOrderRequest, OrderItem } from "./proto/order_pb";

const client = new OrderServiceClient("http://localhost:8080");

const App = () => {
  useEffect(() => {
    const request = new CreateOrderRequest();
    request.setCustomerId("12345");

    const item1 = new OrderItem();
    item1.setProductId("p1");
    item1.setQuantity(2);

    const item2 = new OrderItem();
    item2.setProductId("p2");
    item2.setQuantity(1);

    request.addItems(item1);
    request.addItems(item2);
    request.setPaymentMethod("Credit Card");
    request.setTotalPrice(150.75);

    client.createOrder(request, {}, (err, response) => {
      if (err) {
        console.error("Error:", err.message);
      } else {
        console.log("Order Created:", response.toObject());
      }
    });
  }, []);

  return <div>gRPC React Client</div>;
};

export default App;

```