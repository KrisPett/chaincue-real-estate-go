package configs

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"io"
	"net/http"
	"net/url"
	"strings"
)

var (
	oauthConfig = &oauth2.Config{
		ClientID:     "real-estate-client",
		ClientSecret: "GIra03FtExlT0A9dZsF6dqqT88JJRi6G",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://auth.chaincuet.com/auth/realms/real-estate/protocol/openid-connect/auth",
			TokenURL: "https://auth.chaincuet.com/auth/realms/real-estate/protocol/openid-connect/token",
		},
		Scopes: []string{"openid", "profile", "email"},
	}
	introspectURL = "https://auth.chaincuet.com/auth/realms/real-estate/protocol/openid-connect/token/introspect"
)

func ProtectRoute(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		return
	}

	parts := strings.Fields(authHeader)
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
		return
	}

	token := parts[1]

	client := oauthConfig.Client(c, &oauth2.Token{AccessToken: token})

	data := url.Values{}
	data.Set("client_id", oauthConfig.ClientID)
	data.Set("client_secret", oauthConfig.ClientSecret)
	data.Set("token", token)

	resp, err := client.Post(introspectURL, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error validating token: %v", err)})
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error closing response body: %v", err)})
			return
		}
	}(resp.Body)

	var introspectionResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&introspectionResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error decoding introspection response: %v", err)})
		return
	}

	if valid, ok := introspectionResponse["active"].(bool); !ok || !valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	username, _ := introspectionResponse["sub"].(string)
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Authenticated endpoint for user: %s", username)})
}
