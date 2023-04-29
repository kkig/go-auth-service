## Authentication Service with GO

Restful API to authenticante/add user with Postgres database.

### Go workspaces (Go 1.18+)

Starting with Go 1.18, using go.work is recommended for multi-module workspaces. Read [here](https://github.com/golang/tools/blob/master/gopls/doc/workspace.md) for more details.

```
# Initialize go.work
go work init

# Add directory to go.work
go work use ./dir

# Recursively add all directories to go.work
go work use -r .

# Pushes the dependencies in go.work file back into go.mod
go work sync
```

## Start

Start service in `/auth-service` directory by running:

```
go run .
```

or

```
go run main.go
```

## APIs

- **auth/register**: Register new user. Send `username` and `passord` in POST request.

### Request:

```
curl -i -H "Content-Type: application/json" \
    -X POST \
    -d '{"username":"<USERNAME>", "password":"<PASSWORD>"}' \
    http://localhost:8000/auth/register
```

### Response:

```
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Date: Tue, 11 Oct 2022 03:37:08 GMT
Content-Length: 166
{
	"user": {
		"ID": 4,
		"CreatedAt": "2023-04-29T20:07:06.318028+09:00",
		"UpdatedAt": "2023-04-29T20:07:06.318028+09:00",
		"DeletedAt": null,
		"username": "Pochi",
		"Entries": null
	}
}
```

- **auth/login**:

Send `username` and `password` with POST request to log in.

### Request:

```
curl -i -H "Content-Type: application/json" \
    -X POST \
    -d '{"username":"<USERNAME>", "password":"<PASSWORD>"}' \
    http://localhost:8000/auth/login
```

### Response:

```
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Tue, 11 Oct 2022 03:43:08 GMT
Content-Length: 147

{"jwt":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlYXQiOjE2NjU0NjE3ODgsImlhdCI6MTY2NTQ1OTc4OCwiaWQiOjF9.4agGQACwKSZpPCpHeXnoqXfc3WZqYtE8b0SFcoH40uo"}
```

## DB

This project use [PostgresSQL](https://www.postgresql.org/docs/15/app-createdb.html)(RDBM). If you want to try running this project, database configuration wil be required.

Create database with:

```
createdb diary_app
```
