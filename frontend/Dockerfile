FROM golang:1.26 AS builder

WORKDIR /app

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build \
    -trimpath \
    -ldflags="-s -w" \
    -o frontend .

###################################################

FROM gcr.io/distroless/static-debian12

WORKDIR /

COPY --from=builder /app/frontend .
COPY --from=builder /app/templates ./templates

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/frontend"]
