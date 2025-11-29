# TreeMOn

**Montlich Version** – Microservices Simulation in Go

TreeMOn is a project designed to demonstrate the differences between **network-based service calls** and **direct function calls**. This project is **educational** and **does not aim to be a production-ready microservice architecture**.

---

## Features

* Monolithic-style gateway that communicates with services via HTTP.
* Separate services for:

  * **Auth Service** – validates users.
  * **Vote Service** – collects votes for `cat` or `dog`.
  * **Main Gateway** – orchestrates calls to Auth and Vote services.
* Logging middleware to track **request duration** for each endpoint.
* Fully in-memory storage for simplicity.
* Easy to test using tools like `curl` or `hurl`.

---

## Endpoints 

### Auth Service (`:8000`)

* **POST /login** – validate username and password
  Request body:

  ```json
  {
    "username": "admin",
    "password": "1234"
  }
  ```

  Response: `200 OK` if valid, `401 Unauthorized` otherwise.

### Vote Service (`:8001`)

* **POST /vote** – submit a vote
* **GET /results** – get current vote counts

Request body for `/vote`:

```json
{
  "username": "admin",
  "password": "1234",
  "animal": "cat"
}
```

### Main Gateway (`:8080`)

* **POST /main_login** – proxy to Auth service
* **POST /main_vote** – proxy to Vote service (requires valid credentials)
* **GET /main_results** – get vote results via Vote service

---

## Logging

All requests through the main gateway are logged with:

* Endpoint name
* HTTP method
* Request path
* Processing duration

Example:

```
[MAIN_LOGIN] POST /main_login processed in 1.234ms
[MAIN_VOTE] POST /main_vote processed in 2.567ms
[MAIN_RESULTS] GET /main_results processed in 0.987ms
```

---

## How to Run

1. Run Auth service:

```bash
cd auth
go run main.go
```

2. Run Vote service:

```bash
cd vote
go run main.go
```

3. Run Main Gateway:

```bash
cd main
go run main.go
```

4. Test endpoints using `curl` or `hurl`.

---

## Example `curl` Commands

* **Login**

```bash
curl -X POST http://localhost:8080/main_login \
-H "Content-Type: application/json" \
-d '{"username":"admin","password":"1234"}'
```

* **Vote**

```bash
curl -X POST http://localhost:8080/main_vote \
-H "Content-Type: application/json" \
-d '{"username":"admin","password":"1234","animal":"cat"}'
```

* **Get Results**

```bash
curl http://localhost:8080/main_results
```

---

## Notes

* This project is meant for **learning purposes**.
* All data is stored **in-memory**; restarting services will reset votes and users.
* Demonstrates **network calls vs function calls** in Go microservices architecture.

---
