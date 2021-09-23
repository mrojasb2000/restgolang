# Greenlinght Project

## Create home folder
```
$ mkdir $HOME/Projects/greenlight
```

## Enable modules
```
$ go mod init
```

## High-level skeleton structure

* The **bin** directory will contain our compiled application binaries, ready for deployment to a production server.

* The **cmd/api** directory will contain the application-specific code for our Greenlight API. This will include the code for running the server, reading and writing HTTP requests, and managing authentication.

* The **internal** directory will contain various ancillary packages used by our API. It will contain the code for interacting with our database, doing data validation, sending emails and so on. Basically, any code which isn't application-specific and can potencially be reused will live in here. Our Go code under **cmd/api** will import the packages in the **internal** directory.

* The **migrations** directory will contain the SQL migration files for our datbase.

* The **remote** directory will contain the configuration files and setup scripts for our production server.

* The **go.mod** file will declare our project dependencies, versions and module path.

* The **Makefile** will contain recipes for automating common administrative tasks -- like auditing our Go code, building binaries, and executing database migrations.

It's important to point out that the directory name **internal** carries a special meaning and behavior in Go: any packages which live under this directory can only be imported by code *inside* the parent of the **internal** directory. In our case, this means that any packages which live in **internal** can only be imported by code inside our **greenlight** project directory.

Or, looking at it the other way, this means that any packages under **internal** cannot be *imported by code outside of our project*

This is useful because it prevents other codebases from importing and relying on the (potencially unversioned and unsupported) packages in our **internal** directory -- even if the project code is publicly available somewhere like Github.

## A basic HTTP Server

| URL Pattern | Handler | Action |
| ----------- | ----------- | ----------- |
| /v1/healthcheck | healthcheck | Show application information |

### Test access to APIs
Open your browser and visiting http://localhost:4000/v1/healthcheck
```
status: available
environment: development
version: 1.0.0
```

Or you can use *curl* to make request from your terminal

```
$ curl -i localhost:4000/v1/healthcheck

HTTP/1.1 200 OK
Date: Thu, 23 Sep 2021 23:02:02 GMT
Content-Length: 58
Content-Type: text/plain; charset=utf-8

status: available
environment: development
version: 1.0.0
```
