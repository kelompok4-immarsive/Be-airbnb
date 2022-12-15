package delivery

import (
	"fajar/testing/features/reservation"
	"time"
)

type ReservationRequest struct {
	Price       float64 `json:"price" form:"price"`
	Total_Price float64 `json:"total_price" form:"total_price"`
	Start_date  string  `json:"start_date" form:"start_date"`
	End_date    string  `json:"end_date" form:"end_date"`
	PayID       uint    `json:"pay_id " form:"pay_id"`
	Status      string  `json:"status" form:"status"`
	RoomID      uint    `json:"room_id" form:"room_id"`
	UserID      uint    `json:"id_user" form:"id_user"`
}

//	func (data *ReservationRequest) reqToCore() reservation.ReservationCore {
//		return reservation.ReservationCore{
//			Price:       data.Price,
//			Total_Price: data.Total_Price,
//			Start_date:  data.Start_date,
//			Status:      data.Status,
//			End_date:    data.End_date,
//			RoomID:      data.RoomID,
//			UserID:      data.UserID,
//		}
//	}
func RequestToCore(data ReservationRequest, Start_date, End_date time.Time) reservation.ReservationCore {
	return reservation.ReservationCore{
		RoomID:      data.RoomID,
		Total_Price: data.Total_Price,
		UserID:      data.UserID,
		Start_date:  Start_date,
		End_date:    End_date,
		Price:       data.Price,
	}
}
