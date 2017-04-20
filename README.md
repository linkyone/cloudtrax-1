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


# Data models

Here's an example of the top level data object being offered by the API. Some
nested objects will have their own endpoints, but all data should be associated
with a user some how.

```json
{
  "users": [
    {
      "id": "{{GUID}}",
      "createdAt": "{{JS representation of date created, example: '2012-04-23T18:25:43.511Z'}}",
      "deletedAt": "{{JS representation of date deleted, example: '2012-04-23T18:25:43.511Z'}}",
      "email" : "{{first email used to sign up with the wifi system}}",
      "firstName" : "{{first name string}}",
      "lastName" : "{{first name string}}",
      "username": "{{generated for the first time from email (most likely)}}",
      "accounts": [
        {
          "id": "{{GUID}}",
          "externalId": "{{The SSO service identifier for this account}}",
          "createdAt": "{{JS representation of date created, example: '2012-04-23T18:25:43.511Z'}}",
          "deletedAt": "{{JS representation of date deleted, example: '2012-04-23T18:25:43.511Z'}}",
          "type": "{{Twitter/Facebook/Google/Microsoft/etc.}}",
          "meta": {
            "{{key}}": "{{most SSO types offer key-value pair values for account information, can be JSONP?}}"
          }
        }
      ],
      "sessions": [
        {
          "id": "{{AP generated Session ID}}",
          "createdAt": "{{JS representation of date, example: '2012-04-23T18:25:43.511Z'}}",
          "nodes": ["{{An array of mac addresses}}"],
          "devices": ["{{An array of mac addresses}}"],
          "ipAddresses": ["{{An array of IPv4 IP addresses}}"],
          "downloadTotal": "{{int32 value for total download}}",
          "uploadTotal": "{{int32 value for total download}}",
          "seconds": "{{int32 total number of seconds the session is alive for}}",
          "timeoutSeconds": "{{int32 default timeout seconds for the session}}"
        }
      ]
    }
  ]
}

```
