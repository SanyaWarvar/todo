package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/SanyaWarvar/todo-app"
	"github.com/SanyaWarvar/todo-app/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo       repository.Authorization
	salt       string
	tokenTTL   time.Duration
	signingKey string
}

func NewAuthService(repo repository.Authorization, salt, signingKey string, tokenTTL time.Duration) *AuthService {
	return &AuthService{
		repo:       repo,
		salt:       salt,
		signingKey: signingKey,
		tokenTTL:   tokenTTL,
	}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password_hash = s.generatePasswordHash(user.Password_hash)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(s.salt)))
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		&tokenClaims{
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(s.tokenTTL).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
			user.Id,
		},
	)

	return token.SignedString([]byte(s.signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(s.signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}
