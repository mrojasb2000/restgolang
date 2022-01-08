# Greenlinght Project

## Create home folder
```
$ mkdir $HOME/Projects/greenlight
```

## Enable modules
```
$ go mod init
```

### Running program
```
$ go run ./cmd/api
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

Or alternatively, you can use *curl* to make request from your terminal

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
**Note**: The -i flag in the command above instructs *curl* to disply HTTP response headers as well as the response body.


### Verify command-line flags
Specifying alternative port and environment values when starting the application.

| Flag   | Description      | Default     |
| ------ | ---------------- | ----------: | 
| port   | Port number      | 4000        |
| env    | Environment name | development | 


### Running app with params
```
$ go run ./cmd/api -port=3030 -env=production
2021/09/23 20:43:43 starting production server on :3030
```


## API Endpoind and RESTful Routing

Over the next few sections of this book we're going to gradually up out API so that the endpoints start to look like this:


| Method      | URL Pattern | Handler     | Action      |
| ----------- | ----------- | ----------- | ----------- |
| GET         | /v1/healthcheck | healthcheckHandler | Show application informaction |
| GET         | /v1/movies | listMovieHandler | Show the details of all movies |
| POST         | /v1/movies | createMovieHandler | Create a new movies |
| GET         | /v1/movies/:id | showMovieHandler | Show the details of a specific movie |
| PUT         | /v1/movies/:id | editMovieHandler | Update the details of a specific movie |
| DELETE       | /v1/movies/:id | deleteMovieHandler | Delete a specific movie |


## HTTP methods and your use


| Method | Usage |
| ------ | ----- |
| GET | Use for actions that retrieve information only and don't change the state of your application or any data. |
| POST | Use for **non-idempotent** actions that modify state. In the context of a REST API. POST is generally used for actions that create a new resource. |
| PUT | Use for **idempotent** actions that modify the state of a resource at a specify URL. In the context of a REST API, PUT is generally used for actions that replace or update an existing resource. |
| PATCH | Use for actions that partially update a resource at a specific URL. It's OK for the action to be either **idempotent** or **non-idempotent**. |
| DELETE | Use for actions that delete a resoource at a specific URL. |



## Test Movies enpoints

Test healthcheck enpoint API.
```
$ curl localhost:4000/v1/healthcheck
status: available
environment: development
version: 1.0.0
```
Test create movies
```
$ curl -X POST localhost:4000/v1/movies
create a new movie
```
Test show movies
```
$ curl localhost:4000/v1/movies/123
show the details of movie 123
```