package utils

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type contctKey string

const userKey contctKey = "userID"

func CreateJWT(secret []byte, userID string, role string, jwtExpirationInSeconds int64) (string, error) {
	expiration := time.Now().Add(time.Second * time.Duration(jwtExpirationInSeconds))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    userID,
		"role":      role,
		"expiresAt": expiration.Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func CreateTransactionJWT(secret []byte, workspaceID string, jwtExpirationInSeconds int64) (string, error) {
	expiration := time.Now().Add(time.Second * time.Duration(jwtExpirationInSeconds))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"workspaceID": workspaceID,
		"expiresAt":   expiration.Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func WithJWTAuth(handlerFunc http.HandlerFunc, secret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := getTokenFromReq(r)
		const BEARER_SCHEME = "Bearer "
		if tokenString == "" {
			log.Printf("failed to validate token: token not passed")
			permissionDenied(w, "permission denied")
			return
		}
		newTokenString := tokenString[len(BEARER_SCHEME):]

		token, err := validateToken(newTokenString, secret)
		if err != nil {
			log.Printf("failed to validate token: %v", err)
			permissionDenied(w, err.Error())
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			log.Println("Invalid token")
			permissionDenied(w, "jwt claims failed")
			return
		}

		if claims["expiresAt"].(float64) < float64(time.Now().Unix()) {
			log.Println("Expired token")
			permissionDenied(w, "token has expired")
			return
		}

		str := claims["userID"].(string)

		ctx := r.Context()
		ctx = context.WithValue(ctx, userKey, str)
		r = r.WithContext(ctx)

		handlerFunc(w, r)
	}
}

func WithENCJWTAuth(handlerFunc http.HandlerFunc, secret string, conf []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := getTokenFromReq(r)
		const BEARER_SCHEME = "Bearer "
		if tokenString == "" {
			log.Printf("failed to validate token: token not passed")
			permissionDenied(w, "permission denied")
			return
		}
		newTokenString := tokenString[len(BEARER_SCHEME):]

		newToken, err := GetAESDecrypted(newTokenString, conf)
		if err != nil {
			log.Printf("Failed to create JWT token: %v", err)
			permissionDenied(w, "permission denied")
			return
		}

		token, err := validateToken(string(newToken), secret)
		if err != nil {
			log.Printf("failed to validate token: %v", err)
			permissionDenied(w, err.Error())
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			log.Println("Invalid token")
			permissionDenied(w, "jwt claims failed")
			return
		}

		if claims["expiresAt"].(float64) < float64(time.Now().Unix()) {
			log.Println("Expired token")
			permissionDenied(w, "token has expired")
			return
		}

		str := claims["userID"].(string)

		ctx := r.Context()
		ctx = context.WithValue(ctx, userKey, str)
		r = r.WithContext(ctx)

		handlerFunc(w, r)
	}
}

func getTokenFromReq(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")

	if tokenAuth != "" {
		return tokenAuth
	}

	return ""
}

func validateToken(t string, secret string) (*jwt.Token, error) {
	token, err := jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse JWT token")
	}
	return token, nil
}

func permissionDenied(w http.ResponseWriter, err string) {
	WriteError(w, http.StatusForbidden, err)
}

func GetUserIDFromCtx(ctx context.Context) string {
	userID, ok := ctx.Value(userKey).(string)
	if !ok {
		return ""
	}
	return userID
}
