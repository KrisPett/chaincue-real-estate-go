go version go1.21.3 linux/amd64

# Setup

```
docker-compose up -d
docker exec -it postgres-monolith psql -U admin -d postgres

docker exec -it postgres-monolith psql -U admin -d postgres -c "CREATE DATABASE \"chaincue-real-estate-db\";"
```

# .env

```
touch .env .env.test

Example .env file:

CORS_ALLOW_ORIGINS=http://localhost:3000
OAUTH_CLIENT_ID=client-name
OAUTH_CLIENT_SECRET=...
OAUTH_AUTH_URL=https://example.com/auth/realms/client-name/protocol/openid-connect/auth
OAUTH_TOKEN_URL=https://example.com/auth/realms/client-name/protocol/openid-connect/token
OAUTH_INTROSPECT_URL=https://example.com/auth/realms/client-name/protocol/openid-connect/token/introspect
REDIS_PASSWORD=redis
REDIS_HOST=localhost
REDIS_PORT=6379
POSTGRES_HOST=localhost
POSTGRES_USER=admin
POSTGRES_PASSWORD=admin
POSTGRES_DB=chaincue-real-estate-db
POSTGRES_PORT=5432
```

