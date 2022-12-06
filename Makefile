default: run

build:
	cd cmd/migrate && go build .

run:
	cd cmd/migrate && go run .

install:
	cd cmd/migrate && go install
