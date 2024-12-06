# help me init this file

test:
	go test ./...

build:
	DOCKER_BUILDKIT=1 docker build --platform linux/arm64 -t web-api .

run:
	docker-compose up --build

kill: 
	docker kill $(docker ps -q)

clean:
	go clean
	go fmt ./...
	go mod tidy

docker-clean:
	docker system prune -f -a --volumes
	docker rmi $(docker images -a -q)
	docker rm $(docker ps -a -f status=exited -q)

mock:
	mockgen -source=internal/db/db.go -destination=internal/db/mock/mock_db.go -package=mock_db
	mockgen -source=internal/dao/dao.go -destination=internal/dao/mock/mock_dao.go -package=mock_dao
	mockgen -source=internal/controller/controller.go -destination=internal/controller/mock/mock_controller.go -package=mock_controller
