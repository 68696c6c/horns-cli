APP_PATH = /go/src/horns-cli

.DEFAULT:
	@echo 'App targets:'
	@echo
	@echo '    image     build app image for local development'
	@echo '    build     build app image and compile the app'
	@echo '    init      initialize go modules'
	@echo '    deps      install dependancies'
	@echo '    setup     build image and install dependencies'
	@echo '    states    run the US Map State component generator'
	@echo '    test      run unit tests'
	@echo

default: .DEFAULT

image:
	docker build . --target dev -t horns-cli:dev

build:
	docker-compose run --rm app go build -i -o horns-cli

init:
	docker-compose run --rm app go mod init

deps:
	docker-compose run --rm app go mod tidy
	docker-compose run --rm app go mod vendor

setup: image init deps

states: build
	docker-compose run --rm app ./horns-cli gen:states

test:
	docker-compose run --rm app go test ./... -cover
