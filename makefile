GO_GENERATOR_BIN_URL=https://github.com/StephanHCB/go-generator-cli/releases/download/v1.3.2/go-generator-cli_1.3.2_Darwin_x86_64
GO_GENERATOR_BIN=./bin/go-generator-cli

install-gomplate-macos:
	brew install gomplate

install-templating-macos:
	mkdir -p bin
	brew install wget
	wget ${GO_GENERATOR_BIN_URL} -O ${GO_GENERATOR_BIN}
	chmod +x ${GO_GENERATOR_BIN}

hello-service:
# ${GO_GENERATOR_BIN} --generator=./hello --target=./hello --render
	gomplate -f ./hello/prompt.tmpl -o ./hello/prompt.gpt --template ./hello/hello.proto