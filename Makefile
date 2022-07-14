default: develop

develop:
	@if ![[ command -v air ]]; then go install github.com/cosmtrek/air@latest; fi;
	@air

test:
	@CONFIG_PATH=$${PWD} GIN_MODE=release go test handler/* -v
