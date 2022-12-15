package delivery

import "fajar/testing/features/komentar"

type KomentarResponse struct {
	ID           uint
	Vote_star    uint
	Isi_komentar string
	UserID       uint
	RoomID       uint
}

func fromCore(dataCore komentar.CoreKomentar) KomentarResponse {
	return KomentarResponse{
		ID:           dataCore.ID,
		Vote_star:    dataCore.Vote_star,
		Isi_komentar: dataCore.Isi_komentar,
		UserID:       dataCore.UserID,
		RoomID:       dataCore.RoomID,
	}
}

func fromCoreList(dataCore []komentar.CoreKomentar) []KomentarResponse {
	var dataResponse []KomentarResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
