
local-dev:
	docker-compose -f docker-compose.yml up -d

clean:
	docker-compose down