
test:
	export GOPROXY="https://proxy.golang.org" && go test ./...

testv:
	export GOPROXY="https://proxy.golang.org" && go test ./... -v

vet: 
	go vet ./...

migrateup:
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:5433/awstdev_collection?sslmode=disable" -verbose up

migratedown:
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:5433/awstdev_collection?sslmode=disable" -verbose down

migrateforce:
	migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5433/awstdeawstdev_collectionv?sslmode=disable" force 2

# orm volatiletech/sqlboiler mapping at local
generate-models:
	sqlboiler psql --config sqlboiler.toml

# mockery
generate-mock:
	GO111MODULE=on go get github.com/vektra/mockery/.../
	mockery -dir ./internal/repository/ -all -keeptree
