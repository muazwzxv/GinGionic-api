package auth

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// AuthDetails DTO for token
type AuthDetails struct {
	AuthUUID string
	UserID   uint64
}

// CreateToken generate jwt
func CreateToken(auth AuthDetails) (string, error) {
	claim := jwt.MapClaims{}
	claim["authorized"] = true
	claim["auth_uuid"] = auth.AuthUUID
	claim["user_id"] = auth.UserID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

// ExtractToken := Extract token from header
func ExtractToken(req *http.Request) string {
	keys := req.URL.Query()
	token := keys.Get("token")

	if token != "" {
		return token
	}

	bearer := req.Header.Get("Authorization")

	splittedToken := strings.Split(bearer, " ")
	if len(splittedToken) == 2 {
		return splittedToken[1]
	}
	return ""
}

// VerifyToken := verify token with api_secret
func VerifyToken(req *http.Request) (*jwt.Token, error) {

	tokenstring := ExtractToken(req)
	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

// ValidateToken := validates the token
func ValidateToken(req *http.Request) error {
	token, err := VerifyToken(req)

	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}

	return nil
}

// ExtractTokenAuth := extracts auth token for DTO
func ExtractTokenAuth(req *http.Request) (*AuthDetails, error) {
	token, err := VerifyToken(req)

	if err != nil {
		return nil, err
	}

	claim, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		authUUID, ok := claim["auth_uuid"].(string)
		if !ok {
			return nil, err
		}

		userID, err := strconv.ParseUint(fmt.Sprintf("%.f", claim["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}

		return &AuthDetails{
			AuthUUID: authUUID,
			UserID:   userID,
		}, nil
	}
	return nil, err

}
