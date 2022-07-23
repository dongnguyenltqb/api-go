# Install air 
# go install github.com/cosmtrek/air@latest; fi;

default: develop

develop:
	@air

test:
	@CONFIG_PATH=$${PWD} GIN_MODE=release go test handler/* -v
