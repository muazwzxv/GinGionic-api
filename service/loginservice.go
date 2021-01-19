package service

import "Go-Learn-API/auth"

type signinInterface interface {
	SignIn(auth.AuthDetails) (string, error)
}

type signinStruct struct{}

var (
	// Authorize := expose for other package
	Authorize signinInterface = &signinStruct{}
)

func (si *signinStruct) SignIn(au auth.AuthDetails) (string, error) {
	token, err := auth.CreateToken(au)
	if err != nil {
		return "", nil
	}

	return token, nil
}
