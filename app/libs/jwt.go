package libs

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func ValidateJWT(jwtString string) bool {
	_, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
	
		return []byte(CfgSecret()), nil
	})

	return err == nil
}

func DecodeJWT(jwtString string) map[string]string {
	token, _ := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
	
		return []byte(CfgSecret()), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return map[string]string{"userId": claims["userId"].(string)}
	} else {
		return map[string]string{}
	}
}

func EncodeJWT(jwtStruct map[string]string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": jwtStruct["userId"],
	})
	
	tokenString, err := token.SignedString([]byte(CfgSecret()))

	if err != nil { panic(err) }

	return tokenString
}