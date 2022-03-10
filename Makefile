.PHONY: build push

build:
	docker build -t ghcr.io/warehouse-13/camo:latest .

push:
	docker push ghcr.io/warehouse-13/camo:latest
