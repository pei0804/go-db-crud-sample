GO=$(shell which go)

.PHONY: p1 p2 p3 p4

p1:
	$(GO) run p1/main.go

p2:
	$(GO) run p2/main.go

p3:
	$(GO) run p3/main.go

p4:
	$(GO) run p4/main.go

test:
	go test -v ./...

deps:
	go get -u github.com/pei0804/scaneo
	go get -u github.com/rubenv/sql-migrate/...

gen:
	cd model && go generate && cd ../

migrate-status:
	sql-migrate status -env=$(ENV) -config=config/db.yml

migrate-up:
	sql-migrate up -env=$(ENV) -config=config/db.yml

migrate-down:
	sql-migrate down -env=$(ENV) -config=config/db.yml

migrate-dryrun:
	sql-migrate up -env=$(ENV) -dryrun -config=config/db.yml
