BINARY_NAME=goto
BINARY_OUT_DIR=out
DEST=/usr/local/bin

all: build test

build:
	go build -o ./${BINARY_OUT_DIR}/${BINARY_NAME} .

clean:
	go clean
	rm ./${BINARY_OUT_DIR}/${BINARY_NAME}

runserver:
	go run ./main.go server

test:
	go test -v ./...

install:
	install -D ./${BINARY_OUT_DIR}/${BINARY_NAME} ${DEST}/${BINARY_NAME}
	setcap 'CAP_NET_BIND_SERVICE=+ep' ${DEST}/${BINARY_NAME}