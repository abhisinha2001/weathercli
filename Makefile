BINARY_NAME=weathercli

.PHONY: all
all: build

.PHONY: build
build: 
	go build -o $(BINARY_NAME) main.go

.PHONY: run
run: 
	go run main.go $(city)

.PHONY: clean
clean:
	rm -f $(BINARY_NAME)