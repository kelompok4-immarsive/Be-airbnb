package delivery

import (
	"fajar/testing/features/gambar"
)

type GambarRequest struct {
	Image_url string `json:"Image_url" form:"Image_url"`
	RoomID    string `json:"RoomID " form:"RoomID "`
}

func (data *GambarRequest) reqToCore() gambar.GambarCore {
	return gambar.GambarCore{
		Image_url: data.Image_url,
		RoomID:    data.RoomID,
	}
}
