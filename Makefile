
# Image URL to use all building/pushing image targets
REGISTRY ?= quay.io
REPOSITORY ?= $(REGISTRY)/redhat-cop/quay-operator
DEV_TAG ?= dev

IMG := $(REPOSITORY):latest

VERSION := v0.0.5

OPENSHIFT_VERSION=3.10.172
QUAY_NAMESPACE=quay-enterprise
SDK_VERSION=v0.10.0
GOPATH ?= "$(HOME)/go"

BUILD_COMMIT := $(shell ./scripts/build/get-build-commit.sh)
BUILD_TIMESTAMP := $(shell ./scripts/build/get-build-timestamp.sh)
BUILD_HOSTNAME := $(shell ./scripts/build/get-build-hostname.sh)

LDFLAGS := "-X github.com/redhat-cop/quay-operator/version.Version=$(VERSION) \
	-X github.com/redhat-cop/quay-operator/version.Vcs=$(BUILD_COMMIT) \
	-X github.com/redhat-cop/quay-operator/version.Timestamp=$(BUILD_TIMESTAMP) \
	-X github.com/redhat-cop/quay-operator/version.Hostname=$(BUILD_HOSTNAME)"

all: manager

# Run tests
native-test: generate fmt vet
	go test ./pkg/... ./cmd/... -coverprofile cover.out

# Build manager binary
manager: generate fmt vet
	go build -o build/_output/bin/quay-operator  -ldflags $(LDFLAGS) github.com/redhat-cop/quay-operator/cmd/manager

# Build manager binary
manager-osx: generate fmt vet
	go build -o build/_output/bin/quay-operator GOOS=darwin  -ldflags $(LDFLAGS) github.com/redhat-cop/quay-operator/cmd/manager

# Run against the configured Kubernetes cluster in ~/.kube/config
run: generate fmt vet
	go run ./cmd/manager/main.go

# Install CRDs into a cluster
install:
	cat deploy/crds/*crd.yaml | kubectl apply -f-

# Run go fmt against code
fmt:
	go fmt ./pkg/... ./cmd/...

# Run go vet against code
vet:
	go vet ./pkg/... ./cmd/...

# Generate code
generate:
	go generate ./pkg/... ./cmd/...

# Docker Login
docker-login:
	@docker login -u $(DOCKER_USER) -p $(DOCKER_PASSWORD) $(REGISTRY)

# Tag for Dev
docker-tag-dev:
	@docker tag $(IMG) $(REPOSITORY):$(DEV_TAG)

# Tag for Dev
docker-tag-release:
	@docker tag $(IMG) $(REPOSITORY):$(VERSION)
	@docker tag $(IMG) $(REPOSITORY):latest	

# Push for Dev
docker-push-dev:  docker-tag-dev
	@docker push $(REPOSITORY):$(DEV_TAG)

# Push for Release
docker-push-release:  docker-tag-release
	@docker push $(REPOSITORY):$(VERSION)
	@docker push $(REPOSITORY):latest

# Build the docker image
docker-build:
	docker build . -t ${IMG}

# Push the docker image
docker-push:
	docker push ${IMG}

# Travis Latest Tag Deployment
travis-latest-deploy: docker-login docker-build docker-push

# Travis Dev Deployment
travis-dev-deploy: docker-login docker-build docker-push-dev

# Travis Release
travis-release-deploy: docker-login docker-build docker-push-release

#Travis E2E
travis-e2e-tests: install-minishift install-sdk run-go-mod run-unit-tests run-e2e-tests

#Install dependencies
run-go-mod:
	go mod vendor

#Run Unit Tests
run-unit-tests:
	go test -v ./pkg/...

#Run E2E
run-e2e-tests:
	operator-sdk test local ./test/e2e --namespace "quay-enterprise" --up-local --no-setup --verbose --go-test-flags "-timeout 20m"

#Install SDK
install-sdk:
	@echo Installing SDK ${SDK_VERSION}
	@SDK_VERSION=$(SDK_VERSION) GOPATH=$(GOPATH) ./.travis/setup-sdk.sh

#Install Minishift
install-minishift:
	@echo Installing Minishift
	@OPENSHIFT_VERSION=$(OPENSHIFT_VERSION) QUAY_NAMESPACE=$(QUAY_NAMESPACE) ./.travis/setup-minishift.sh