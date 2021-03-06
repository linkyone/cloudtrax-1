[![Build Status](https://travis-ci.org/ryanhatfield/cloudtrax.svg?branch=master)](https://travis-ci.org/ryanhatfield/cloudtrax)

# Setup

Everything in the server is configured with environment variables from the
hosted system. If this is Heroku, Environment Variables are saved in the app
settings section of the online dashboard. If this app is running locally, all
variables are pulled from the current environment. All variables are prefixed
with `CLOUDTRAX_SERVER_`, but not included in the variable name below for
brevity.

## Environment Variables

| Name          | Req | Description                                           |
| ------------- | --  | ----------------------------------------------------- |
| `PORT`        | YES | The port for the server to run on. <sub>1</sub>       |
| `DATABASEURI` | YES | A complete and valid Postgres connection URI.         |
| `SECRET`      | YES  | A salt used for communication with the APs.          |
| `DEBUG`       | NO  | Default: `FALSE`, enables debug output. <sub>2</sub>  |
| `MAXDBCONNECTIONS` | NO | Defaults to Heroku default, 20 |

#### Environment Variable Notes:

1. This defaults to `PORT` without the `CLOUDTRAX_SERVER_` prefix when running on Heroku
2. This can cause undue server load, as it enables debug mode with the underlying SQL library.

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

# External Documentation

* Post about [PostgreSQL key/value storage](http://blog.creapptives.com/post/14062057061/the-key-value-store-everyone-ignored-postgresql)

# Routes

| Route | Description |
| ----- | ----------- |
| `/:site/auth.html`<br />(`/joescoffee/auth.html`) | Main authorization endpoint for the access point. This should be entered in the admin interface for each site, including a predetermined site name. |
| `/:site/sessions/:session/:device`<br />(`/joescoffee/01234567/01:02:03:04:05:06/`) | Return a single session from the server, by site name and session token |
| `/:site/sessions/:session/:device/authorize`<br />(`/joescoffee/01234567/01:02:03:04:05:06/authorize`) | Authorize a session, returns HTTP OK if accepted. |
