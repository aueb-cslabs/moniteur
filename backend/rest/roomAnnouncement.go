package rest

import (
	"github.com/aueb-cslabs/moniteur/backend/types"
	"github.com/aueb-cslabs/moniteur/backend/utils"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

//TODO REMOVE EXCESS CODE AND CLEAN IT UP

// createRoomAnn Method that accepts POSTs a room announcement
func createRoomAnn(e echo.Context) error {
	c := e.(*types.Context)

	post := &types.AnnouncementPost{}
	err := e.Bind(post)

	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	ann := &types.Announcement{}
	ann.End, err = types.ConvertDateToUnix(post.End)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}
	ann.Msg = post.Msg

	//Because of the room names that are in Greek we need to convert it to English
	room := c.Plugin().ConvertNameInverse(e.Param("room"))

	announcements[room] = ann

	redisClient.Set(room+"_ann", ann.Msg, 0)
	redisClient.ExpireAt(room+"_ann", time.Unix(ann.End, 0))
	redisClient.Set(room+"_ann_dt", ann.End, 0)

	writeRoomAnnouncements()

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

	ann := &types.Announcement{}
	ann.End, err = types.ConvertDateToUnix(post.End)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}
	ann.Msg = post.Msg

	//Because of the room names that are in Greek we need to convert it to English
	room := c.Plugin().ConvertNameInverse(e.Param("room"))

	announcements[room] = ann

	redisClient.Set(room+"_ann", ann.Msg, 0)
	redisClient.ExpireAt(room+"_ann", time.Unix(ann.End, 0))
	redisClient.Set(room+"_ann_dt", ann.End, 0)

	writeRoomAnnouncements()

	return e.NoContent(http.StatusOK)
}

// deleteRoomAnn Method that accepts DELETEs a room announcement
func deleteRoomAnn(e echo.Context) error {
	c := e.(*types.Context)
	room := c.Plugin().ConvertNameInverse(e.Param("room"))
	delete(announcements, room)

	writeRoomAnnouncements()
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

// checkRoomAnnouncementsExpiration checks every hour if any room announcement has expired
func checkRoomAnnouncementsExpiration() {
	for {
		for i, k := range announcements {
			now := time.Now().Unix()
			if now >= k.End {
				delete(announcements, i)
			}
		}
		time.Sleep(time.Hour)
	}
}

// writeRoomAnnouncements writes room announcement to existing.yml
func writeRoomAnnouncements() {
	_ = utils.WriteRoomAnnouncement(announcements)
}
