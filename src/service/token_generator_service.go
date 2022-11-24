package service

import "github.com/thanhpk/randstr"

func generateToken() string {
	token := randstr.String(8)
	return token
}
