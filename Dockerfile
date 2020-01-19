# Build
FROM golang:1.13-alpine3.11 AS build

WORKDIR /go/src/app
COPY . .
RUN go build sensors_exporter.go

# Runtime
FROM alpine:3.11
COPY --from=build /go/src/app/sensors_exporter /
EXPOSE 9612/tcp
ENTRYPOINT ["/sensors_exporter"]
