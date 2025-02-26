FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/amf cmd/amf/main.go

FROM alpine:3.18
RUN apk add --no-cache ca-certificates tzdata
RUN adduser -D -H -h /app appuser
USER appuser
WORKDIR /app
COPY --from=builder /bin/amf /app/amf
RUN mkdir -p /app/configs
COPY configs/amf/ /app/configs/
EXPOSE 8080 9090
ENTRYPOINT ["/app/amf"]