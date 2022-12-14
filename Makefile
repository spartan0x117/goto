BINARY_NAME=goto
BINARY_OUT_DIR=out
DEST=/usr/local/bin

all: build test

build:
	go build -o ./${BINARY_OUT_DIR}/${BINARY_NAME} .

clean:
	go clean
	rm ./${BINARY_OUT_DIR}/${BINARY_NAME}

test:
	go test -v ./...

install:
	install -D ${BINARY_OUT_DIR}/${BINARY_NAME} ${DEST}/${BINARY_NAME}