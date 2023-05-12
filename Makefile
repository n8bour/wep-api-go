
local-dev:
	@docker-compose -f docker-compose.yml up -d

clean:
	@docker-compose down

run:
	@go build -o bin/app cmd/main.go
	@./bin/app

test:
	@go test ./... -v
