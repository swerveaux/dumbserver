FROM golang:latest AS builder

WORKDIR /build
COPY . /build
RUN CGO_ENABLED=0 go build ./...

FROM busybox:latest
EXPOSE 8080
WORKDIR /app
COPY --from=builder /build/dumbserver /app/dumbserver

ENTRYPOINT ["./dumbserver"]