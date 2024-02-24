.PHONY: run/sprut
run/sprut:
	@go run cmd/sprut/main.go

## Generate server
.PHONY: gen/server
gen/server:
	ogen --package sprut_server --target gen/server/ --clean schemas/sprut/api.yaml --no-client
