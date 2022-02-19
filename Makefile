
# #####################################
# Application
run:
	echo 'Docker run start'
	go run cmd/main.go;

build:
	echo 'docker build start'
	docker build  . -t api

test:
	echo "Docker test start"
	go test -cover ./...

# #####################################
# Modules
tidy:
	echo  "Go mod tidy and vendor start"
	go mod tidy
	go mod vendor

# #####################################
# Docker
dup:
	docker-compose up
down:
	docker-compose down