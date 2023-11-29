package configs

import (
	"chaincue-real-estate-go/internal/utilities"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"golang.org/x/oauth2"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func AuthenticateRoutes(c *gin.Context, roles ...string) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		return
	}

	token := utilities.TrimAndGetToken(authHeader)

	introspectionResponse, err := getFromCache(c, token)
	if err != nil {
		introspectionResponse, err = introspectToken(c, token)
		if err != nil {
			return
		}
		err = cacheIntrospectionResponse(c, token, introspectionResponse)
		if err != nil {
			log.Printf("Error caching introspection response: %v", err)
		}
	}

	if err := checkTokenValidity(introspectionResponse, c); err != nil {
		return
	}

	if err := checkRolesAccess(introspectionResponse, roles, c); err != nil {
		return
	}

	username, _ := introspectionResponse["sub"].(string)
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Authenticated endpoint for user: %s", username)})
}

func getFromCache(ctx context.Context, token string) (map[string]interface{}, error) {
	result, err := redisClient.Get(ctx, token).Result()
	if err == redis.Nil {
		return nil, err
	} else if err != nil {
		log.Printf("Error retrieving from cache: %v", err)
		return nil, err
	}

	var introspectionResponse map[string]interface{}
	err = json.Unmarshal([]byte(result), &introspectionResponse)
	if err != nil {
		log.Printf("Error decoding introspection response from cache: %v", err)
		return nil, err
	}

	return introspectionResponse, nil
}

func cacheIntrospectionResponse(ctx context.Context, token string, response map[string]interface{}) error {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error encoding introspection response: %v", err)
		return err
	}

	err = redisClient.Set(ctx, token, jsonResponse, 5*time.Minute).Err()
	if err != nil {
		log.Printf("Error caching introspection response: %v", err)
		return err
	}

	return nil
}

func introspectToken(c *gin.Context, token string) (map[string]interface{}, error) {
	if err := godotenv.Load(".env"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error loading .env file: %v", err)})
		log.Fatal("Error loading .env file: ", err)
		return nil, err
	}

	oauthConfig := &oauth2.Config{
		ClientID:     os.Getenv("OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("OAUTH_CLIENT_SECRET"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  os.Getenv("OAUTH_AUTH_URL"),
			TokenURL: os.Getenv("OAUTH_TOKEN_URL"),
		},
		Scopes: []string{"openid", "profile", "email"},
	}
	client := oauthConfig.Client(c, &oauth2.Token{AccessToken: token})

	data := url.Values{
		"client_id":     {oauthConfig.ClientID},
		"client_secret": {oauthConfig.ClientSecret},
		"token":         {token},
	}

	introspectURL := os.Getenv("OAUTH_INTROSPECT_URL")
	resp, err := client.Post(introspectURL, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error validating token: %v", err)})
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error closing response body: %v", err)})
		}
	}(resp.Body)

	var introspectionResponse map[string]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&introspectionResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error decoding introspection response: %v", err)})
		return nil, err
	}

	return introspectionResponse, nil
}

func checkTokenValidity(introspectionResponse map[string]interface{}, c *gin.Context) error {
	if valid, ok := introspectionResponse["active"].(bool); !ok || !valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return errors.New("Invalid token")
	}
	return nil
}

func checkRolesAccess(introspectionResponse map[string]interface{}, requiredRoles []string, c *gin.Context) error {
	rolesClaim, ok := introspectionResponse["realm_access"].(map[string]interface{})["roles"].([]interface{})
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Roles not found in the token"})
		return errors.New("Roles not found in the token")
	}

	userRoles := make([]string, len(rolesClaim))
	for i, role := range rolesClaim {
		userRole, ok := role.(string)
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"error": "Invalid role type"})
			return errors.New("Invalid role type")
		}
		userRoles[i] = userRole
	}

	for _, requiredRole := range requiredRoles {
		hasRole := false
		for _, userRole := range userRoles {
			if requiredRole == userRole {
				hasRole = true
				break
			}
		}
		if !hasRole {
			c.JSON(http.StatusForbidden, gin.H{"error": fmt.Sprintf("User does not have the required role: %s", requiredRole)})
			return errors.New("User does not have the required role")
		}
	}

	return nil
}
