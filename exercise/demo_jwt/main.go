package main

import (
	"Explore_Go/exercise/demo_jwt/utils"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	// 通用JWT实现,使用HS256签名算法
	userJWT := &utils.UserJWT{
		Username: "admin",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	t, err := userJWT.GenerateJWT(userJWT, "hello")
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

	claims, err := userJWT.ParseJWT(t, "hello")
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
	block, _ := pem.Decode(privateKeyBytes)
	if block == nil || block.Type != "EC PRIVATE KEY" {
		panic("failed to decode PEM block containing private key")
	}
	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	// 解析公钥
	publicKey := &privateKey.PublicKey

	t1, err := userJWT.GenerateECDSAJWT(userJWT, privateKey)
	if err != nil {
		panic(err)
	}
	fmt.Println(t1)

	claims1, err := userJWT.ParseECDSAJWT(t1, publicKey)
	if err != nil {
		panic(err)
	}
	fmt.Println(claims1)
}
