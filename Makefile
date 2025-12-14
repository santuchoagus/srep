.ONESHELL:

build:
	set -e
	rm -f $(HOME)/.local/share/srep/app.db
	sh storage/sqlite/migrate.sh
	go build .

migrate_build:
	set -e
	test -f $(HOME)/.local/share/srep/app.db
	sh storage/sqlite/migrate.sh
	go build .

.PHONY: build migrate
