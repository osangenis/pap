GO_GENERATOR_BIN_URL=https://github.com/StephanHCB/go-generator-cli/releases/download/v1.3.2/go-generator-cli_1.3.2_Darwin_x86_64
GO_GENERATOR_BIN=./bin/go-generator-cli

install-gomplate-macos:
	brew install gomplate

install-protolint-macos:
	brew install protolint

proto-lint:
	docker run --volume "$(PWD):/workspace" --workdir /workspace yoheimuta/protolint lint -fix hello

hello-service: proto-lint
	gomplate -f ./hello/prompt.tmpl -o ./hello/prompt.gpt --template ./hello/hello.proto --template ./common/code-rules.tmpl
	go run cmd/openai-render/main.go -output_dir=$(PWD)/hello < ./hello/prompt.gpt > ./hello/result.gpt