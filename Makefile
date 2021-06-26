.PHONY: engine
engine: # musl tag will be used by docker alpine image
	GOOS=linux go build -tags musl -a -o engine app/main.go

.PHONY: stop
stop:
	@docker-compose down

# Short test is used for testing the whole unit-test
.PHONY: short-test
short-test:
	@go test -v -cover --short -race ./...

# Full test is used for testing the whole application including the database query directly to a live database
# This may takes time.
.PHONY: full-test
full-test:
	@echo "Running the full test..."
	@go test -v -cover -race ./...

.PHONY: docker-test
docker-test:
	@make clean
	@docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
	@docker-compose -f docker-compose.test.yml down --volumes

.PHONY: clean
clean:
	@make stop
	@docker-compose -f docker-compose.test.yml down --volumes