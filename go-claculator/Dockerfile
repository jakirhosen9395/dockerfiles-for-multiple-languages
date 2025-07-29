FROM golang:tip-bullseye

WORKDIR /go-claculator

COPY . .

RUN go build -o jakir main.go

CMD ["./jakir"]