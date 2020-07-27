package announcements

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/aueb-cslabs/moniteur/backend/test"
	"github.com/aueb-cslabs/moniteur/backend/types"
)

var (
	ann = types.Announcement{
		End: time.Now().Add(time.Hour).Unix(),
		Msg: "Test Announcement",
	}
)

func TestCreateAnnouncement(t *testing.T) {
	testData, _ := json.Marshal(ann)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(testData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	ctx := test.TestFactory()
	ctx.Context = e.NewContext(req, rec)

	if assert.NoError(t, createAnnouncement(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, int64(1), ctx.RedisClient.Exists("Announcement").Val())
	}
}

func TestGetAnnouncement(t *testing.T) {
	testData, _ := json.Marshal(ann)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	ctx := test.TestFactory()
	ctx.Context = e.NewContext(req, rec)

	if assert.NoError(t, announcement(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(testData), strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}

func TestDeleteAnnouncement(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()

	ctx := test.TestFactory()
	ctx.Context = e.NewContext(req, rec)

	if assert.NoError(t, deleteAnnouncement(ctx)) {
		assert.Equal(t, int64(0), ctx.RedisClient.Exists("Announcement").Val())
	}
}
