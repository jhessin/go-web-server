package main

import (
	"fmt"

	// traffic

	// auth

	"github.com/qor/auth"
	"github.com/qor/auth/claims"
	"github.com/qor/auth/providers/password"

	// mongo
	// bcrypt
	"golang.org/x/crypto/bcrypt"
	// jwt
)

type any = interface{}

var Auth = auth.New(&auth.Config{
	UserModel: User{},
})

func init() {
	Auth.RegisterProvider(password.New(&password.Config{
		Encryptor:        Encryptor{},
		AuthorizeHandler: AuthorizeHandler,
		RegisterHandler:  RegisterHandler,
	}))
}

type Encryptor struct{}

func (e Encryptor) Digest(password string) (string, error) {
	p, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(p), err
}

func (e Encryptor) Compare(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func AuthorizeHandler(ctx *auth.Context) (*claims.Claims, error) {
	// TODO authentication
	return nil, fmt.Errorf("AuthorizeHandler not implemented")
}

func RegisterHandler(ctx *auth.Context) (*claims.Claims, error) {
	// TODO register the authentication
	return nil, fmt.Errorf("RegisterHandler not implemented")
}
