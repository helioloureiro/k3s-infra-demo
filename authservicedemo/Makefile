BINARY := webapp
CONTAINER_NAME := authservicedemo
REPO_NAME := quay.io/helioloureiro/proj_vas

BUILD_OPTIONS = -modcacherw
BUILD_OPTIONS += -ldflags="-s -w -X 'main.Version=$$(git rev-parse HEAD | cut -c1-10)'"
## not running on alpine
# BUILD_OPTIONS += -buildmode=pie

## this options requires CGO enabled
BUILD_OPTIONS_RACE = -race

$(BINARY): check-race main.go go.mod
	export CGO_ENABLED=0 && \
		go build -o $(BINARY) $(BUILD_OPTIONS) .

check-race:
	go run $(BUILD_OPTIONS) $(BUILD_OPTIONS_RACE) . --version

container:
	docker build \
		-t $(REPO_NAME)/$(CONTAINER_NAME):latest \
		--squash=true \
		--force-rm=true \
		--rm=true \
		.

push:
	docker push $(REPO_NAME)/$(BINARY):latest

clean:
	rm -f $(BINARY)