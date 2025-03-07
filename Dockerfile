FROM docker.io/golang:1.24 AS builder
RUN mkdir /app
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o tasks-server -ldflags="-s -w" ./cmd 

FROM docker.io/alpine:latest
RUN mkdir /app && adduser -h /app -D skillsrock
WORKDIR /app
COPY --chown=skillsrock --from=builder /app/.env .
COPY --chown=skillsrock --from=builder /app/tasks-server .
EXPOSE 8080

CMD ["./tasks-server"]
