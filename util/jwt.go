package util

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	SECRETKEY = "chatWeb" //私钥
)

//自定义Claims
type CustomClaims struct {
	UserId int64
	jwt.StandardClaims
}

func GenerateToken() string {
	//生成token
	maxAge := 60 * 60 * 24

	//或者用下面自定义claim
	claims := jwt.MapClaims{
		"id":   11,
		"name": "jerry",
		"exp":  time.Now().Add(time.Duration(maxAge) * time.Second).Unix(), // 过期时间，必须设置,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SECRETKEY))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("token: %v\n", tokenString)

	//解析token
	// ret, err := ParseToken(tokenString)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("userinfo: %v\n", ret)
	return tokenString
}

//解析token
func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRETKEY), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
