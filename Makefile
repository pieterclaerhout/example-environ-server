REVISION := 1.3

build:
	@go build -o example-environ-server github.com/twixlmedia/example-environ-server

run: build
	@TWX_ENVIRONMENT=production ./example-environ-server

publish-docker-image: build
	@docker build -t example-environ-server .
	@docker tag example-environ-server pieterclaerhout/example-environ-server:$(REVISION)
	@docker push pieterclaerhout/example-environ-server:$(REVISION)
