
CMDS=build
BINARY=gan-circle-server

all: build
build:
	GOGS=linux go build -o ./bin/${BINARY} ./api
