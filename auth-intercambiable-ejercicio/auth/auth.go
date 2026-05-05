package auth

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidCredentials = errors.New("credenciales invalidas")
	ErrInvalidToken       = errors.New("token invalido")
)

type Provider interface {
	Name() string
	Login(username, password string) (string, error)
	Validate(token string) (string, error)
}

type provider struct {
	mode    string
	secret  string
	users   map[string]string
	apiKeys map[string]string
}

func New(mode, secret string) (Provider, error) {
	if mode != "jwt" && mode != "apikey" {
		return nil, fmt.Errorf("modo de auth no soportado: %s", mode)
	}

	return &provider{
		mode:   mode,
		secret: secret,
		users: map[string]string{
			"admin": "admin123",
			"ana":   "clave123",
		},
		apiKeys: map[string]string{
			"admin": "apikey-admin-123",
			"ana":   "apikey-ana-456",
		},
	}, nil
}

func (p *provider) Name() string {
	return p.mode
}

func (p *provider) Login(username, password string) (string, error) {
	if p.users[username] != password {
		return "", ErrInvalidCredentials
	}

	if p.mode == "apikey" {
		return p.apiKeys[username], nil
	}

	return "", fmt.Errorf("TODO: generar JWT en Login cuando AUTH_MODE=jwt")
}

func (p *provider) Validate(token string) (string, error) {
	if p.mode == "apikey" {
		for username, apiKey := range p.apiKeys {
			if apiKey == token {
				return username, nil
			}
		}
		return "", ErrInvalidToken
	}

	return "", fmt.Errorf("TODO: validar JWT en Validate cuando AUTH_MODE=jwt")
}
