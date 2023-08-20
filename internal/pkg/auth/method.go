package auth

import (
	"github.com/golang-jwt/jwt/v4"
)

var methods = map[string]jwt.SigningMethod{
	"256": jwt.SigningMethodHS256,
	"384": jwt.SigningMethodHS384,
	"512": jwt.SigningMethodHS512,
}

func SigningMethod(method string) jwt.SigningMethod {
	m, ok := methods[method]
	if !ok {
		return jwt.SigningMethodHS256
	}
	return m
}
