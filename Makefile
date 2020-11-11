build: pack-migrations
	GOOS=linux GARCH=amd64 CGO_ENABLED=0 go build -v -a -installsuffix cgo -o bin ./cmd/jaeger-postgres-storage
	chmod +x cmd/jaeger-postgres-storage

run:
	docker-compose -f hack/docker-compose.yaml up

pack-migrations:
ifeq (, $(shell which go-bindata))
	$(error "No go-bindata in $(PATH), consider installing it from https://github.com/jteeuwen/go-bindata")
endif
	cd assets/migrations && go-bindata -pkg migrations .
