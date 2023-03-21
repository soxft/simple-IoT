package jwtutil

import (
	"door/config"
	"door/lib/hash"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"regexp"
	"strconv"
	"time"
)

type JwtClaims struct {
	ID       string `json:"jti,omitempty"`
	ExpireAt int64  `json:"exp,omitempty"`
	IssuedAt int64  `json:"iat,omitempty"`
	Issuer   string `json:"iss,omitempty"`
	jwt.RegisteredClaims
}

// GetJwtFromAuth
// 从 Authorization 中获取JWT
func GetJwtFromAuth(Authorization string) string {
	reg, _ := regexp.Compile(`^Bearer\s+(.*)$`)
	if reg.MatchString(Authorization) {
		return reg.FindStringSubmatch(Authorization)[1]
	}
	return ""
}

// GenerateJwt
// @description generate JWT token for user
func GenerateJwt(clientIp string) (string, error) {
	timeNow := time.Now().Unix()

	return jwt.NewWithClaims(jwt.SigningMethodHS512, JwtClaims{
		ID:       hash.Sha1(strconv.FormatInt(timeNow, 10) + clientIp),
		ExpireAt: time.Now().Add(24 * 30 * time.Hour).Unix(),
		IssuedAt: time.Now().Unix(),
		Issuer:   "xcsoft",
	}).SignedString([]byte(config.Jwt.Secret))
}

// CheckPermission
// @description check user permission
func CheckPermission(_jwt string) error {
	_, err := JwtDecode(_jwt)
	if err != nil {
		return err
	}

	return nil
}

// JwtDecode
// @description check JWT token
func JwtDecode(_jwt string) (JwtClaims, error) {
	token, err := jwt.ParseWithClaims(_jwt, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Jwt.Secret), nil
	})
	if err != nil {
		return JwtClaims{}, err
	}

	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		return *claims, nil
	}

	return JwtClaims{}, errors.New("jwt token error")
}
