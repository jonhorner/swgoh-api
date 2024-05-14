CMD=go
BUILD=$(CMD) build
SAMPLE_BINARY_NAME=bootstrap
FUNCTION_NAME=WuBotTbOperations

all: build package deploy
build:
	@echo "Building $(SAMPLE_BINARY_NAME)"
	env GOOS=linux GOARCH=arm64 $(BUILD) -o $(SAMPLE_BINARY_NAME)
package:
	@echo "Packaging $(SAMPLE_BINARY_NAME).zip"
	zip -j $(SAMPLE_BINARY_NAME).zip $(SAMPLE_BINARY_NAME)
deploy:
	@echo "Deploying to lambda"
	aws lambda update-function-code --function-name $(FUNCTION_NAME) --zip-file fileb://$(CURDIR)/$(SAMPLE_BINARY_NAME).zip
test:
	aws lambda invoke --function-name $(FUNCTION_NAME) /tmp/output.json
clean:
	rm -f sample/$(SAMPLE_BINARY_NAME)
	rm -f sample/$(SAMPLE_BINARY_NAME).zip
