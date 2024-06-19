package utils

import (
	"crypto/ecdsa"
	"github.com/golang-jwt/jwt/v5"
)

type JWTGeneric interface {
	GenerateJWT(username, secretKey string) (string, error)
	ParseJWT(tokenstr, secretKey string) (any, error)
	GenerateECDSSAJWT(username, key *ecdsa.PrivateKey) (string, error)
	ParseECDSAJWT(tokenstr string, key *ecdsa.PublicKey) (any, error)
}

type UserJWT struct {
	Username             string `json:"username"`
	jwt.RegisteredClaims        // v5版本新加的方法
}

func (u *UserJWT) GenerateJWT(userJWT *UserJWT, secretKey string) (string, error) {
	//	使用HS256签名算法
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, userJWT)
	s, err := t.SignedString([]byte(secretKey))
	return s, err
}

func (u *UserJWT) ParseJWT(tokenstr, secretKey string) (*UserJWT, error) {
	t, err := jwt.ParseWithClaims(tokenstr, &UserJWT{}, func(token *jwt.Token) (any, error) {
		return []byte(secretKey), nil
	})

	if claims, ok := t.Claims.(*UserJWT); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

type MapJWT struct {
	Username      string `json:"username"`
	jwt.MapClaims        // v4版本的方法
}

func (m *MapJWT) GenerateMapJWT(mapJWT *MapJWT, secretKey string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, mapJWT)
	s, err := t.SignedString([]byte(secretKey))
	return s, err
}

func (m *MapJWT) ParseMapJWT(tokenstr, secretKey string) (*MapJWT, error) {
	t, err := jwt.ParseWithClaims(tokenstr, &MapJWT{}, func(token *jwt.Token) (any, error) {
		return []byte(secretKey), nil
	})

	if claims, ok := t.Claims.(*MapJWT); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func (u *UserJWT) GenerateECDSAJWT(userJWT *UserJWT, key *ecdsa.PrivateKey) (string, error) {
	// 使用ES256签名算法
	t := jwt.NewWithClaims(jwt.SigningMethodES256, userJWT)
	s, err := t.SignedString(key)

	return s, err
}

func (u *UserJWT) ParseECDSAJWT(tokenstring string, key *ecdsa.PublicKey) (*UserJWT, error) {
	t, err := jwt.ParseWithClaims(tokenstring, &UserJWT{}, func(token *jwt.Token) (any, error) {
		return key, nil
	})

	if claims, ok := t.Claims.(*UserJWT); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
