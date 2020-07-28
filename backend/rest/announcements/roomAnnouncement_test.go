package announcements

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/aueb-cslabs/moniteur/backend/test"
	"github.com/aueb-cslabs/moniteur/backend/types"
)

func TestCreateRoomAnn(t *testing.T) {
	testData, _ := json.Marshal(ann)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(testData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	ctx := test.TestFactory()
	ctx.Context = e.NewContext(req, rec)

	ctx.SetParamNames("room")
	ctx.SetParamValues("a")

	if assert.NoError(t, createRoomAnn(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, int64(1), ctx.RedisClient.Exists("a_ann").Val())
		assert.Equal(t, int64(1), ctx.RedisClient.Exists("a_ann_dt").Val())
	}
}

func TestUpdateRoomAnn(t *testing.T) {
	newAnn := types.Announcement{
		End: time.Now().Add(time.Hour).Unix(),
		Msg: "Test Announcement New",
	}
	testData, _ := json.Marshal(newAnn)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", bytes.NewReader(testData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	ctx := test.TestFactory()
	ctx.Context = e.NewContext(req, rec)

	ctx.SetParamNames("room")
	ctx.SetParamValues("a")

	if assert.NoError(t, updateRoomAnn(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, int64(1), ctx.RedisClient.Exists("a_ann").Val())
		assert.Equal(t, int64(1), ctx.RedisClient.Exists("a_ann_dt").Val())
		assert.Equal(t, "Test Announcement New", ctx.RedisClient.Get("a_ann").Val())
	}
}

func TestGetRoomAnn(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	ctx := test.TestFactory()
	ctx.Context = e.NewContext(req, rec)

	ctx.SetParamNames("room")
	ctx.SetParamValues("a")

	if assert.NoError(t, getRoomAnn(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Test Announcement New", ctx.RedisClient.Get("a_ann").Val())
	}
}

func TestDeleteRoomAnn(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()

	ctx := test.TestFactory()
	ctx.Context = e.NewContext(req, rec)

	ctx.SetParamNames("room")
	ctx.SetParamValues("a")

	if assert.NoError(t, deleteRoomAnn(ctx)) {
		assert.Equal(t, int64(0), ctx.RedisClient.Exists("a_ann").Val())
	}
}
