default: develop

develop:
	@air

test:
	@export CONFIG_PATH=$PWD
	@export ENV=test
	@go test handler/* -v