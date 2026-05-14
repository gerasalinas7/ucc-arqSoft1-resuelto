package auth

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
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
	mode     string
	secret   []byte
	users    map[string]string
	apiKeys  map[string]string
	tokenTTL time.Duration
}

func New(mode, secret string) (Provider, error) {
	if mode != "jwt" && mode != "apikey" {
		return nil, fmt.Errorf("modo de auth no soportado: %s", mode)
	}

	if strings.TrimSpace(secret) == "" {
		secret = "secret-demo"
	}

	return &provider{
		mode:   mode,
		secret: []byte(secret),
		users: map[string]string{
			"admin": "admin123",
			"ana":   "clave123",
		},
		apiKeys: map[string]string{
			"admin": "apikey-admin-123",
			"ana":   "apikey-ana-456",
		},
		tokenTTL: time.Hour,
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

	claims := jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(p.tokenTTL).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(p.secret)
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

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, ErrInvalidToken
		}
		return p.secret, nil
	})
	if err != nil || !parsedToken.Valid {
		return "", ErrInvalidToken
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", ErrInvalidToken
	}

	username, ok := claims["sub"].(string)
	if !ok || strings.TrimSpace(username) == "" {
		return "", ErrInvalidToken
	}

	return username, nil
}
