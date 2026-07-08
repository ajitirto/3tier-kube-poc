FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum* .

RUN go mod download

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build \
    -ldflags="-s -w" \
    -trimpath \
    -o backend .

########################################

FROM gcr.io/distroless/static-debian12

WORKDIR /

COPY --from=builder /app/backend .

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/backend"]
