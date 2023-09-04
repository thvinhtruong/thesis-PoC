package utils

import (
	"errors"
	"fmt"
	"server/MainService/config"

	"github.com/golang-jwt/jwt/v4"
)

type JwtUtils struct {
	c config.Config
}

type InfoInJwt struct {
	UserId int
}

func NewJwtUtils(c config.Config) *JwtUtils {
	return &JwtUtils{c: c}
}

func (i *InfoInJwt) convertObjectToMap() map[string]interface{} {
	return map[string]interface{}{
		"UserId": i.UserId,
	}
}

func (i *InfoInJwt) convertMapToObject(m map[string]interface{}) {
	i.UserId = int(m["UserId"].(float64))
}

func (j *JwtUtils) DecodeToken(tokenString string) (*InfoInJwt, error) {
	// parse tokenString
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		hmac_key := fmt.Sprintf("%v", j.c.GetConfig(config.HMAC_KEY))
		return []byte(hmac_key), nil
	})
	if err != nil {
		return nil, err
	}

	claims := token.Claims.(jwt.MapClaims)
	if claims.Valid() != nil {
		return nil, errors.New("token is invalid")
	}
	data := claims["data"].(map[string]interface{})

	var info InfoInJwt
	info.convertMapToObject(data)
	return &info, nil
}

func (j *JwtUtils) GenerateToken(infoInJwt InfoInJwt) (string, error) {

	// Create the JWT claims
	expTime := jwt.TimeFunc().Unix() + 24*60*3600 // expire after 1 day
	claims := jwt.MapClaims{
		"data": infoInJwt.convertObjectToMap(),
		"iss":  "MainService",
		"exp":  expTime,
	}

	// create 2 first part of token: header and payload
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Get secret key to sign 2 first part
	HmacKey := fmt.Sprintf("%v", j.c.GetConfig(config.HMAC_KEY))

	// sign first 2 part and put result to thrid part
	tokenString, err := token.SignedString([]byte(HmacKey))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *JwtUtils) VerifyToken(tokenString string) error {
	// parse tokenString
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		hmac_key := fmt.Sprintf("%v", j.c.GetConfig(config.HMAC_KEY))
		return []byte(hmac_key), nil
	})

	if err != nil {
		return err
	}

	claims := token.Claims.(jwt.MapClaims)
	if claims.Valid() != nil {
		return errors.New("token is invalid")
	}

	return nil
}
