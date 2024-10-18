BINARY := initdb.sh
CONTAINER_NAME := initcontainer
REPO_NAME := quay.io/helioloureiro/proj_vas


all: $(BINARY) container push

$(BINARY): 

container: $(BINARY)
	docker build \
		-t $(REPO_NAME)/$(CONTAINER_NAME):latest \
		--squash=true \
		--force-rm=true \
		--rm=true \
		.

push:
	docker push $(REPO_NAME)/$(CONTAINER_NAME):latest

clean:
	rm -f $(BINARY)