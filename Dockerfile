FROM golang:buster AS go-build-env

WORKDIR /app

COPY . .

ARG TARGETARCH
RUN CGO_ENABLED=0 GOARCH=$TARGETARCH go build -o /app/app

FROM quay.io/synpse/alpine:3.9

RUN apk --update add ca-certificates

COPY --from=go-build-env /app/app /bin/
COPY ./ui /ui

RUN chmod +x /bin/app

ENTRYPOINT ["/bin/app"]
