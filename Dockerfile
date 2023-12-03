FROM golang:1.21.3 as builder

WORKDIR /usr/src/app

# COPY go.sum ./
COPY go.mod ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o clean

FROM alpine:latest

COPY --from=builder /usr/src/app/clean /clean

CMD ["/clean"]
