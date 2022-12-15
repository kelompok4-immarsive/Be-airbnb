package delivery

import (
	"fajar/testing/features/gambar"
)

type GambarRequest struct {
	Image_url string `json:"Image_url" form:"Image_url"`
	RoomID    uint   `json:"room_id " form:"room_id "`
}

func (data *GambarRequest) reqToCore() gambar.GambarCore {
	return gambar.GambarCore{
		Image_url: data.Image_url,
		RoomID:    data.RoomID,
	}
}
