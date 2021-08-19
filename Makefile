BIN_DIR=_output/bin
RELEASE_VER=v0.1.0
REL_OSARCH="linux/amd64"

all: server

init:
	mkdir -p ${BIN_DIR}

server: init
	go build -v -o=${BIN_DIR}/server ./cmd/server/

clean:
	rm -rf _output/
