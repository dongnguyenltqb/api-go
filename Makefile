default: develop

develop:
	@air

test:
	@CONFIG_PATH=$${PWD} ENV=test go test handler/* -v
