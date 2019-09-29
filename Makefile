APP_PATH = /go/src/horns-cli
HORNS_PATH = ~/Code/horns

.DEFAULT:
	@echo 'App targets:'
	@echo
	@echo '    build      compile the app'
	@echo '    init       initialize go modules'
	@echo '    deps       install dependancies'
	@echo '    setup      build image and install dependencies'
	@echo '    install    compile the project and copy the binary to the horns project'
	@echo '    states     run the US Map State component generator'
	@echo '    test       run unit tests'
	@echo

default: .DEFAULT

build:
	go build -i -o horns

init:
	docker-compose run --rm app go mod init

deps:
	go mod tidy
	go mod vendor

setup: init deps

install: deps build
	cp horns $(HORNS_PATH)/horns

states: build
	./horns gen:states "."

test:
	go test ./... -cover
