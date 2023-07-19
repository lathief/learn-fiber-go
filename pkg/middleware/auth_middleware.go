package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
	"github.com/lathief/learn-fiber-go/pkg/constant"
	"github.com/lathief/learn-fiber-go/pkg/handlers"
	"github.com/lathief/learn-fiber-go/pkg/models"
	"github.com/lathief/learn-fiber-go/pkg/utils"
	"net/http"
	"strconv"
)

type Security struct {
	db *sqlx.DB
}
type SecurityInterface interface {
	Authentication(c *fiber.Ctx) error
}

func NewSecurityRepo(DB *sqlx.DB) SecurityInterface {
	return &Security{
		db: DB,
	}
}
func (s *Security) Authentication(c *fiber.Ctx) error {
	verifyToken, err := utils.VerifyAccessToken(c)
	if err != nil {
		return handlers.HandleResponse(c, constant.GetStatusMsg(http.StatusUnauthorized), http.StatusUnauthorized)
	}
	token := verifyToken.(jwt.MapClaims)
	userId := int(token["id"].(float64))
	userName := token["username"].(string)
	var user models.User
	err = s.db.GetContext(c.Context(), &user, "SELECT * FROM users WHERE id = ?", userId)
	if err != nil {
		return handlers.HandleResponse(c, "the user belonging to this token no logger exists", http.StatusForbidden)
	}
	c.Locals("userId", strconv.Itoa(userId))
	c.Locals("username", userName)
	return c.Next()
}
