package auth

import (
	"errors"
	"time"
	"ticket-system/internal/models"
	"ticket-system/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	secret string
	store  *repository.MemoryStore
}

func NewService(secret string, store *repository.MemoryStore) *Service {
	return &Service{
		secret: secret,
		store:  store,
	}
}

func (s *Service) Register(username, password string) (models.User, error) {
	if username == "" || password == "" {
		return models.User{}, errors.New("invalid input")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}
	return s.store.CreateUser(username, string(hash))
}

func (s *Service) Authenticate(username, password string) (string, models.User, error) {
	user, err := s.store.GetUserByUsername(username)
	if err != nil {
		return "", models.User{}, errors.New("invalid credentials")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", models.User{}, errors.New("invalid credentials")
	}
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(s.secret))
	if err != nil {
		return "", models.User{}, err
	}
	return signed, user, nil
}

func (s *Service) ValidateToken(tokenString string) (int, error) {
	parsed, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, errors.New("invalid signing method")
		}
		return []byte(s.secret), nil
	})
	if err != nil || !parsed.Valid {
		return 0, errors.New("invalid token")
	}
	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}
	idf, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("invalid user id")
	}
	return int(idf), nil
}