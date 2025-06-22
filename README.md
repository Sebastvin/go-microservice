# Microservice with SaaS

<p align="center">
  <img src="https://github.com/user-attachments/assets/5dc49efe-01bf-43d4-96f3-9a38cb452e4a" width="200"/>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Status-Under%20Construction-orange?style=for-the-badge&logo=construction" alt="Under Construction"/>
  <img src="https://img.shields.io/badge/Go-1.22.2+-blue?style=for-the-badge&logo=go" alt="Go Version"/>
  <img src="https://img.shields.io/badge/Node.js-18+-green?style=for-the-badge&logo=node.js" alt="Node.js Version"/>
</p>

## 📋 Project Overview

Project is microservices-based system built with Go and React. The system handles order processing, payment processing via Stripe, inventory management, and provides a modern web interface for order management. Simulates single-payment subscriptions.

### 🏗️ Architecture

- **Backend**: Microservices architecture using Go with gRPC communication
- **Frontend**: React + TypeScript + Vite with Tailwind CSS
- **Database**: MongoDB for order storage
- **Message Queue**: RabbitMQ for service communication
- **Service Discovery**: Consul for service registration and discovery
- **Observability**: Jaeger for distributed tracing
- **Payment Processing**: Stripe integration

## 🚀 Prerequisites

Before running this project, ensure you have the following:

### Required Software
- **Go 1.22.2+** - [Download here](https://golang.org/dl/)
- **Node.js 18+** - [Download here](https://nodejs.org/)
- **Docker & Docker Compose** - [Download here](https://docs.docker.com/get-docker/)
- **Air** (for Go hot reloading) - Install with: `go install github.com/cosmtrek/air@latest`

### Required Accounts & API Keys

#### 🔑 Stripe Account & API Keys
1. Create a [Stripe account](https://stripe.com)
2. Get your API keys from the [Stripe Dashboard](https://dashboard.stripe.com/apikeys)
3. You'll need:
   - **Publishable Key** (starts with `pk_`)
   - **Secret Key** (starts with `sk_`)
   - **Webhook Secret** (for webhook verification)

#### 🤖 OpenAI API Key
1. Create an [OpenAI account](https://platform.openai.com/)
2. Get your API key from the [OpenAI API Keys page](https://platform.openai.com/api-keys)
3. The system uses OpenAI's  model for **image editing and style application**
4. You'll need:
   - **API Key** (starts with `sk-`)

## 🛠️ Setup Instructions

### 1. Clone and Setup
```bash
git clone <your-repo-url>
cd <project-name>
```

### 2. Start Infrastructure Services
```bash
docker compose up -d
```

This will start:
- **MongoDB** (port 27017) - Database
- **Mongo Express** (port 8082) - Database admin interface
- **RabbitMQ** (port 5672) - Message queue
- **RabbitMQ Management** (port 15672) - Queue admin interface
- **Consul** (port 8500) - Service discovery
- **Jaeger** (port 16686) - Distributed tracing

### 3. Configure Stripe Webhook
```bash
# Install Stripe CLI
# macOS: brew install stripe/stripe-cli/stripe
# Windows: scoop install stripe
# Linux: See https://stripe.com/docs/stripe-cli

# Forward webhooks to your local payment service
stripe listen --forward-to localhost:8080/webhook
```

### 4. Start Backend Services

Open multiple terminal windows and run each service:

```bash
# Gateway Service
cd gateway
air

#  Orders Service  
cd orders
air

# Payments Service
cd payments
air

#  Stock Service
cd stock
air
```

### 5. Start Frontend
```bash
cd frontend
npm install
npm run dev
```


## 🔧 Development

### Database Access
- **Mongo Express**: http://localhost:8082 (no auth required)

### Service Discovery
- **Consul UI**: http://localhost:8500

### Monitoring
- **Jaeger UI**: http://localhost:16686
