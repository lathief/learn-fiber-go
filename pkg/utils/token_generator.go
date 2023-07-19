package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lathief/learn-fiber-go/pkg/constant"
	"github.com/lathief/learn-fiber-go/pkg/dtos"
	"github.com/lathief/learn-fiber-go/pkg/models"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func GenerateToken(id int, username, role string) string {
	claims := jwt.MapClaims{
		"id":       id,
		"username": username,
		"role":     role,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Second * time.Duration(604800)).Unix(),
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := parseToken.SignedString([]byte("s3cretk3y"))
	if err != nil {
		log.Fatal(err)
	}
	return signedToken
}

func VerifyAccessToken(c *fiber.Ctx) (interface{}, error) {

	headerToken := c.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, constant.ErrUserNeedLogin
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, constant.ErrTokenInvalid
		}
		return []byte("s3cretk3y"), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, constant.ErrTokenInvalid
	}

	return token.Claims.(jwt.MapClaims), nil
}
func VerifyRefreshToken(token dtos.TokenAuth) (interface{}, error) {
	hash := sha1.New()
	_, err := io.WriteString(hash, os.Getenv("SECRET_KEY"))
	if err != nil {
		return nil, err
	}

	user := models.User{}
	salt := string(hash.Sum(nil))[0:16]
	block, err := aes.NewCipher([]byte(salt))
	if err != nil {
		return user, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return user, err
	}

	data, err := base64.URLEncoding.DecodeString(token.RefreshToken)
	if err != nil {
		return user, err
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	plain, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return user, err
	}

	if string(plain) != token.AccessToken {
		return user, constant.ErrTokenInvalid
	}
	getDataToken, _ := jwt.Parse(token.AccessToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, constant.ErrTokenInvalid
		}
		return []byte("s3cretk3y"), nil
	})

	if _, ok := getDataToken.Claims.(jwt.MapClaims); !ok && !getDataToken.Valid {
		return nil, constant.ErrTokenInvalid
	}

	verifyToken := getDataToken.Claims.(jwt.MapClaims)

	return verifyToken["username"].(string), nil
}
func RefreshToken(accessToken string) (string, error) {
	hash := sha1.New()
	_, err := io.WriteString(hash, "s3cretk3y")
	if err != nil {
		return "", err
	}
	salt := string(hash.Sum(nil))[0:16]
	block, err := aes.NewCipher([]byte(salt))
	if err != nil {
		fmt.Println(err.Error())

		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return "", err
	}
	token := base64.URLEncoding.EncodeToString(gcm.Seal(nonce, nonce, []byte(accessToken), nil))
	return token, nil
}
