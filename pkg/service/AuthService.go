package service

import (
	"crypto/sha1"
	"strconv"
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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(tokenTTL).Unix(),
		IssuedAt:  time.Now().Unix(),
		Id: strconv.FormatInt(int64(user.Id), 10),
	})

	return token.SignedString([]byte(secretKey))

}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return string(hash.Sum([]byte(solt)))
}
