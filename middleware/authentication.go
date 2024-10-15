package middleware

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"templates/infrastructure"

	jwt "github.com/go-chi/jwtauth"
)

type Authentication interface {
	GetTokenAuth() *jwt.JWTAuth
}

func NewAuthentication() Authentication {
	publicByte, err := ioutil.ReadFile(infrastructure.PathPublicKey)
	if err != nil {
		infrastructure.ErrLog.Fatal(err)
	}
	PublicKeyRS256String := string(publicByte)
	publicKeyBlock, _ := pem.Decode([]byte(PublicKeyRS256String))
	publicKey, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
	if err != nil {
		infrastructure.ErrLog.Fatal(err)
	}

	return &authentication{
		tokenAuthDecode: jwt.New("RS256", publicKey, nil),
	}
}

type authentication struct {
	tokenAuthDecode *jwt.JWTAuth
}

func (a authentication) GetTokenAuth() *jwt.JWTAuth {
	return a.tokenAuthDecode
}
