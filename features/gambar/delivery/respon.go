package delivery

import "fajar/testing/features/gambar"

type GambarRespon struct {
	ID        uint
	Image_url string
	RoomID    uint
}

func GambarCoreToGambarRespon(dataCore gambar.GambarCore) GambarRespon { // data user core yang ada di controller yang memanggil user repository
	return GambarRespon{
		ID:        dataCore.ID,
		Image_url: dataCore.Image_url,
		RoomID:    dataCore.RoomID,
	}
}
func ListGambarCoreToGambarRespon(dataCore []gambar.GambarCore) []GambarRespon { //data user.core data yang diambil dari entities ke respon struct
	var ResponData []GambarRespon

	for _, value := range dataCore { //memanggil paramete data core yang berisi data user core
		ResponData = append(ResponData, GambarCoreToGambarRespon(value)) // mengambil data mapping dari user core to respon
	}
	return ResponData
}
