FROM golang:1.22 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o openai-proxy .

FROM alpine:latest  

WORKDIR /root/
RUN addgroup -S nobody && adduser -S nobody -G nobody
COPY --from=builder /app/openai-proxy .
# Use an unprivileged user
USER nobody:nobody