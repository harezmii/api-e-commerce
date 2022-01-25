package firabaseAuth

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/option"
	"strings"
)

func FirebaseAuthorize() (*auth.Client, error) {
	opt := option.WithCredentialsFile("serviceAccount.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}
	auth, authError := app.Auth(context.Background())
	if authError != nil {
		return nil, authError
	}
	return auth, nil
}

func FirebaseMiddleWare(ctx *fiber.Ctx) error {
	authorize, authError := FirebaseAuthorize()
	if authError != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"statusCode": 401,
			"error":      "Firebase Service Error",
		})
	}
	authToken := ctx.Get("Authorization")
	idToken := strings.TrimSpace(strings.Replace(authToken, "Bearer", "", 1))

	if idToken == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"statusCode": 401,
			"error":      "UnAuthorization Error",
		})
	}

	token, err := authorize.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"statusCode": 401,
			"error":      "UnAuthorization Token Error",
		})
	}
	ctx.Set("UUID", token.UID)
	return ctx.Next()
}

func GetFirebaseUsersID(ctx *fiber.Ctx) string {
	uuid := ctx.Get("UUID")
	return uuid
}
