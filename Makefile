dev-up:
	docker-compose -f ./docker-compose.yml up -d

dev-down:
	docker-compose -f ./docker-compose.yml down

dev-clean:
	docker-compose -f ./docker-compose.yml rm -f