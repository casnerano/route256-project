CUR_DIR=$(shell pwd)
BIN_DIR=${CUR_DIR}/bin
GO_VERSION=$(shell go version | perl -nle '/(go\d\S+)/; print $$1;')
SMART_IMPORTS=${BIN_DIR}/smartimports_${GO_VERSION}
LINT_VER=v1.51.1
LINT_BIN=${BIN_DIR}/lint_${GO_VERSION}_${LINT_VER}
PACKAGE=route256/notifications/cmd/app

all: format build test lint

build: bindir
	go build -o ${BIN_DIR}/app ${PACKAGE}

test:
	go test ./...

run:
	go run ${PACKAGE}

lint: install-lint
	${LINT_BIN} run

precommit: format build test lint
	echo "OK"

bindir:
	mkdir -p ${BIN_DIR}

format: install-smartimports
	${SMART_IMPORTS} -exclude internal/mocks

install-lint: bindir
	test -f ${LINT_BIN} || \
		(GOBIN=${BIN_DIR} go install github.com/golangci/golangci-lint/cmd/golangci-lint@${LINT_VER} && \
		mv ${BIN_DIR}/golangci-lint ${LINT_BIN})

install-smartimports: bindir
	test -f ${SMART_IMPORTS} || \
		(GOBIN=${BIN_DIR} go install github.com/pav5000/smartimports/cmd/smartimports@latest && \
		mv ${BIN_DIR}/smartimports ${SMART_IMPORTS})
