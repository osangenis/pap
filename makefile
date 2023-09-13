.EXPORT_ALL_VARIABLES:

GO_GENERATOR_BIN_URL=https://github.com/StephanHCB/go-generator-cli/releases/download/v1.3.2/go-generator-cli_1.3.2_Darwin_x86_64
GO_GENERATOR_BIN=./bin/go-generator-cli

COMMON_TEMPLATES:=--template common/
SERVICE_TEMPLATES:=--template svc/
HELLO_SVC_PATH:=svc/hello

unit-test:
	go test ./...
	
install-gomplate-macos:
	brew install gomplate

install-protolint-macos:
	brew install protolint

proto-lint:
	docker run --volume "$(PWD):/workspace" --workdir /workspace yoheimuta/protolint lint -fix ${HELLO_SVC_PATH}/hello.proto

render-hello-service-prompt: proto-lint
	SVC_NAME=hello \
	SVC_PATH=${HELLO_SVC_PATH} \
	gomplate -f ${HELLO_SVC_PATH}/prompt.tmpl -o ${HELLO_SVC_PATH}/prompt.gpt \
		--template proto-definition=${HELLO_SVC_PATH}/hello.proto \
		${COMMON_TEMPLATES} \
		${SERVICE_TEMPLATES}

hello-service: render-hello-service-prompt 
	go run cmd/openai-render/main.go -output_dir=${HELLO_SVC_PATH} < ${HELLO_SVC_PATH}/prompt.gpt > ${HELLO_SVC_PATH}/result.gpt