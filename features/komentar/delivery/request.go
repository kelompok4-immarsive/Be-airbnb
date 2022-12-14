package delivery

import "fajar/testing/features/komentar"

type KomentarRequest struct {
	Vote_star    uint   `json:"vote_star"`
	UserID       uint   `json:"user_id"`
	RoomID       uint   `json:"room_id"`
	Isi_komentar string `json:"isi_komentar"`
}

func toCore(data KomentarRequest) komentar.CoreKomentar {
	return komentar.CoreKomentar{
		Vote_star:    data.Vote_star,
		UserID:       data.UserID,
		RoomID:       data.RoomID,
		Isi_komentar: data.Isi_komentar,
	}
}
