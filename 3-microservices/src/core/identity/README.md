# wallet.awst.io identity service

- For dev env
    start db: docker-compose up db_awstdev_identity
    migrate db: docker-compose up migrate
    start api: docker-compose up api_awstdev_identity

#### download and use sqlboiler:
  - GO111MODULE=off go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql
  - go get github.com/volatiletech/sqlboiler/v4
  - go get github.com/volatiletech/null/v8
  - link: https://github.com/volatiletech/sqlboiler

#### download and install go migrate:
  - link: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
  - go get github.com/golang-migrate/migrate/v4

#### generate models use sqlboiler
  - make generate-models
  
##### generate db/migrations file
migrate create -ext sql -dir [path_to/migrations] [name_file]

