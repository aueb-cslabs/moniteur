package authentication

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/aueb-cslabs/moniteur/backend/test"
	"github.com/aueb-cslabs/moniteur/backend/types"
)

func TestRegister(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()

	ctx := test.TestFactory()
	ctx.Context = e.NewContext(req, rec)

	ctx.SetParamNames("id")
	ctx.SetParamValues("p3180260")

	user := &types.User{}

	if assert.NoError(t, Register(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		ctx.DB.Find(user, "username = ?", "p3180260")
		assert.Equal(t, "p3180260", user.Username)
	}

	rec = httptest.NewRecorder()
	ctx.Context = e.NewContext(req, rec)

	ctx.SetParamNames("id")
	ctx.SetParamValues("p3180260")

	// Second request should return HTTP 202 code
	if assert.NoError(t, Register(ctx)) {
		assert.Equal(t, http.StatusAccepted, rec.Code)
	}

	// Empty parameter. It should return HTTP 400 code
	rec = httptest.NewRecorder()
	ctx.Context = e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("")

	if assert.NoError(t, Register(ctx)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestUnregister(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()

	ctx := test.TestFactory()
	ctx.Context = e.NewContext(req, rec)

	ctx.SetParamNames("id")
	ctx.SetParamValues("p3180260")

	if assert.NoError(t, Unregister(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, true, ctx.DB.Where("username = ?", "p3180260").Find(&types.User{}).RecordNotFound())
	}

	rec = httptest.NewRecorder()
	ctx.Context = e.NewContext(req, rec)

	ctx.SetParamNames("id")
	ctx.SetParamValues("")

	if assert.NoError(t, Unregister(ctx)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}
