build:
	go build

clean:
	rm -f ./bookkeeping.db bookkeeping.db-shm bookkeeping.db-wal

newdb:
	go run . database create

generate:
	ruby ./generate-data.rb ./documents.csv ./segments.csv

populate:
	go run . -v database import --documents ./documents.csv --segments ./segments.csv

resetdb: clean newdb populate

fullinit: clean newdb generate populate

.PHONY: rest
rest:
	go run . -v server run