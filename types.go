package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Status struct {
	Status     string `json:"status" bson:"status"`
	RetryCount int    `json:"retryCount" bson:"retry_count"`
}

type Info struct {
	Id    int    `json:"id" bson:"user_id"`
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
}
type Event struct {
	Id         primitive.ObjectID `json:"id" bson:"_id"`
	EventType  string             `json:"eventType" bson:"event_type"`
	Status     string             `json:"status" bson:"status"`
	Payload    Info               `json:"payload" bson:"payload"`
	RetryCount int                `json:"retryCount" bson:"retry_count"`
	CreatedAt  time.Time          `json:"createdAt" bson:"created_at"`
	UpdatedAt  time.Time          `json:"updatedAt" bson:"updated_at"`
}
type DeadEvent struct{
	Id         primitive.ObjectID `json:"id" bson:"_id"`
	EventType  string             `json:"eventType" bson:"event_type"`
	Payload    Info               `json:"payload" bson:"payload"`
	CreatedAt  time.Time          `json:"createdAt" bson:"created_at"`
	FailedAt  time.Time          `json:"failedAt" bson:"failed_at"`
	FailureReason string		`json:"failureReason" bson:"failure_reason"`


}
func NewDeadEvent(event *Event,failureReason string,failureTime time.Time)*DeadEvent{
	return &DeadEvent{
		Id: event.Id,
		FailedAt: failureTime,
		FailureReason: failureReason,
		Payload: event.Payload,
		CreatedAt: event.CreatedAt,
		EventType: event.EventType,

	}
}

func NewRegisterEvent(information Info) *Event {
	return &Event{
		Id:         primitive.NewObjectID(),
		EventType:  "user_registration",
		Payload:    information,
		RetryCount: 0,
		Status:     "pending",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}
