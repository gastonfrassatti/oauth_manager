
# Oauth Manager

Client for managing the oauths grants of external apps.

This client connects to a mock server created in https://pipedream.com to obtain the grant data.

## Features

- Generate new grants if there isn't any.
- If the grants are expired new ones are generated.
- Saves and updates the data.

## API Reference

#### Obtain grants

```http
  POST /oauthManager/v1/grants
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `client_uuid` | `string` | **Required**. Oauth client id |
| `secret` | `string` | **Required**. Oauth secret |

### Valid request body example

```json
{
"client_uuid":"8aa18595c31f4b489cf773102c48d48d",
"secret":"NqhaBAGNlW0zfquk"
} 
```

Returns grants(a new one or existing one) for the given uuid.

## Run Locally

Clone the project

#### DB
This proyect works with mysql database, if you don't have it installed you can run the docker-compose file in ./docker/db/

The database is "oauth_db" and the table "grants"

I still working in the table migrations, so you can execute manually the script:  ./migrations/000001_init_shcema_up.sql

#### Run project
Run the ./cmd/main.go file
```bash
  go run main.go
```