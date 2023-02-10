ifdef VERSION
        @echo 'VERSION: ${VERSION}'
else
        VERSION="latest"
endif


.PHONY: all test clean build

coverage:
	mkdir -p ./artifacts
	go test ./internal/... -cover -coverprofile=./artifacts/coverage.out fmt
	go tool cover -html=artifacts/coverage.out -o artifacts/coverage.html
	rm ./artifacts/coverage.out

build:
	go build -o build/c3 main.go

run:
	go run .

clean:
	kubectl delete jobs `kubectl get jobs -o custom-columns=:.metadata.name`
	kubectl delete deployment,service logstash
	kubectl delete deployment,service elasticsearch
	kubectl delete deployment,service kibana
	kubectl delete deployment,service web
	kubectl delete cm hackathon
	kubectl delete cm cm-compose
	

.PHONY: docker
docker:
	@docker buildx build \
		--build-arg BUILD_DATE="$(date)" --build-arg VCS_REF="$(git rev-parse --short HEAD)" \
		--platform linux/amd64,linux/arm64 --tag ipedrazas/c3:${VERSION} . # --progress=plain