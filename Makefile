REPOSITORY := configwatch
VERSION := $(shell date +%Y%m%d-%H%M)

build-docker:
	docker build -t $(REPOSITORY):$(VERSION) \
		--file ./Dockerfile \
		.

deploy:
	helm upgrade --install --debug configwatch helm \
		--set image.tag=$(VERSION)

.PHONY: build-docker deploy
