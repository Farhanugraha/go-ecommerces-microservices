
# 🛒 Go E-Commerce Microservices

This project is an E-Commerce platform built using Golang with a microservices architecture, integrated with Midtrans Payment Gateway, and fully containerized using Docker.
The system is designed to be modular, scalable, and maintainable.

# 🚀 Features
Microservices Included :
- Auth Service — User registration, login, and authentication
- Product Service — Product list, stock, and details
- Order Service — Checkout, order creation, and order tracking
- Payment Service (Midtrans) — Payment processing using Midtrans Snap API

Additional Features :
- JWT Authentication
- RESTful API design
- Environment-based configuration
- Dockerized for deployment
Lightweight and scalable architecture

+-------------------+
|    API Gateway    |
+---------+---------+
          |
          v
+---------------------+
|   User Service      |
+---------------------+

+---------------------+
|   Product Service   |
+---------------------+

+---------------------+
|   Order Service     |
+---------------------+

+-----------------------------+
|   Payment Service           |
|   (Midtrans Integration)    |
+-----------------------------+

(All services run as separate Docker containers)




| Component        | Technology                                   |
| ---------------- | -------------------------------------------- |
| Backend          | Golang                                       |
| Architecture     | Microservices                                |
| Containerization | Docker & Docker Compose                      |
| Payment Gateway  | Midtrans Snap API                            |
| Database         | MySQL / PostgreSQL (depending on your setup) |
| Auth             | JWT                                          |

# 💳 Midtrans Payment Flow
1. Order Service sends charge request → Payment Service
2. Payment Service triggers Midtrans Snap API
3. Midtrans returns Snap payment URL
4. User completes payment
5. Midtrans sends callback notification to Payment Service
6. Order status is updated
