package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewJwtService() *jwtService {
	return &jwtService{}
}

var SecretKey = []byte("Y7dtw0*whw8r6-wg0dsh_&sb0fs0hyw8e2ew8fg8r38rs0fs8")

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{
		"user_id": userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(SecretKey)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	parseToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid Token")
		}

		return SecretKey, nil
	})
	if err != nil {
		return parseToken, err
	}
	return parseToken, nil
}
