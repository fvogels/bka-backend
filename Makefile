build:
	go build

clean:
	rm -f ./bookkeeping.db bookkeeping.db-shm bookkeeping.db-wal

newdb:
	go run . database create

populate:
	ruby ./generate-data.rb ./documents.csv ./segments.csv
	go run . -v database import --documents ./documents.csv --segments ./segments.csv

fullinit: clean newdb populate

.PHONY: rest
rest:
	go run . -v server run