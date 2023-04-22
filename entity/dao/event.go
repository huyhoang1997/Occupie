package dao

import (
	"occupie/entity"
	"time"
)

const (
	ACCEPTED = "ACCEPTED"
	CANCELED = "CANCELED"
	DECLINED = "DECLINED"
)

type Event struct {
	entity.BaseModel
	Booker         User
	BookingSubject BookingSubject
	BookingDate    time.Time
}

type EventLog struct {
	entity.BaseModel
	EventID int64 //define foreign key to Event

}
