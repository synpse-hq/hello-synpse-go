
build-amd64:
	docker build --tag quay.io/synpse/hello-synpse-go:latest -f Dockerfile .

push-amd64: build-amd64
	docker push quay.io/synpse/hello-synpse-go:latest