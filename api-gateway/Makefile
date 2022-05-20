update-private-dependencies:
	./scripts/update-private-dependencies.sh

generate-swagger:
	swag init -g api/router.go -o api/docs

build-docker:
	docker build -t crud-api-gateway .