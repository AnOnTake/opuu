MAKEFLAGS += --silent
GIT_COMMIT:= $(shell git rev-list -1 HEAD)

DATE := `date --utc +'%a %b %_d %H:%M:%S UTC %Y'`

BUILD_NUMBER := $(shell git rev-list --count HEAD)
GIT_HASH := $(shell git rev-parse HEAD)

build:
	go build -ldflags "-s -w -X 'main.BuildNumber=$(BUILD_NUMBER)' -X 'main.BuildDate=$(DATE)' -X 'main.GitHash=$(GIT_HASH)'"

install:
	go build -ldflags "-s -w -X 'main.BuildNumber=$(BUILD_NUMBER)' -X 'main.BuildDate=$(DATE)' -X 'main.GitHash=$(GIT_HASH)'"
	cp ./opuu /usr/local/bin
