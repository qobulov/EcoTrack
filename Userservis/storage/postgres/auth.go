package postgres

import (
	"database/sql"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type AuthServer struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) *AuthServer {
	return &AuthServer{db: db}
}

// Yangi foydalanuvchini yaratish
func (s *AuthServer) CreateUser(username, email, password string) error {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return err
	}
	query := "INSERT INTO users (username, email, password_hash, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"
	_, err = s.db.Exec(query, username, email, hashedPassword, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}

// Foydalanuvchini autentifikatsiya qilish
func (s *AuthServer) AuthenticateUser(email, password string) (string, error) {
	var storedHash string
	query := "SELECT password_hash FROM users WHERE email = $1"
	err := s.db.QueryRow(query, email).Scan(&storedHash)
	if err != nil {
		return "", err
	}
	if !checkPasswordHash(password, storedHash) {
		return "", errors.New("foydalanuvchi yoki parol noto'g'ri")
	}
	token, err := generateToken(email)
	if err != nil {
		return "", err
	}
	return token, nil
}

// Tokenni bekor qilish
func (s *AuthServer) InvalidateToken(token string) error {
	// Bu yerda tokenni bekor qilish logikasi amalga oshirilishi mumkin
	return nil
}

// Tokenni yangilash
func (s *AuthServer) RefreshToken(token string) (string, error) {
	newToken, err := generateNewToken(token)
	if err != nil {
		return "", err
	}
	return newToken, nil
}

// Parolni xesh qilish
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Parol xeshini tekshirish
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Token yaratish
func generateToken(email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Subject:   email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(claims)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Yangi token yaratish
func generateNewToken(token string) (string, error) {
	claims := &jwt.StandardClaims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return claims, nil
	})
	if err != nil {
		return "", err
	}
	if !tkn.Valid {
		return "", errors.New("token yaroqsiz")
	}
	if time.Until(time.Unix(claims.ExpiresAt, 0)) > 30*time.Second {
		return "", errors.New("tokenni yangilashga hali vaqt bor")
	}
	newToken, err := generateToken(claims.Subject)
	if err != nil {
		return "", err
	}
	return newToken, nil
}
