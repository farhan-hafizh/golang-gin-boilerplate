package authMiddleware

import (
	"encoding/base64"
	"errors"

	"github.com/golang-jwt/jwt"
)

type Service interface {
	GenerateToken(userId int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
	secretKey string
}

func CreateService(secretKey string) *jwtService {
	return &jwtService{secretKey}
}

func (s *jwtService) GenerateToken(userId int) (string, error) {
	// create claim object
	claim := jwt.MapClaims{}
	// initiate claim
	claim["user_id"] = userId
	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// sign token
	signedToken, err := token.SignedString([]byte(s.secretKey))

	if err != nil {
		return signedToken, err
	}

	base64Token := base64.StdEncoding.EncodeToString([]byte(signedToken))

	return base64Token, nil
}

func (s *jwtService) ValidateToken(paramToken string) (*jwt.Token, error) {
	token, err := base64.StdEncoding.DecodeString(paramToken)
	// parse a token with key function that return any type(interface{}) and error
	decodedToken, err := jwt.Parse(string(token), func(t *jwt.Token) (interface{}, error) {
		// check if the token is signed with HMAC method
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		// if not return the invalid token error
		if !ok {
			return nil, errors.New("Invalid token")
		}
		// if ok, retun the byte so it can be used in jwt.Parse
		return []byte(s.secretKey), nil
	})

	if err != nil {
		return decodedToken, err
	}

	return decodedToken, nil
}
