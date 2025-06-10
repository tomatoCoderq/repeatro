package security

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Security struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey *ecdsa.PublicKey
	ExpirationDelta time.Duration
}

type CustomClaims struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims // includes exp, nbf, iat, etc.
}

func (s *Security) GenerateKey() (error){
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return err
	}

	s.PrivateKey = key
	s.PublicKey = &key.PublicKey

	return  nil
}

func (s *Security) EncodeString(input string, user_id uuid.UUID) (string, error) {
	claims := CustomClaims{
		UserID: user_id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.ExpirationDelta)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	token, err := unsignedToken.SignedString(s.PrivateKey)
	if err != nil {
		return "", nil
	}
	return token, nil
}

func (s *Security) DecodeToken(token string) (CustomClaims, error) {
	claimsToGet := &CustomClaims{}
	parser := jwt.NewParser(jwt.WithLeeway(0 * time.Second))
	tokenGot, err := parser.ParseWithClaims(token, claimsToGet, func(token *jwt.Token) (interface{}, error) {
        return &s.PublicKey, nil
    })
	if err != nil {
		return CustomClaims{}, err
	}

	if !tokenGot.Valid {
		return CustomClaims{}, fmt.Errorf("invalid token")
	}

	return *claimsToGet, nil
}