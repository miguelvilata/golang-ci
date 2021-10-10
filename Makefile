
DOCKER_COMPOSE := docker-compose --file docker-compose.yml
DOCKER_COMPOSE_EXEC := docker-compose -f docker-compose.yml exec -T golang

run:
	$(DOCKER_COMPOSE) up -d --remove-orphans

recreate:
	$(DOCKER_COMPOSE) up -d --build --force-recreate

test:
	make run
	$(DOCKER_COMPOSE_EXEC) bash -c "go get github.com/mfridman/tparse && cd pkg && go test ./... -json -cover | tparse -all"

lint:
	make run
	$(DOCKER_COMPOSE_EXEC) bash -c "go install golang.org/x/lint/golint && golint -set_exit_status ./..."

build:
	make run
	# MacOS
	$(DOCKER_COMPOSE_EXEC) bash -c "cd cmd/geekshubs-library && GOOS=darwin GOARCH=amd64 go build -o ../../bin/main-darwin-amd64 main.go"
	# Linux
	$(DOCKER_COMPOSE_EXEC) bash -c "cd cmd/geekshubs-library && GOOS=linux GOARCH=amd64 go build -o ../../bin/main-linux-amd64 main.go"
	# Windows
	$(DOCKER_COMPOSE_EXEC) bash -c "cd cmd/geekshubs-library && GOOS=windows GOARCH=amd64 go build -o ../../bin/main-windows-amd64 main.go"

check_app:
	make run
	$(DOCKER_COMPOSE_EXEC) bash -c "curl --fail http://localhost:8080/api/"	

shell:
	make run
	$(DOCKER_COMPOSE_EXEC) bash 
