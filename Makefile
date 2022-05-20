build-services:
	./scripts/build-services.sh

local-run: build-services
	docker-compose up

local-stop:
	docker-compose down