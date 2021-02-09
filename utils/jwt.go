package utils

import (
	"crypto/rsa"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"time"
)

var (
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
)

func init() {
	publicKeyByte, err := ioutil.ReadFile("cert/jwt_rsa_public.pem")
	if err != nil {
		log.Println(err)
	}
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicKeyByte)
	if err != nil {
		log.Println(err)
	}
	privateKeyByte, err := ioutil.ReadFile("cert/jwt_rsa_private.pem")
	if err != nil {
		log.Println(err)
	}
	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateKeyByte)
	if err != nil {
		log.Println(err)
	}
}

func CreateToken(id int) (tokenStr string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iat": time.Now().Unix(),                     //颁发时间
		"nbf": time.Now().Unix(),                     //生效时间
		"exp": time.Now().Add(time.Hour * 24).Unix(), //过期时间 24小时
		"iss": "iprun",
		"id":  id,
	})
	return token.SignedString(privateKey)
}

func ValidateToken(tokenStr string) (id int, err error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (i interface{}, e error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("token加密类型错误")
		}
		return publicKey, nil
	})
	if err != nil {
		return 0, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return int(claims["id"].(float64)), nil
	}
	return 0, errors.New("token无效")
}
