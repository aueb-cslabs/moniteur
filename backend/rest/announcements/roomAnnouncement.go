package announcements

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/aueb-cslabs/moniteur/backend/types"
)

// createRoomAnn Method that accepts POSTs a room announcement
func createRoomAnn(e echo.Context) error {
	ctx := e.(*types.Context)

	post := &types.Announcement{}
	err := e.Bind(post)

	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	if err := post.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	//Because of the room names that are in Greek we need to convert it to English
	room := ctx.Plugin().ConvertNameInverse(e.Param("room"))

	ctx.RedisClient.Set(room+"_ann", post.Msg, 0)
	ctx.RedisClient.ExpireAt(room+"_ann", time.Unix(post.End, 0))
	ctx.RedisClient.Set(room+"_ann_dt", post.End, 0)

	return e.NoContent(http.StatusOK)
}

// updateRoomAnn Method that accepts PUTs a room announcement
func updateRoomAnn(e echo.Context) error {
	ctx := e.(*types.Context)

	post := &types.Announcement{}
	err := e.Bind(post)

	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	if err := post.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	//Because of the room names that are in Greek we need to convert it to English
	room := ctx.Plugin().ConvertNameInverse(e.Param("room"))

	ctx.RedisClient.Set(room+"_ann", post.Msg, 0)
	ctx.RedisClient.ExpireAt(room+"_ann", time.Unix(post.End, 0))
	ctx.RedisClient.Set(room+"_ann_dt", post.End, 0)

	return e.NoContent(http.StatusOK)
}

// deleteRoomAnn Method that accepts DELETEs a room announcement
func deleteRoomAnn(e echo.Context) error {
	ctx := e.(*types.Context)

	room := ctx.Plugin().ConvertNameInverse(e.Param("room"))

	ctx.RedisClient.Del(room + "_ann")
	ctx.RedisClient.Del(room + "_ann_dt")

	return e.NoContent(http.StatusOK)
}

// getRoomAnn Method that accepts GETs a room announcement
func getRoomAnn(e echo.Context) error {
	ctx := e.(*types.Context)

	room := ctx.Plugin().ConvertNameInverse(e.Param("room"))
	room = room + "_ann"
	exists := ctx.RedisClient.Exists(room).Val()
	if exists == 1 {
		res := ctx.RedisClient.Get(room).Val()
		end, _ := ctx.RedisClient.Get(room + "_dt").Int64()
		roomAnn := &types.Announcement{End: end, Msg: res}
		return e.JSON(http.StatusOK, roomAnn)
	} else {
		return e.JSON(http.StatusOK, nil)
	}
}
