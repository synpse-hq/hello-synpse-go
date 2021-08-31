go-mod-download:
	go mod download

binary-arm64: go-mod-download
	CGO_ENABLED=0 GOARCH=arm64 GOOS=linux go build -o app 

binary-arm32: go-mod-download
	CGO_ENABLED=0 GOARCH=arm GOOS=linux go build -o app 

binary-amd64: go-mod-download
	CGO_ENABLED=0 go build -o app 

image-arm64: binary-arm64
	docker build --tag quay.io/synpse/hello-synpse-go:arm64 -f Dockerfile .

image-arm32: binary-arm32
	docker build --tag quay.io/synpse/hello-synpse-go:arm32 -f Dockerfile .

image-amd64: binary-amd64
	docker build --tag quay.io/synpse/hello-synpse-go:amd64 -f Dockerfile .

push-amd64: image-amd64
	docker push quay.io/synpse/hello-synpse-go:amd64

push-arm64: image-arm64
	docker push quay.io/synpse/hello-synpse-go:arm64

push-arm32: image-arm32
	docker push quay.io/synpse/hello-synpse-go:arm32


binary-all: binary-arm64 binary-arm32 binary-amd64

image-all: image-arm64 image-arm32 image-amd64

push-all: push-arm64 push-arm32 push-amd64

