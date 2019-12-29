package authentication

import (
	"errors"
	"github.com/aueb-cslabs/moniteur/backend/rest"
	"github.com/aueb-cslabs/moniteur/backend/types"
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

	if ctx.AuthUsers.Exists(user.Username).Val() != 1 {
		return e.NoContent(http.StatusUnauthorized)
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

		tokenString, err := token.SignedString([]byte(rest.Secret))
		if err != nil {
			return e.JSON(http.StatusUnauthorized, err)
		}

		authToken := &types.AuthToken{}
		authToken.Token = tokenString
		authToken.TokenType = "Bearer"
		authToken.ExpiresIn = expiresAt

		rest.Authorized[authToken.Token] = authTokenClaim

		return e.JSON(http.StatusOK, authToken)
	} else {
		return e.JSON(http.StatusUnauthorized, ldapErr)
	}
}

func AuthenticateToken(e echo.Context) error {
	ctx := e.(*types.Context)

	authHeader := e.Request().Header.Get("Authorization")
	if authHeader == "" {
		return e.NoContent(http.StatusUnauthorized)
	}
	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) != 2 {
		return e.NoContent(http.StatusUnauthorized)
	}
	token, err := jwt.Parse(bearerToken[1], jwtKey)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, err)
	}
	if !token.Valid {
		return e.NoContent(http.StatusUnauthorized)
	}
	claim := rest.Authorized[bearerToken[1]]
	name := e.Request().Header.Get("Username")
	if ctx.AuthUsers.Exists(name).Val() != 1 {
		return e.NoContent(http.StatusUnauthorized)
	}
	if claim == nil {
		return e.NoContent(http.StatusUnauthorized)
	}
	expiresAt := claim.StandardClaims.ExpiresAt
	username := claim.Username

	nowUnix := time.Now().Unix()
	if nowUnix >= expiresAt {
		delete(rest.Authorized, bearerToken[1])
		return e.NoContent(http.StatusUnauthorized)
	}
	if name != username {
		return e.NoContent(http.StatusUnauthorized)
	} else {
		return e.NoContent(http.StatusOK)
	}
}

// Validate user validation of JWT token
func Validate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.(*types.Context)

		authHeader := c.Request().Header.Get("authorization")
		if authHeader == "" {
			return c.NoContent(http.StatusUnauthorized)
		}
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			return c.NoContent(http.StatusUnauthorized)
		}
		token, err := jwt.Parse(bearerToken[1], jwtKey)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err)
		}
		if !token.Valid {
			return c.NoContent(http.StatusUnauthorized)
		}
		claim := rest.Authorized[bearerToken[1]]
		name := c.Request().Header.Get("Username")
		if ctx.AuthUsers.Exists(name).Val() != 1 {
			return c.NoContent(http.StatusUnauthorized)
		}
		if claim == nil {
			return c.NoContent(http.StatusUnauthorized)
		}
		expiresAt := claim.StandardClaims.ExpiresAt
		username := claim.Username

		nowUnix := time.Now().Unix()
		if nowUnix >= expiresAt {
			delete(rest.Authorized, bearerToken[1])
			return c.NoContent(http.StatusUnauthorized)
		}
		if name != username {
			return c.NoContent(http.StatusUnauthorized)
		} else {
			return next(c)
		}
	}
}

func Invalidate(e echo.Context) error {
	authToken := e.Request().Header.Get("Authorization")
	user := e.Request().Header.Get("Username")

	if authToken == "" {
		return e.NoContent(http.StatusBadRequest)
	}
	bearerToken := strings.Split(authToken, " ")
	if len(bearerToken) != 2 {
		return e.NoContent(http.StatusBadRequest)
	}
	token, err := jwt.Parse(bearerToken[1], jwtKey)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}
	if !token.Valid {
		return e.NoContent(http.StatusUnauthorized)
	}
	claim := rest.Authorized[bearerToken[1]]
	if claim == nil {
		return e.NoContent(http.StatusBadRequest)
	}
	if claim.Username == user {
		delete(rest.Authorized, bearerToken[1])
		return e.NoContent(http.StatusOK)
	}
	return e.NoContent(http.StatusBadRequest)
}

// jwtKey checks if the token is signed
func jwtKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("unexpected error")
	}
	return []byte(rest.Secret), nil
}

func Users(e echo.Context) error {
	ctx := e.(*types.Context)

	authorizedUsers := ctx.AuthUsers.Do("SCAN", "0", "COUNT", "1000")
	data := authorizedUsers.Val().([]interface{})
	users := data[1].([]interface{})
	return e.JSON(http.StatusOK, users)
}
