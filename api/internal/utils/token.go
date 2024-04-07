package utils

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(userId string, privateKeyStr string, ttl time.Duration) (string, error) {
	privateK, err := ConvertBase64PEMToPrivateKey(privateKeyStr)
	if err != nil {
		return "", err
	}
	token := jwt.New(jwt.SigningMethodES256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(ttl * time.Minute).Unix()
	claims["userid"] = userId
	tokenString, err := token.SignedString(privateK)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string, publicKeyStr string) (*jwt.Token, error) {
	publicKey, err := ConvertBase64PEMToPublicKey(publicKeyStr)
	if err != nil {
		return nil, err
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, err
		}
		return publicKey, nil
	})
	return token, err
}

func ConvertBase64PEMToPrivateKey(base64PEM string) (*ecdsa.PrivateKey, error) {
	pemBytes, err := base64.StdEncoding.DecodeString(base64PEM)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(pemBytes)
	if block == nil || block.Type != "EC PRIVATE KEY" {
		return nil, err
	}
	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}
func ConvertBase64PEMToPublicKey(base64PEM string) (*ecdsa.PublicKey, error) {
	pemBytes, err := base64.StdEncoding.DecodeString(base64PEM)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(pemBytes)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, err
	}
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return publicKey.(*ecdsa.PublicKey), nil
}
