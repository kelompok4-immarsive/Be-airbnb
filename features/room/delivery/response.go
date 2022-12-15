package delivery

import "fajar/testing/features/room"

type RoomRespon struct {
	Name_room string  `json:"name_room"`
	Price     float64 `json:"price"`
	Deskripsi string  `json:"deskripsi"`
	Status    string  `json:"status"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	UserID    uint    `json:"id_user"`
}

func RoomCoreToRoomRespon(dataCore room.RoomCore) RoomRespon { // data user core yang ada di controller yang memanggil user repository
	return RoomRespon{

		Name_room: dataCore.Name_room,
		Price:     dataCore.Price, //mapping data core ke data gorm model
		Deskripsi: dataCore.Deskripsi,
		Status:    dataCore.Status,
		Longitude: dataCore.Longitude,
		Latitude:  dataCore.Latitude,
	}
}
func ListRoomCoreToRoomRespon(dataCore []room.RoomCore) []RoomRespon { //data user.core data yang diambil dari entities ke respon struct
	var ResponData []RoomRespon

	for _, value := range dataCore { //memanggil paramete data core yang berisi data user core
		ResponData = append(ResponData, RoomCoreToRoomRespon(value)) // mengambil data mapping dari user core to respon
	}
	return ResponData
}
