FROM golang:1.24

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -v -ldflags='-s -w' -trimpath -o ./home-assistant cmd/main.go
CMD ["./home-assistant"]