package main

import (
	"Explore_Go/exercise/demo_jwt/utils"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	jwtService := utils.JWTService{}
	// 通用JWT实现,使用HS256签名算法
	userJWT := &utils.UserJWT{
		Username: "admin",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	t, err := jwtService.GenerateJWT(userJWT, "hello")
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

	claims, err := jwtService.ParseJWT(t, "hello")
	if err != nil {
		panic(err)
	}
	fmt.Println(claims)

	// 通用JWT实现,使用ES256签名算法
	// 读取外部私钥和公钥
	privateKeyBytes, err := os.ReadFile("key/private.pem")
	if err != nil {
		panic(err)
	}

	// 解析私钥
	privateKey, err := jwt.ParseECPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		panic(err)
	}

	// 解析公钥
	publicKey := &privateKey.PublicKey

	t1, err := jwtService.GenerateECDSAJWT(userJWT, privateKey)
	if err != nil {
		panic(err)
	}
	fmt.Println(t1)

	claims1, err := jwtService.ParseECDSAJWT(t1, publicKey)
	if err != nil {
		panic(err)
	}
	fmt.Println(claims1)
}
