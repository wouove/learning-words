ARTIFACTS_DIR=out

export CGO_ENABLED=1
export GOFLAGS=-ldflags=-s -ldflags=-w -ldflags=-X=main.Version=${VERSION}

clean:
	@echo "Cleaning build directory"
	@rm -Rf ./${ARTIFACTS_DIR}
	@mkdir -p ${ARTIFACTS_DIR}

build:
	@echo "Building for local ARCH and OS"
	go build -o ${ARTIFACTS_DIR}/learning-words cmd/main.go

rebuild: clean build

run: build
	@out/learning-words translate