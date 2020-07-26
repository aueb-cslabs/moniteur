package authentication

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

	"github.com/aueb-cslabs/moniteur/backend/types"
)

// Authenticate user authentication based on plugin authentication and generates JWT auth token
func Authenticate(e echo.Context) error {
	ctx := e.(*types.Context)
	user := &types.User{}
	db := ctx.DB

	err := e.Bind(user)

	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	if err := user.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	if err := db.Find(&types.User{Username: user.Username}).Error; err != nil {
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

		tokenString, err := token.SignedString([]byte(ctx.Secret))
		if err != nil {
			return e.JSON(http.StatusUnauthorized, err)
		}

		authToken := &types.AuthToken{}
		authToken.Token = tokenString
		authToken.TokenType = "Bearer"
		authToken.ExpiresIn = expiresAt

		res, err := json.Marshal(authTokenClaim)

		if err != nil {
			return e.NoContent(http.StatusBadRequest)
		}

		ctx.Tokens.Set(authToken.Token, res, 0)
		ctx.Tokens.ExpireAt(authToken.Token, time.Unix(authToken.ExpiresIn, 0))

		return e.JSON(http.StatusOK, authToken)
	} else {
		return e.JSON(http.StatusUnauthorized, ldapErr)
	}
}

func AuthenticateToken(e echo.Context) error {
	ctx := e.(*types.Context)
	db := ctx.DB

	authHeader := e.Request().Header.Get("Authorization")
	if authHeader == "" {
		return e.NoContent(http.StatusUnauthorized)
	}
	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) != 2 {
		return e.NoContent(http.StatusUnauthorized)
	}
	var jwtKey = []byte(ctx.Secret)
	token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return e.JSON(http.StatusUnauthorized, err)
	}
	if !token.Valid {
		return e.NoContent(http.StatusUnauthorized)
	}
	redisClaim, err := ctx.Tokens.Get(bearerToken[1]).Bytes()
	claim := &types.AuthTokenClaim{}
	_ = json.Unmarshal(redisClaim, claim)
	name := e.Request().Header.Get("Username")
	if err := db.Find(&types.User{Username: name}).Error; err != nil {
		return e.NoContent(http.StatusUnauthorized)
	}
	if len(claim.Username) == 0 {
		return e.NoContent(http.StatusUnauthorized)
	}
	expiresAt := claim.StandardClaims.ExpiresAt
	username := claim.Username

	nowUnix := time.Now().Unix()
	if nowUnix >= expiresAt {
		ctx.Tokens.Del(bearerToken[1])
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
		db := ctx.DB

		authHeader := c.Request().Header.Get("authorization")
		if authHeader == "" {
			return c.NoContent(http.StatusUnauthorized)
		}
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			return c.NoContent(http.StatusUnauthorized)
		}
		var jwtKey = []byte(ctx.Secret)
		token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err)
		}
		if !token.Valid {
			return c.NoContent(http.StatusUnauthorized)
		}
		redisClaim, err := ctx.Tokens.Get(bearerToken[1]).Bytes()
		claim := &types.AuthTokenClaim{}
		_ = json.Unmarshal(redisClaim, claim)
		name := c.Request().Header.Get("Username")
		if err := db.Find(&types.User{Username: name}).Error; err != nil {
			return c.NoContent(http.StatusUnauthorized)
		}
		if len(claim.Username) == 0 {
			return c.NoContent(http.StatusUnauthorized)
		}
		expiresAt := claim.StandardClaims.ExpiresAt
		username := claim.Username

		nowUnix := time.Now().Unix()
		if nowUnix >= expiresAt {
			ctx.Tokens.Del(bearerToken[1])
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
	ctx := e.(*types.Context)

	authToken := e.Request().Header.Get("Authorization")
	user := e.Request().Header.Get("Username")

	if authToken == "" {
		return e.NoContent(http.StatusBadRequest)
	}
	bearerToken := strings.Split(authToken, " ")
	if len(bearerToken) != 2 {
		return e.NoContent(http.StatusBadRequest)
	}
	//token, err := jwt.Parse(bearerToken[1], NewKeyFunc(ctx.Secret))
	var jwtKey = []byte(ctx.Secret)
	token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}
	if !token.Valid {
		return e.NoContent(http.StatusUnauthorized)
	}
	redisClaim, err := ctx.Tokens.Get(bearerToken[1]).Bytes()
	claim := &types.AuthTokenClaim{}
	_ = json.Unmarshal(redisClaim, claim)
	if len(claim.Username) == 0 {
		return e.NoContent(http.StatusBadRequest)
	}
	if claim.Username == user {
		ctx.Tokens.Del(bearerToken[1])
		return e.NoContent(http.StatusOK)
	}
	return e.NoContent(http.StatusBadRequest)
}

/*// jwtKey checks if the token is signed
func jwtKey(token *jwt.Token, secret string) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("unexpected error")
	}
	return []byte(secret), nil
}*/

func Users(e echo.Context) error {
	ctx := e.(*types.Context)
	db := ctx.DB

	var users []string
	db.Model(&types.User{}).Pluck("username", &users)

	return e.JSON(http.StatusOK, users)
}
