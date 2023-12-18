package web

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/RonkZeDonk/uogcal/pkg/database"
	"github.com/RonkZeDonk/uogcal/pkg/redis"
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

type GoogleAccessToken struct {
	AccessToken string `json:"accessToken"`
}
type GoogleUserInfo struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verifiedEmail"`
	Name          string `json:"name"`
	GivenName     string `json:"givenName"`
	FamilyName    string `json:"familyName"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

type ClaimsMap struct {
	User string
	Id   string
}

const JWT_COOKIE_KEY = "uogcal_token"

func getClaims(c *fiber.Ctx) (jwt.MapClaims, error) {
	token := c.Cookies(JWT_COOKIE_KEY)
	parsed, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return jwt.MapClaims{}, err
	}
	return parsed.Claims.(jwt.MapClaims), nil
}

func GetAuthClaims(c *fiber.Ctx) (ClaimsMap, error) {
	claims, err := getClaims(c)
	if err != nil {
		return ClaimsMap{}, err
	}

	res := ClaimsMap{
		User: claims["user"].(string),
		Id:   claims["id"].(string),
	}
	return res, nil
}

func GetJWTExpiration(c *fiber.Ctx) (time.Duration, error) {
	claims, err := getClaims(c)
	if err != nil {
		return 0, err
	}

	exp, err := claims.GetExpirationTime()
	if err != nil {
		return 0, err
	}

	return time.Until(exp.Time), nil
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

func removeJWTCookie(c *fiber.Ctx) {
	cookie := fiber.Cookie{
		Name:    JWT_COOKIE_KEY,
		Expires: time.UnixMilli(0),
		Path:    "/",
	}
	c.Cookie(&cookie)
}

func getGoogleIdFromAT(accessToken string) (GoogleUserInfo, error) {
	req, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v1/userinfo?alt=json", strings.NewReader(""))
	if err != nil {
		return GoogleUserInfo{}, err
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return GoogleUserInfo{}, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return GoogleUserInfo{}, err
	}

	var gRes GoogleUserInfo
	if err = json.Unmarshal(body, &gRes); err != nil {
		return GoogleUserInfo{}, err
	}
	return gRes, nil
}

func deauthJWT(c *fiber.Ctx, expiry time.Duration) error {
	token := c.Cookies(JWT_COOKIE_KEY)
	return redis.Set(token, 1, expiry)
}

func AuthRoutes(r fiber.Router) {
	r.Use([]string{"/auth", "/me", "/upload"}, jwtware.New(jwtware.Config{
		TokenLookup: "cookie:" + JWT_COOKIE_KEY,
		Filter: func(c *fiber.Ctx) bool {
			url := c.OriginalURL()
			switch url {
			case "/auth/login/":
				return true
			case "/auth/register/":
				return true
			case "/auth/google/login":
				return true
			case "/auth/google/register":
				return true
			}
			return false
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			c.Set(fiber.HeaderCacheControl, "no-store, must-revalidate")
			return c.Redirect("/login")
		},
		SuccessHandler: func(c *fiber.Ctx) error {
			if redis.Exists(c.Cookies(JWT_COOKIE_KEY)) {
				removeJWTCookie(c)

				c.Set(fiber.HeaderCacheControl, "no-store, must-revalidate")
				return c.Redirect("/login")
			}
			return c.Next()
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

	r.Post("/auth/google/login", func(c *fiber.Ctx) error {
		var token GoogleAccessToken
		if err := json.Unmarshal(c.Request().Body(), &token); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		gRes, err := getGoogleIdFromAT(token.AccessToken)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		username, uuid, err := database.GetOAuthUser(gRes.Id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		jwtCookie, err := createJWTCookie(ClaimsMap{
			User: username,
			Id:   uuid,
		})
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		c.Cookie(jwtCookie)

		c.Set(fiber.HeaderCacheControl, "no-store, must-revalidate")
		return c.Redirect("/")
	})

	r.Post("/auth/google/register", func(c *fiber.Ctx) error {
		var token GoogleAccessToken
		if err := json.Unmarshal(c.Request().Body(), &token); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		gRes, err := getGoogleIdFromAT(token.AccessToken)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		username, uuid, err := database.AddOAuthUser(gRes.Id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		jwtCookie, err := createJWTCookie(ClaimsMap{
			User: username,
			Id:   uuid.String(),
		})
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		c.Cookie(jwtCookie)

		c.Set(fiber.HeaderCacheControl, "no-store, must-revalidate")
		return c.Redirect("/")
	})

	r.Get("/auth/logout", func(c *fiber.Ctx) error {
		timeUntilExpire, err := GetJWTExpiration(c)
		if err != nil {
			return c.SendString(err.Error())
		}

		deauthJWT(c, timeUntilExpire)
		removeJWTCookie(c)

		c.Set(fiber.HeaderCacheControl, "no-store, must-revalidate")
		return c.Redirect("/")
	})
}
