default: run

build:
	cd cmd/migrate && go build .

run:
	cd cmd/migrate && go run .

install:
	cd cmd/migrate && go install

clean:
	cd cmd/migrate && go clean

lint:
	golangci-lint run
