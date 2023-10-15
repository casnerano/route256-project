CUR_DIR=$(shell pwd)
BIN_DIR=${CUR_DIR}/bin
GO_VERSION=$(shell go version | perl -nle '/(go\d\S+)/; print $$1;')
SMART_IMPORTS=${BIN_DIR}/smartimports_${GO_VERSION}
LINT_VER=v1.51.1
LINT_BIN=${BIN_DIR}/lint_${GO_VERSION}_${LINT_VER}
PACKAGE=route256/cart/cmd/app

all: format build test lint

build: bindir .install-bin-deps .vendor-proto build-proto
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

.install-bin-deps:
	GOBIN=$(BIN_DIR) go install github.com/vadimi/grpc-client-cli/cmd/grpc-client-cli@latest
	GOBIN=$(BIN_DIR) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(BIN_DIR) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	GOBIN=$(BIN_DIR) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	GOBIN=$(BIN_DIR) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	GOBIN=$(BIN_DIR) go install github.com/envoyproxy/protoc-gen-validate@latest

load-product-service-proto:
	mkdir -p api/client/product/
	${BIN_DIR}/grpc-client-cli -s route256.product.ProductService discover route256.pavl.uk:8082 > api/client/product/product_service.proto

build-proto: load-product-service-proto
	mkdir -p api/v1/openapiv2
	protoc \
		--proto_path=vendor-proto \
		--proto_path=api/v1 \
		--plugin=protoc-gen-go=${BIN_DIR}/protoc-gen-go --go_out="." --go_opt="paths=import" \
		--plugin=protoc-gen-go-grpc=${BIN_DIR}/protoc-gen-go-grpc --go-grpc_out="." --go-grpc_opt="paths=import" \
		--plugin=protoc-gen-openapiv2=${BIN_DIR}/protoc-gen-openapiv2 --openapiv2_out="api/v1/openapiv2" --openapiv2_opt="logtostderr=true" \
		--plugin=protoc-gen-grpc-gateway=$(BIN_DIR)/protoc-gen-grpc-gateway --grpc-gateway_out="." --grpc-gateway_opt="logtostderr=true" --grpc-gateway_opt="paths=import" --grpc-gateway_opt="generate_unbound_methods=true" \
		--plugin=protoc-gen-validate=$(BIN_DIR)/protoc-gen-validate --validate_out="lang=go,paths=import:." \
		api/v1/*.proto

	mkdir -p pkg/proto/client/product
	protoc \
    	--proto_path=vendor-proto \
    	--proto_path=api/client/product \
    	--plugin=protoc-gen-go=${BIN_DIR}/protoc-gen-go --go_out="pkg/proto/client/product" --go_opt="paths=source_relative" \
    	--plugin=protoc-gen-go-grpc=${BIN_DIR}/protoc-gen-go-grpc --go-grpc_out="pkg/proto/client/product" --go-grpc_opt="paths=source_relative" \
		api/client/product/*.proto

	go mod tidy

.vendor-rm:
	rm -rf vendor-proto

.vendor-proto: .vendor-rm vendor-proto/google/api vendor-proto/google/protobuf vendor-proto/protoc-gen-openapiv2/options vendor-proto/validate

vendor-proto/google/api:
	git clone -b master --single-branch -n --depth=1 --filter=tree:0 \
 		https://github.com/googleapis/googleapis vendor-proto/googleapis && \
 	cd vendor-proto/googleapis && \
	git sparse-checkout set --no-cone google/api && \
	git checkout
	mkdir -p  vendor-proto/google
	mv vendor-proto/googleapis/google/api vendor-proto/google
	rm -rf vendor-proto/googleapis

vendor-proto/protoc-gen-openapiv2/options:
	git clone -b main --single-branch -n --depth=1 --filter=tree:0 \
 		https://github.com/grpc-ecosystem/grpc-gateway vendor-proto/grpc-ecosystem && \
 	cd vendor-proto/grpc-ecosystem && \
	git sparse-checkout set --no-cone protoc-gen-openapiv2/options && \
	git checkout
	mkdir -p vendor-proto/protoc-gen-openapiv2
	mv vendor-proto/grpc-ecosystem/protoc-gen-openapiv2/options vendor-proto/protoc-gen-openapiv2
	rm -rf vendor-proto/grpc-ecosystem

vendor-proto/google/protobuf:
	git clone -b main --single-branch -n --depth=1 --filter=tree:0 \
		https://github.com/protocolbuffers/protobuf vendor-proto/protobuf &&\
	cd vendor-proto/protobuf &&\
	git sparse-checkout set --no-cone src/google/protobuf &&\
	git checkout
	mkdir -p vendor-proto/google
	mv vendor-proto/protobuf/src/google/protobuf vendor-proto/google
	rm -rf vendor-proto/protobuf

vendor-proto/validate:
	git clone -b main --single-branch --depth=2 --filter=tree:0 \
		https://github.com/bufbuild/protoc-gen-validate vendor-proto/tmp && \
		cd vendor-proto/tmp && \
		git sparse-checkout set --no-cone validate &&\
		git checkout
		mkdir -p vendor-proto/validate
		mv vendor-proto/tmp/validate vendor-proto/
		rm -rf vendor-proto/tmp

tests:
	go test -cover ./...
tests-cover:
	go test -coverprofile=coverage.out ./... > /dev/null && go tool cover -func=coverage.out
tests-cover-html:	tests-cover
	go tool cover -html=coverage.out