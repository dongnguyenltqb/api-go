default: develop

develop:
	@air

test:
	@CONFIG_PATH=$${PWD} ENV=test GIN_MODE=release go test handler/* -v
