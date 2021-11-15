APP_REPO ?= quay.io/synpse/hello-synpse-go

.PHONY: image
image:
	docker buildx build --platform linux/amd64,linux/arm64,linux/arm/v7 -t ${APP_REPO}:latest --push -f Dockerfile .
