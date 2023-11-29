package utilities

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Claims struct {
	Exp             int64          `json:"exp"`
	Iat             int64          `json:"iat"`
	AuthTime        int64          `json:"auth_time"`
	Jti             string         `json:"jti"`
	Iss             string         `json:"iss"`
	Aud             string         `json:"aud"`
	Sub             string         `json:"sub"`
	Typ             string         `json:"typ"`
	Azp             string         `json:"azp"`
	SessionState    string         `json:"session_state"`
	Acr             string         `json:"acr"`
	RealmAccess     RealmAccess    `json:"realm_access"`
	ResourceAccess  ResourceAccess `json:"resource_access"`
	Scope           string         `json:"scope"`
	Sid             string         `json:"sid"`
	EmailVerified   bool           `json:"email_verified"`
	Name            string         `json:"name"`
	PreferredUserID string         `json:"preferred_username"`
	GivenName       string         `json:"given_name"`
	FamilyName      string         `json:"family_name"`
	Email           string         `json:"email"`
}

type RealmAccess struct {
	Roles []string `json:"roles"`
}

type ResourceAccess struct {
	Account AccountAccess `json:"account"`
}

type AccountAccess struct {
	Roles []string `json:"roles"`
}

func TrimAndGetToken(token string) string {
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

func GenerateToken(privateKey *rsa.PrivateKey) (string, error) {
	claims := Claims{
		Exp:          time.Now().Add(5 * time.Minute).Unix(),
		Iat:          time.Now().Unix(),
		AuthTime:     time.Now().Unix(),
		Jti:          "b6e28ea6-77db-4531-bea5-4bee4eff5736",
		Iss:          "https://auth.chaincue.com/auth/realms/real-estate",
		Aud:          "account",
		Sub:          "d74e8499-376d-454b-8b08-37f948ba6afa",
		Typ:          "Bearer",
		Azp:          "real-estate-client",
		SessionState: "0940ddc5-938e-4f04-8e42-4f2b0eafbfa4",
		Acr:          "1",
		RealmAccess: RealmAccess{
			Roles: []string{
				"offline_access",
				"default-roles-real-estate",
				"admin",
				"uma_authorization",
				"user",
			},
		},
		ResourceAccess: ResourceAccess{
			Account: AccountAccess{
				Roles: []string{
					"manage-account",
					"manage-account-links",
					"view-profile",
				},
			},
		},
		Scope:           "openid email profile",
		Sid:             "0940ddc5-938e-4f04-8e42-4f2b0eafbfa4",
		EmailVerified:   true,
		Name:            "testuser user",
		PreferredUserID: "testuser2@chaincue.com",
		GivenName:       "testuser2@chaincue.com",
		FamilyName:      "testuser2@chaincue.com",
		Email:           "testuser2@chaincue.com",
	}

	claimsJSON, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}

	header := base64.URLEncoding.EncodeToString([]byte(`{"alg":"RS256","typ":"JWT","kid":"uSFw1wVgrNQRj9wg67PQsSnlLWTqNgyG-mfGHCJoo-0"}`))

	payload := base64.URLEncoding.EncodeToString(claimsJSON)

	hashed := sha256.Sum256([]byte(header + "." + payload))
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return "", err
	}

	token := fmt.Sprintf("%s.%s.%s", header, payload, base64.URLEncoding.EncodeToString(signature))
	return token, nil
}
