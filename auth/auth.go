package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(username string) (map[string]any, error) {
	privateKey, errPrivateKey := loadPrivateKey("certs/jwt/ed25519_private.pem")
	if errPrivateKey != nil {
		return map[string]any{}, errPrivateKey
	}

	expiresAt := time.Now().Add(time.Hour * 24).Unix()
	issuedAt := time.Now().Unix()
	token := jwt.NewWithClaims(
		jwt.SigningMethodEdDSA,
		jwt.MapClaims{
			"username": username,
			"exp":      expiresAt,
			"iat":      issuedAt,
		})
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return map[string]any{}, err
	}

	tokenMap := map[string]any{
		"token":      tokenString,
		"expires_at": time.Unix(expiresAt, 0).Format("2006-01-02 15:04:05 MST"),
		"issued_at":  time.Unix(issuedAt, 0).Format("2006-01-02 15:04:05 MST"),
	}
	return tokenMap, nil
}

func VerifyToken(tokenString string) error {
	publicKey, errPubKey := loadPublicKey("./certs/jwt/ed25519_public.pem")
	if errPubKey != nil {
		return errPubKey
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodEd25519); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

//func JWTMiddleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		tokenString := r.Header.Get("Authorization")
//
//		// Парсим и верифицируем токен
//		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//			// Убедитесь, что метод подписи токена ожидаемый
//			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
//			}
//
//			return []byte(signingSecretKey), nil
//		})
//
//		if err != nil || !token.Valid {
//			http.Error(w, "Unauthorized", http.StatusUnauthorized)
//			return
//		}
//
//		// Токен валиден, продолжаем обработку запроса
//		next.ServeHTTP(w, r)
//	})
//}
