# Golang Kubernetes Simple Application

Simple application menggunakan:

- Golang Backend
- Golang Frontend
- PostgreSQL StatefulSet
- Kubernetes Deployment
- Kubernetes Service
- NGINX Ingress
- Distroless Multi Stage Docker Image

Project ini dibuat untuk belajar deployment aplikasi Go ringan di Kubernetes dengan database yang persistent.

---

# Architecture

```
                 Browser
                    |
                    |
              NGINX Ingress
                    |
        +-----------+-----------+
        |                       |
        v                       v

   Frontend Service        Backend Service
      (Go)                   (Go)
        |                       |
        |                       |
        +-----------+-----------+
                    |
                    v

              PostgreSQL
             StatefulSet
                    |
                    |
             Persistent Volume
```

---

# Project Structure

```
.
├── backend
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   └── handler.go
│
├── frontend
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   └── templates
│       └── index.html
│
└── k8s
    ├── postgres.yaml
    ├── backend.yaml
    ├── frontend.yaml
    └── ingress.yaml
```

---

# Technology Stack

| Component | Technology |
|---|---|
| Backend | Golang + Gin |
| Frontend | Golang HTTP Server |
| Database | PostgreSQL |
| Container | Docker |
| Image | Distroless |
| Orchestration | Kubernetes |
| Local Cluster | Kind |
| Ingress | NGINX Ingress Controller |

---

# Backend API

## Health Check

```
GET /api/health
```

Test:

```bash
curl http://localhost:8080/api/health
```

Response:

```json
{
  "status": "ok"
}
```

---

## Get Visitors

```
GET /api/visitors
```

Test:

```bash
curl http://localhost:8080/api/visitors
```

Response:

```json
[
  {
    "id": 1,
    "name": "Aji"
  }
]
```

---

## Create Visitor

```
POST /api/visitors
```

Test:

```bash
curl -X POST \
http://localhost:8080/api/visitors \
-H "Content-Type: application/json" \
-d '{"name":"Aji"}'
```

Response:

```json
{
  "id": 1,
  "name": "Aji",
  "message": "visitor created"
}
```

---

## Delete Visitors

```
DELETE /api/visitors
```

Test:

```bash
curl -X DELETE \
http://localhost:8080/api/visitors
```

---

# Docker Image

## Build Backend

```bash
docker build \
-t ajitirto/image-kecil:backend-1 \
./backend
```

## Build Frontend

```bash
docker build \
-t ajitirto/image-kecil:frontend-1 \
./frontend
```

---

# Image Optimization

Menggunakan multi stage build:

```dockerfile
FROM golang:1.26 AS builder

...

FROM gcr.io/distroless/static-debian12
```

Keuntungan:

- Tidak membawa compiler Go
- Tidak membawa shell
- Image lebih kecil
- Attack surface lebih kecil


Estimasi ukuran:

```
Builder image
~1GB

Final image
~20-30MB
```


---

# Deploy PostgreSQL

```bash
kubectl apply -f k8s/postgres.yaml
```

Check:

```bash
kubectl get pods
```

Expected:

```
postgres-0 Running
```

---

# Deploy Backend

```bash
kubectl apply -f k8s/backend.yaml
```

---

# Deploy Frontend

```bash
kubectl apply -f k8s/frontend.yaml
```

---

# Deploy Ingress

```bash
kubectl apply -f k8s/ingress.yaml
```

---

# Check Resource

## Pods

```bash
kubectl get pods
```

Example:

```
frontend-xxxxx Running
backend-xxxxx Running
postgres-0 Running
```

---

## Services

```bash
kubectl get svc
```

Expected:

```
frontend
backend
postgres
```

---

## Persistent Storage

```bash
kubectl get pvc
```

Example:

```
postgres-data-postgres-0 Bound
```

---

# Local Domain

Tambahkan hosts:

```bash
sudo nano /etc/hosts
```

Tambahkan:

```
127.0.0.1 golang.local
```

Access:

```
http://golang.local
```

---

# Persistence Test

Create data:

```bash
curl -X POST \
http://golang.local/api/visitors \
-H "Content-Type: application/json" \
-d '{"name":"Aji"}'
```

Delete backend:

```bash
kubectl delete deployment backend
```

Deploy kembali:

```bash
kubectl apply -f k8s/backend.yaml
```

Data tetap ada.

---

Delete PostgreSQL Pod:

```bash
kubectl delete pod postgres-0
```

Kubernetes akan membuat ulang:

```
postgres-0 Running
```

Data tetap ada karena menggunakan:

```
StatefulSet
+
PersistentVolumeClaim
```

---

# Cleanup

Delete aplikasi:

```bash
kubectl delete -f k8s/
```

Delete Kind cluster:

```bash
kind delete cluster
```

---

# Learning Goal

Project ini mendemonstrasikan:

- Docker multi stage build
- Golang static binary
- Distroless container image
- Kubernetes Deployment
- Kubernetes Service
- StatefulSet PostgreSQL
- Persistent Volume
- Kubernetes DNS
- NGINX Ingress
- Service communication
- Database persistence
