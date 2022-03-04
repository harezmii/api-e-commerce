package jwtManage

import (
	"api/internal/entity/response"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"
)

type TokenManager interface {
	NewToken(ctx *fiber.Ctx, userId int, role string) string
	ParseToken(ctx *fiber.Ctx, token string) (jwt.Claims, error)
	RefreshToken()
}

var (
	ErrorAlgorithm = errors.New("algorithm error")

	ErrorTokenParam = errors.New("token parameter error")

	ErrorTokenQuery = errors.New("token query error")

	ErrorTokenHeader = errors.New("token header error")

	ErrorTokenForm = errors.New("token form error")

	ErrorTokenCookie = errors.New("token cookie error")
)

type TokenManage struct {
	// sign in algorithm HS256 etc.
	TokenSignInAlgorithm string

	Expires time.Time

	TokenLookup string

	TokenHeaderName string

	TokenKey string

	CookieName string

	CookieSecure bool

	CookieExpires time.Time

	CookieHttpOnly bool

	Authorized func(ctx *fiber.Ctx, code int, message string, token string) error

	UnAuthorized func(ctx *fiber.Ctx, code int, message string) error
}

func (t TokenManage) Initialize() TokenManage {
	if t.CookieName == "" {
		t.CookieName = "jwt"
	}
	if t.TokenLookup != "" {
		t.TokenLookup = "header:Authorization"
	}
	if t.TokenKey == "" {
		t.TokenKey = "secret"
	}
	if t.TokenSignInAlgorithm == "" {
		t.TokenSignInAlgorithm = "HS256"
	}
	t.TokenHeaderName = strings.TrimSpace(t.TokenHeaderName)
	if len(t.TokenHeaderName) == 0 {
		t.TokenHeaderName = "Bearer"
	}

	t.CookieExpires = time.Now().Add(time.Minute * 1)
	t.Expires = time.Now().Add(time.Minute * 1)

	t.CookieSecure = false
	t.CookieHttpOnly = true
	t.UnAuthorized = func(ctx *fiber.Ctx, code int, message string) error {
		return ctx.Status(code).JSON(response.ErrorResponse{StatusCode: code, Message: message})
	}
	t.Authorized = func(ctx *fiber.Ctx, code int, message string, token string) error {
		return ctx.Status(code).JSON(response.SuccessResponse{StatusCode: code, Message: message, Data: token})
	}
	return t
}
func (t TokenManage) TokenAlgorithm() *jwt.SigningMethodHMAC {
	if t.TokenSignInAlgorithm == "HS256" {
		return jwt.SigningMethodHS256
	}
	return nil
}

func (t TokenManage) TokenDetect(ctx *fiber.Ctx, key string, tokenType string) string {
	var detectToken string

	switch tokenType {
	case "header":
		{
			detectToken, _ = t.TokenHeader(ctx, key)
			break
		}
	case "query":
		{
			detectToken, _ = t.TokenQuery(ctx, key)
		}
	case "form":
		{
			detectToken, _ = t.TokenForm(ctx, key)
		}
	case "param":
		{
			detectToken, _ = t.TokenParam(ctx, key)
		}
	case "cookie":
		{
			detectToken, _ = t.TokenQuery(ctx, key)
		}

	}
	return detectToken
}

func (t TokenManage) TokenCookie(ctx *fiber.Ctx, key string) (string, error) {
	token := ctx.Cookies(key)
	if string(token) == "" {
		return "", ErrorTokenCookie
	}
	return string(token), nil
}

func (t TokenManage) TokenHeader(ctx *fiber.Ctx, key string) (string, error) {
	token := ctx.Context().Request.Header.Peek(key)
	if string(token) == "" {
		return "", ErrorTokenHeader
	}
	return string(token), nil
}
func (t TokenManage) TokenParam(ctx *fiber.Ctx, key string) (string, error) {
	token := ctx.Params(key)
	if token == "" {
		return "", ErrorTokenParam
	}
	return token, nil
}
func (t TokenManage) TokenForm(ctx *fiber.Ctx, key string) (string, error) {
	token := ctx.FormValue(key)
	if token == "" {
		return "", ErrorTokenForm
	}
	return token, nil
}
func (t TokenManage) TokenQuery(ctx *fiber.Ctx, key string) (string, error) {
	token := ctx.Query(key)
	if token == "" {
		return "", ErrorTokenQuery
	}
	return token, nil
}
func (t TokenManage) NewToken(ctx *fiber.Ctx, userId int, role string) string {
	claims := jwt.NewWithClaims(t.TokenAlgorithm(), jwt.MapClaims{
		"exp":   t.Expires,
		"id":    userId,
		"login": true,
		"role":  role,
	})
	signedString, err := claims.SignedString([]byte(t.TokenKey))
	if err != nil {
		return ""
	}
	cookie := fiber.Cookie{
		Name:     t.CookieName,
		Value:    signedString,
		Expires:  t.CookieExpires,
		HTTPOnly: t.CookieHttpOnly,
		Secure:   t.CookieSecure,
	}
	ctx.Cookie(&cookie)
	ctx.Context().Response.Header.Set("jwt", signedString)
	return signedString
}
func (t TokenManage) ParseToken(ctx *fiber.Ctx, token string) (jwt.Claims, error) {
	parse, parseError := jwt.Parse(token, jwt.Keyfunc(func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != t.TokenSignInAlgorithm {
			return nil, ErrorAlgorithm
		}
		return []byte(t.TokenKey), nil
	}))
	if parseError != nil {
		return nil, nil
	}
	return parse.Claims, nil
}
func (t TokenManage) RefreshToken() {

}
