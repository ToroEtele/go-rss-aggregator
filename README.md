# RESTful server API written in Go language

## Setup:

## Tools:

The project is using PostgreSQL. The database migrations are done with goose:

```
go install github.com/pressly/goose/v3/cmd/goose@latest
```

To run a migration up/down use:

```
cd sql/schema

goose postgres postgresql://postgres:240301@192.168.88.28:5432/rssagg up
```

The interaction is done with sqlc:

```
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```
