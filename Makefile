docker-build:
	docker buildx build --platform linux/amd64,linux/arm64 -t choonsiong/go-rest-api .

docker-hub:
	docker push choonsiong/go-rest-api