.ONESHELL:

build:
	set -e
	sh storage/sqlite/seed.sh
	go build .