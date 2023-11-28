package utilities

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

func GetToken(token string) string {
	return strings.TrimPrefix(token, "Bearer ")
}

func GetSubFromToken(token string) (string, error) {
	parts := strings.Split(token, ".")

	if len(parts) != 3 {
		return "", fmt.Errorf("invalid token format")
	}

	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return "", fmt.Errorf("error decoding token payload: %v", err)
	}

	var claims map[string]interface{}
	err = json.Unmarshal(payload, &claims)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling token payload: %v", err)
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		return "", fmt.Errorf("token does not contain a 'sub' claim")
	}

	return sub, nil
}
