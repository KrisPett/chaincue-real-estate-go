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

CORS_ALLOW_ORIGINS=http://localhost:3000
OAUTH_CLIENT_ID=client-name
OAUTH_CLIENT_SECRET=...
OAUTH_AUTH_URL=https://example.com/auth/realms/client-name/protocol/openid-connect/auth
OAUTH_TOKEN_URL=https://example.com/auth/realms/client-name/protocol/openid-connect/token
OAUTH_INTROSPECT_URL=https://example.com/auth/realms/client-name/protocol/openid-connect/token/introspect
```

package main

import (
"sync"
"time"
)

// TokenCache is a simple in-memory cache for token introspection results.
type TokenCache struct {
cache map[string]cacheEntry
mu    sync.Mutex
}

type cacheEntry struct {
value      map[string]interface{}
expiration time.Time
}

// NewTokenCache creates a new TokenCache.
func NewTokenCache() *TokenCache {
return &TokenCache{
cache: make(map[string]cacheEntry),
}
}

// Get retrieves a value from the cache, returning nil if the entry is not present or has expired.
func (c *TokenCache) Get(key string) *map[string]interface{} {
c.mu.Lock()
defer c.mu.Unlock()

	entry, found := c.cache[key]
	if !found || entry.expiration.Before(time.Now()) {
		return nil
	}
	return &entry.value
}

// Set adds or updates a value in the cache.
func (c *TokenCache) Set(key string, value map[string]interface{}, expiration time.Duration) {
c.mu.Lock()
defer c.mu.Unlock()

	c.cache[key] = cacheEntry{
		value:      value,
		expiration: time.Now().Add(expiration),
	}
}

// introspectTokenWithCache performs token introspection with caching.
func introspectTokenWithCache(c *gin.Context, token string, cache *TokenCache) (map[string]interface{}, error) {
// Attempt to retrieve the introspection result from the cache.
if cachedResult := cache.Get(token); cachedResult != nil {
return *cachedResult, nil
}

	// If not found in the cache, perform the actual introspection.
	introspectionResponse, err := introspectToken(c, token)
	if err != nil {
		return nil, err
	}

	// Cache the introspection result for future use.
	cache.Set(token, introspectionResponse, time.Minute) // Adjust the expiration time as needed.

	return introspectionResponse, nil
}
