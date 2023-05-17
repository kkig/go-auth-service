## Authentication Service with GO

Restful API to authenticante/add user with Postgres database.

### Go workspaces (Go 1.18+)

Starting with Go 1.18, using go.work is recommended for multi-module workspaces. Read [here](https://github.com/golang/tools/blob/master/gopls/doc/workspace.md) for more details.

```
# Initialize go.work
go work init

# Add directory to go.work
go work use ./dir/to/goMod

# Recursively add all directories to go.work
go work use -r .

# Run in the directory with go.work
# to push the dependencies in go.work file back into go.mod
go work sync
```

### Database set up

Before running service, please create network bridge and volume container.

Create volume container:

```
docker volume create [container name]
```

Create bridge network:

```
docker network create -d bridge [network name]
```

Build and run containers:

```
docker-compose -f [compose-file name] up --build
```

or run in detached mode:

```
docker-compose -f [compose-file name] up --build -d
```

For the intial set up, log into the database engine:

```
docker exec -it [volume container name] ./cockroach sql --insecure
```

and create database and user:

```
CREATE DATABASE [db name];
CREATE USER [user name];
GRANT ALL ON DATABASE [db name] TO [user name];
quit
```

### APIs

In this project, there are endpoints to register and authenticate users. When we authenticate, server issue JWT token.

As security practice, I would recommend adding `scope` to the request to ensure the requester will be authorized to access resources in the scope.

For more information about OAuth 2.0 Auth Code flow, please read [here](https://auth0.com/docs/get-started/apis/scopes/sample-use-cases-scopes-and-claims) for more details.

**auth/register:**

Register new user. Send `email` and `password` in POST request.

Request:

```
curl -i -H "Content-Type: application/json" \
    -X POST \
    -d '{"email":"<USER_EMAIL>", "password":"<PASSWORD_HASH>"}' \
    http://localhost/auth/register
```

Response:

```
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Date: Tue, 11 Oct 2022 03:37:08 GMT
Content-Length: 166

{
	"user": {
		"ID": 865367747815407777,
		"CreatedAt": "2023-04-29T20:07:06.318028+09:00",
		"UpdatedAt": "2023-04-29T20:07:06.318028+09:00",
		"DeletedAt": null,
		"username": "<USERNAME>",
		"Role": 0
	}
}
```

**auth/login:**

Send `username` and `password` with POST request to log in.

Example Request:

```
curl -i -H "Content-Type: application/json" \
    -X POST \
    -d '{"username":"<USERNAME>", "password":"<PASSWORD>"}' \
    http://localhost:8000/auth/login
```

Example Response:

```
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Tue, 11 Oct 2022 03:43:08 GMT
Content-Length: 147

{"jwt":"<JSON_WEB_TOKEN_HERE>"}
```
