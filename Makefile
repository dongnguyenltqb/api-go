default: develop

develop:
	@air

test:
	@CONFIG_PATH=$${PWD} GIN_MODE=release go test handler/* -v
