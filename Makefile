VERSION := $(or ${CI_COMMIT_SHORT_SHA},latest)

REPO := $(or ${CI_COMMIT_SHORT_SHA},local)

build-local: build
	docker build -t $(REPO)/fuge:$(VERSION) -t $(REPO)/fuge:latest -f ./Dockerfile .

build-image: build
	docker build -t $(REPO)/web:$(VERSION) -t $(REPO)/web:latest -f ./Dockerfile .
	docker push $(REPO)/web:$(VERSION)
	docker push $(REPO)/web:latest
	docker rmi $(REPO)/web:$(VERSION)

build: install
	CGO_ENABLED=0 GOOS=linux go build -o ./dist/fuge ./app/

install:
	go mod download

doc: doc-v1

doc-v1:
	cd app && swag init --parseDependency --parseInternal -g router.go -dir ./routers/v1 --instanceName v1

env-up:
	docker-compose -f docker/docker-compose-basic.yaml up -d

env-down:
	docker-compose -f docker/docker-compose-basic.yaml down