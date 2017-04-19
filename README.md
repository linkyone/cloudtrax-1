# Project status

Currently, the project runs locally, inside docker, and on Heroku. All session requests are denied, and any valid login request is allowed. The `CLOUDTRAX_SECRET` environemnt variable is all that is needed to get a working server running (defaults to `default`).

# Communications

The communication with the server happens in one of 3 ways,

1. Between the Access Point and the server on page `/auth.html`.
2. A single `GET` endpoint to authorize user devices.
3. A callback system, to push data usage messages to an external server.

## Access Point Communications

## Authorization Endpoint

The authorize endpoint accepts an authorized UID, authorization time (in minutes),

##### Example flow:

The user is redirected to the login page. This can either be external or internal, but must retrieve a valid Session ID from the login form

# Running server

* Run the server with the following command:<br>
`docker run -itp 8080:8080 $(docker build -q .)`
* Run the server, with a shell prompt with the following command:<br>
`docker run -itp 8080:8080 $(docker build -q.) /bin/sh`

Server should be available at `http://localhost:8080`.

# Deploying to Heroku

* Documentation for the [Heroku Container Registry](https://devcenter.heroku.com/articles/container-registry-and-runtime)
* Install the plugin, and create a new app<br>
```
$ heroku plugins:install heroku-container-registry
$ heroku container:login
$ heroku create
$ heroku container:push web
$ heroku ps:scale web=1
```

# Environment variables

* `PORT` - Heroku gives this as a predefined variables
* `CLOUDTRAX_SECRET` - this should be set, but defaults to `default`

# External Documentation

* Post about [PostgreSQL key/value storage](http://blog.creapptives.com/post/14062057061/the-key-value-store-everyone-ignored-postgresql)
