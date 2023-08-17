package web

import (
	"os"
	"time"

	"github.com/RonkZeDonk/uogcal/pkg/database"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// TODO create a jwt revoke list

type UserLogin struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
}
type UserSignUp struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
}

type ClaimsMap struct {
	User string
	Id   string
}

const JWT_COOKIE_KEY = "uogcal_token"

func GetAuthClaims(c *fiber.Ctx) (ClaimsMap, error) {
	token := c.Cookies(JWT_COOKIE_KEY)
	parsed, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return ClaimsMap{}, err
	}
	claims := parsed.Claims.(jwt.MapClaims)

	res := ClaimsMap{
		User: claims["user"].(string),
		Id:   claims["id"].(string),
	}
	return res, nil
}

func createJWTCookie(claims ClaimsMap) (*fiber.Cookie, error) {
	expiryDate := time.Now().Add(time.Hour * 72)

	claimsMap := jwt.MapClaims{
		"user": claims.User,
		"id":   claims.Id,
		"exp":  expiryDate.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsMap)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	cookie := fiber.Cookie{
		Name:     JWT_COOKIE_KEY,
		Value:    t,
		Expires:  expiryDate,
		Path:     "/",
		SameSite: fiber.CookieSameSiteStrictMode,
		Secure:   true,
	}

	return &cookie, nil
}

func AuthRoutes(r fiber.Router) {
	r.Use([]string{"/auth", "/me", "/upload"}, jwtware.New(jwtware.Config{
		TokenLookup: "cookie:" + JWT_COOKIE_KEY,
		Filter: func(c *fiber.Ctx) bool {
			url := c.OriginalURL()
			switch url {
			case "/auth/login":
				fallthrough
			case "/auth/login/":
				return true
			case "/auth/register":
				fallthrough
			case "/auth/register/":
				return true
			}
			return false
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			c.Set(fiber.HeaderCacheControl, "no-store, must-revalidate")
			return c.Redirect("/login")
		},
		SigningKey: jwtware.SigningKey{
			Key: []byte(os.Getenv("JWT_SECRET")),
		},
	}))

	r.Post("/auth/login", func(c *fiber.Ctx) error {
		c.Accepts(fiber.MIMEApplicationJSON, fiber.MIMEApplicationForm)

		user := new(UserLogin)
		c.BodyParser(user)

		uid, password, err := database.GetUserPassword(user.Username)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		if bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password)) != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		jwtCookie, err := createJWTCookie(ClaimsMap{
			User: user.Username,
			Id:   uid,
		})
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		c.Cookie(jwtCookie)

		c.Set(fiber.HeaderCacheControl, "no-store, must-revalidate")
		return c.Redirect("/")
	})

	r.Get("/auth/logout", func(c *fiber.Ctx) error {
		cookie := fiber.Cookie{
			Name:    JWT_COOKIE_KEY,
			Expires: time.UnixMilli(0),
			Path:    "/",
		}
		c.Cookie(&cookie)

		c.Set(fiber.HeaderCacheControl, "no-store, must-revalidate")
		return c.Redirect("/")
	})

	r.Post("/auth/register", func(c *fiber.Ctx) error {
		c.Accepts(fiber.MIMEApplicationJSON, fiber.MIMEApplicationForm)

		user := new(UserLogin)
		c.BodyParser(user)

		if len(user.Username) <= 2 || len(user.Username) > 16 {
			return c.Status(400).SendString("Username must be more than 2 characters and less than 16")
		}
		if len(user.Password) < 8 || len(user.Password) > 64 {
			return c.Status(400).SendString("Password was less than 8 characters and less than 64")
		}

		uuid, err := database.AddUser(user.Username, user.Password)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		jwtCookie, err := createJWTCookie(ClaimsMap{
			User: user.Username,
			Id:   uuid.String(),
		})
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		c.Cookie(jwtCookie)

		c.Set(fiber.HeaderCacheControl, "no-store, must-revalidate")
		return c.Redirect("/")
	})
}
