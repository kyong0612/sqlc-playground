.PHONY: init
init:
	direnv allow .
	cp .env.sample .env
	echo 'fill .env value'

.PHONY: gen-sqlc
gen-sqlc:
	go run github.com/sqlc-dev/sqlc/cmd/sqlc generate

.PHONY: test
test:
	go test -v ./...
