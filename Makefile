build:
	go build

clean:
	rm -f ./bookkeeping.db bookkeeping.db-shm bookkeeping.db-wal

newdb:
	go run . database create

populate:
	go run . database import --documents ./documents.csv --segments ./segments.csv

fullinit: clean newdb populate
