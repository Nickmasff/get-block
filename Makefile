up:
	[[ ! -f ./var/config/config.local.yml ]] && cp ./var/config/config.dist.yml ./var/config/config.local.yml

run:
	go run get-block

test:
	go test ./... -v

swag:
	swag init --parseDependency --parseInternal --parseDepth=1