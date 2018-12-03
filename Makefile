CMD=docker
IMG=unused/gocket

all: fetch build

fetch:
	$(CMD) pull $(IMG)

# build for alpine linux
bin/gocket:
	CGO_ENABLED=0 go build -a -installsuffix cgo -o bin/gocket cmd/main.go

build: bin/gocket
	$(CMD) build -t $(IMG) .

push:
	$(CMD) tag $(IMG) $(IMG) && $(CMD) push $(IMG)

sh:
	$(CMD) run --rm -it $(IMG) sh

clean:
	rm -f bin/gocket
	$(CMD) rmi $(IMG)

.PHONY: run all fetch build push sh clean
