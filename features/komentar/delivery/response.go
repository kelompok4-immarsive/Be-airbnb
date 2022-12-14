package delivery

import "fajar/testing/features/komentar"

type KomentarResponse struct {
	ID           uint
	Vote_star    uint
	Isi_komentar string
	User         UserResponse
	Room         RoomResponse
}

type UserResponse struct {
	ID uint
}

type RoomResponse struct {
	ID uint
}

func fromCore(dataCore komentar.CoreKomentar) KomentarResponse {
	return KomentarResponse{
		ID:           dataCore.ID,
		Vote_star:    dataCore.Vote_star,
		Isi_komentar: dataCore.Isi_komentar,
		User:         UserResponse{ID: dataCore.User.ID},
		Room:         RoomResponse{ID: dataCore.Room.ID},
	}
}

func fromCoreList(dataCore []komentar.CoreKomentar) []KomentarResponse {
	var dataResponse []KomentarResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
