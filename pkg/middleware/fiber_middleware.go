package middleware

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/lathief/learn-fiber-go/pkg/handlers"
	"os"
	"time"
)

func FiberMiddleware(app *fiber.App) {
	app.Use(
		compress.New(compress.Config{
			Level: compress.LevelBestSpeed, // 0
		}),
		cors.New(cors.Config{
			AllowOrigins: "*",
			AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		}),
		logger.New(logger.Config{
			Format:       "[${pid} - ${time}] ${status} - ${latency} ${method} ${path} ${error}\n",
			TimeFormat:   "15:04:05",
			TimeZone:     "Local",
			TimeInterval: 500 * time.Millisecond,
			Output:       os.Stdout,
		}),
		// You can make own panic handler and save panic log to file with StackTraceHandler
		recover.New(
			recover.Config{
				EnableStackTrace: true,
			},
		),
		// Error Handling Middleware
		func(c *fiber.Ctx) error {
			err := c.Next()

			if err != nil {
				fmt.Println(err.Error())
				if errors.Is(err, errors.New("Method Not Allowed")) {
					return c.Status(fiber.StatusMethodNotAllowed).JSON(handlers.ReturnResponse{
						Code:    405,
						Message: err.Error(),
					})
				}
				// Handle the error
				fmt.Println("Error occurred:", err)
				return c.Status(fiber.StatusInternalServerError).JSON(handlers.ReturnResponse{
					Code:    500,
					Message: err.Error(),
				})
			}
			return nil
		},
	)
}
