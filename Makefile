PROJECT := marshall-bluetooth-keep

BUILD_ENV   :=
VERSION     := $(shell cat VERSION)
COMMIT      := $(shell git log -1 --pretty=format:"%H")
COMMIT_TIME := $(shell git log -1 --pretty=format:"%at" | xargs -I{} date -d @{} +%Y%m%d%H%M%S)
GO_VERSION  := $(shell go version)
LDFLAGS     := "-X 'main.Version=$(VERSION)' -X 'main.GoVersion=$(GO_VERSION)' -X 'main.Commit=$(COMMIT)' -X 'main.CommitTime=$(COMMIT_TIME)'"
BUILD_ARGS  := -trimpath=true -ldflags ${LDFLAGS}

EXEC        := ${PROJECT}
SRC         := ./main.go
OUTPUT_PATH := ./build/bin

.PHONY: setup build gofmt

setup:
	mkdir -p ${OUTPUT_PATH}
	go env -w GOPROXY=https://goproxy.cn,direct

build: setup
	go vet ./...
	rm -rf ${OUTPUT_PATH}/${EXEC}
	${BUILD_ENV} go build ${BUILD_ARGS} -o ${OUTPUT_PATH}/${EXEC} ${SRC_PATH}

gofmt:
	find . -name '*.go' | xargs -L1 go fmt
