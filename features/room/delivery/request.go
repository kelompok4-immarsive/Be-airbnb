package delivery

import "fajar/testing/features/room"

type RoomRequest struct {
	Name_room string  `json:"name_room" form:"name_room"`
	Price     string  `json:"price" form:"price"`
	Deskripsi string  `json:"deskripsi" form:"deskripsi"`
	Status    string  `json:"status" form:"status"`
	Longitude float64 `json:"longitude" form:"longitude"`
	Latitude  float64 `json:"latitude" form:"latitude"`
	UserID    uint    `json:"id_user" form:"id_user"`
}

func (data *RoomRequest) reqToCore() room.RoomCore {
	return room.RoomCore{
		Name_room: data.Name_room,
		Price:     data.Price,
		Deskripsi: data.Deskripsi,
		Status:    data.Status,
		Longitude: data.Longitude,
		Latitude:  data.Latitude,
		UserID:    data.UserID,
	}
}
