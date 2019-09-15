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

// Authenticate user authentication based on plugin authentication and generates JWT auth token
func Authenticate(e echo.Context) error {
	ctx := e.(*types.Context)
	user := &types.User{}

	err := e.Bind(user)

	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

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

		return e.JSON(http.StatusOK, authToken)
	} else {
		return e.JSON(http.StatusUnauthorized, ldapErr)
	}
}

// Validate user validation of JWT token
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

					name := c.Request().Header.Get("Username")

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

// jwtKey checks if the token is signed
func jwtKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("unexpected error")
	}
	return []byte(secret), nil
}
