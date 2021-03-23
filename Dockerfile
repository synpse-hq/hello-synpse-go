FROM golang:1.16
COPY . /
WORKDIR /

RUN go mod download && CGO_ENABLED=0 go build -o app 

FROM alpine:latest
COPY --from=0 /templates /templates
COPY --from=0 /app /app
ENTRYPOINT ["/app"]