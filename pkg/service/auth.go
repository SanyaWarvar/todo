package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/SanyaWarvar/todo-app"
	"github.com/SanyaWarvar/todo-app/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "pashatexnik"
	tokenTTL   = time.Hour * 24
	signingKey = "radiance"
)

type tokenClaims struct {
	jwt.StandardClaims
	Userid int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
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
				ExpiresAt: time.Now().Add(tokenTTL).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
			user.Id,
		},
	)

	return token.SignedString([]byte(signingKey))
}
