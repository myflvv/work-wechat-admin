package utils

import (
	"crypto/rsa"
	"encoding/json"
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

type Claims struct {
	UserId int `json:"user_id"`
	RoleId int `json:"role_id"`
}

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

//生成
func (c *Claims) CreateToken() (tokenStr string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iat": time.Now().Unix(),                     //颁发时间
		"nbf": time.Now().Unix(),                     //生效时间
		"exp": time.Now().Add(time.Hour * 24).Unix(), //过期时间 24小时
		"iss": "iprun",
		"claims":c,
	})
	return token.SignedString(privateKey)
}

//验证
func ValidateToken(tokenStr string) (claims *Claims, err error) {
	var r_claims *Claims
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (i interface{}, e error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("token加密类型错误")
		}
		return publicKey, nil
	})
	if err != nil {
		return r_claims, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		cc,_:=json.Marshal(claims["claims"])
		_=json.Unmarshal(cc,&r_claims)
		return r_claims, nil
	}
	return r_claims, errors.New("token无效")
}