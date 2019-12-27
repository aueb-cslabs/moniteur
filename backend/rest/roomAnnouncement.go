package rest

import (
	"github.com/aueb-cslabs/moniteur/backend/types"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

// createRoomAnn Method that accepts POSTs a room announcement
func createRoomAnn(e echo.Context) error {
	c := e.(*types.Context)

	post := &types.AnnouncementPost{}
	err := e.Bind(post)

	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	end, err := types.ConvertDateToUnix(post.End)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	//Because of the room names that are in Greek we need to convert it to English
	room := c.Plugin().ConvertNameInverse(e.Param("room"))

	redisClient.Set(room+"_ann", post.Msg, 0)
	redisClient.ExpireAt(room+"_ann", time.Unix(end, 0))
	redisClient.Set(room+"_ann_dt", end, 0)

	return e.NoContent(http.StatusOK)
}

// updateRoomAnn Method that accepts PUTs a room announcement
func updateRoomAnn(e echo.Context) error {
	c := e.(*types.Context)

	post := &types.AnnouncementPost{}
	err := e.Bind(post)

	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	end, err := types.ConvertDateToUnix(post.End)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	//Because of the room names that are in Greek we need to convert it to English
	room := c.Plugin().ConvertNameInverse(e.Param("room"))

	redisClient.Set(room+"_ann", post.Msg, 0)
	redisClient.ExpireAt(room+"_ann", time.Unix(end, 0))
	redisClient.Set(room+"_ann_dt", end, 0)

	return e.NoContent(http.StatusOK)
}

// deleteRoomAnn Method that accepts DELETEs a room announcement
func deleteRoomAnn(e echo.Context) error {
	c := e.(*types.Context)
	room := c.Plugin().ConvertNameInverse(e.Param("room"))

	redisClient.Del(room + "_ann")
	redisClient.Del(room + "_ann_dt")

	return e.NoContent(http.StatusOK)
}

// getRoomAnn Method that accepts GETs a room announcement
func getRoomAnn(e echo.Context) error {
	c := e.(*types.Context)
	room := c.Plugin().ConvertNameInverse(e.Param("room"))
	room = room + "_ann"
	exists := redisClient.Exists(room).Val()
	if exists == 1 {
		res := redisClient.Get(room).Val()
		end, _ := redisClient.Get(room + "_dt").Int64()
		roomAnn := &types.Announcement{End: end, Msg: res}
		return e.JSON(http.StatusOK, roomAnn)
	} else {
		return e.JSON(http.StatusOK, nil)
	}
}
