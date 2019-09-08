package rest

import (
	"errors"
	"github.com/aueb-cslabs/moniteur/types"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"strings"
	"time"
)

func Authenticate(e echo.Context) error {
	ctx := e.(*types.Context)
	user := &types.User{}

	_ = e.Bind(user)
	res, ldapErr := ctx.Plugin().AuthorizeUser(user.Username, user.Password)

	if res {
		expiresAt := time.Now().Add(time.Hour * 24).Unix()
		token := jwt.New(jwt.SigningMethodHS256)

		standardClaim := &jwt.StandardClaims{ExpiresAt: expiresAt}

		authTokenClaim := &types.AuthTokenClaim{}
		authTokenClaim.StandardClaims = standardClaim
		authTokenClaim.Username = user.Username

		token.Claims = authTokenClaim

		tokenString, err := token.SignedString([]byte(secret))
		if err != nil {
			return e.JSON(http.StatusUnauthorized, err)
		}

		authToken := &types.AuthToken{}
		authToken.Token = tokenString
		authToken.TokenType = "Bearer"
		authToken.ExpiresIn = expiresAt

		authorized[authToken.Token] = authTokenClaim

		e.Response().Header().Set("authorization", "Bearer "+authToken.Token)

		return e.NoContent(http.StatusOK)
	} else {
		return e.JSON(http.StatusUnauthorized, ldapErr)
	}
}

func Validate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("authorization")
		if authHeader != "" {
			bearerToken := strings.Split(authHeader, " ")
			if len(bearerToken) == 2 {
				token, err := jwt.Parse(bearerToken[1], jwtKey)
				if err != nil {
					return c.JSON(http.StatusUnauthorized, err)
				}
				if token.Valid {
					claim := authorized[bearerToken[1]]

					name := c.Request().Header.Get("username")

					if claim != nil {
						expiresAt := claim.StandardClaims.ExpiresAt
						username := claim.Username

						nowUnix := time.Now().Unix()

						if nowUnix < expiresAt {
							if name == username {
								return next(c)
							} else {
								return c.NoContent(http.StatusUnauthorized)
							}
						} else {
							delete(authorized, bearerToken[1])
							return c.NoContent(http.StatusUnauthorized)
						}
					}
				}
			} else {
				return c.NoContent(http.StatusUnauthorized)
			}
		} else {
			return c.NoContent(http.StatusUnauthorized)
		}
		return c.NoContent(http.StatusUnauthorized)
	}
}

func jwtKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("unexpected error")
	}
	return []byte(secret), nil
}
