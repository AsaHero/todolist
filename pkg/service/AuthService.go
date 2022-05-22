package service

import (
	"crypto/sha1"
	"errors"
	"time"

	"github.com/AsaHero/todolist/db/models"
	"github.com/AsaHero/todolist/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	solt      = "etkjfgtpokb65s4gwrfkv"
	secretKey = "xlkhgairgv654droejrv"
	tokenTTL  = 12 * time.Hour
)

type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateAccount(user models.Users) (int, error) {
	user.Password_hash = s.generatePasswordHash(user.Password_hash)
	return s.repo.CreateAccount(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		int(user.Id),
	})

	return token.SignedString([]byte(secretKey))

}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)

	if !ok {
		return 0, errors.New("token claims not matching")
	}

	return claims.UserID, nil
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return string(hash.Sum([]byte(solt)))
}
