# SQLC - Compile SQL to type-safe code

## Migration - [golang-migrate](https://github.com/golang-migrate/migrate)
[instalação](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

para criar as migrations basta executar no terminal:
> `migrate create -ext=sql -dir=sql/migrations -seq init`

> `migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/fullcycle" -verbose up | down`
