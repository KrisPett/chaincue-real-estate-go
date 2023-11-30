go version go1.21.3 linux/amd64

# Setup

```
docker-compose up -d
docker exec -it postgres-monolith psql -U admin -d postgres

docker exec -it postgres-monolith psql -U admin -d postgres -c "CREATE DATABASE \"chaincue-real-estate-db\";"
```

# .env

Example .env file: **dev**

```
touch .env .env.test

Example .env file:

CORS_ALLOW_ORIGINS=http://localhost:3000
OAUTH_CLIENT_ID=client-name
OAUTH_CLIENT_SECRET=...
OAUTH_AUTH_URL=https://example.com/auth/realms/client-name/protocol/openid-connect/auth
OAUTH_TOKEN_URL=https://example.com/auth/realms/client-name/protocol/openid-connect/token
OAUTH_INTROSPECT_URL=https://example.com/auth/realms/client-name/protocol/openid-connect/token/introspect

POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_DB=chaincue-real-estate-db
POSTGRES_USER=admin
POSTGRES_PASSWORD=admin

REDIS_HOST=localhost
REDIS_PASSWORD=redis
REDIS_PORT=6379
```

Example .env file: **test container**

```
touch .env .env.test

Example .env file:

CORS_ALLOW_ORIGINS=http://localhost:3000
OAUTH_CLIENT_ID=client-name
OAUTH_CLIENT_SECRET=...
OAUTH_AUTH_URL=https://example.com/auth/realms/client-name/protocol/openid-connect/auth
OAUTH_TOKEN_URL=https://example.com/auth/realms/client-name/protocol/openid-connect/token
OAUTH_INTROSPECT_URL=https://example.com/auth/realms/client-name/protocol/openid-connect/token/introspect

POSTGRES_HOST=postgres-monolith
POSTGRES_PORT=5432
POSTGRES_DB=chaincue-real-estate-db
POSTGRES_USER=admin
POSTGRES_PASSWORD=admin

REDIS_HOST=chaincue-real-estate-redis
REDIS_PASSWORD=redis
REDIS_PORT=6379
```

