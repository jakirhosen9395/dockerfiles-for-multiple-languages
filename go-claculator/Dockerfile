# syntax=docker/dockerfile:1

# -------- Build stage --------
FROM golang:1.22-alpine AS build
ENV CGO_ENABLED=0
WORKDIR /app
COPY . .
RUN go build -trimpath -ldflags="-s -w" -o calculator

# -------- Runtime stage --------
FROM alpine:3.19
WORKDIR /app
COPY --from=build /app/calculator .
COPY --from=build /app/index.html .
ENV PORT=8080
EXPOSE 8080
CMD ["./calculator"]
