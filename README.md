# Environment Variables

Everything in the server is configured with environment variables from the
hosted system. If this is Heroku, Environment Variables are saved in the app
settings section of the online dashboard. If this app is running locally, all
variables are pulled from the current environment. All variables are prefixed
with `CLOUDTRAX_SERVER_`, but not included in the variable name below for
brevity.

| Name          | Required | Description                                      |
| ------------- | -------- | ------------------------------------------------ |
| `PORT`        | YES      | The port number for the server to run on [^1]    |
| `DATABASEURI` | YES      | A complete and valid Postgres connection URI     |
| `SECRET`      | NO       | A salt used for communication with the APs       |
| `DEBUG`       | NO       | Default: `FALSE`, enables debug output [^2]      |

Notes:
[^1]: This defaults to `PORT` without the `CLOUDTRAX_SERVER_` prefix when running on Heroku
[^2]: This can cause undue server load, as it enables debug mode with the underlying SQL library.

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
