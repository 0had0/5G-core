FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/udm cmd/udm/main.go

FROM alpine:3.18
RUN apk add --no-cache ca-certificates tzdata
RUN adduser -D -H -h /app appuser
USER appuser
WORKDIR /app
COPY --from=builder /bin/udm /app/udm
RUN mkdir -p /app/configs
COPY configs/udm/ /app/configs/
EXPOSE 8080 9090
ENTRYPOINT ["/app/udm"]