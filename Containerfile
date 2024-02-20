FROM golang:latest as builder

WORKDIR /build
COPY . .
RUN go build -o WandW -ldflags "-w -s" main.go 

FROM alpine:latest

WORKDIR /opt/w-and-w
COPY --from=builder /build/WandW .
COPY assets ./assets/
CMD ["./WandW"]
