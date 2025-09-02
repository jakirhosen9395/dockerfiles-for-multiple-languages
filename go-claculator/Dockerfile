FROM golang:1.22-bullseye
WORKDIR /app
COPY . .
RUN go build -o app main.go
EXPOSE 8080
CMD ["./app"]
