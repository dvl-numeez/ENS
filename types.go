package main

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Info struct {
	Id    int    `json:"id" bson:"user_id"`
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
}
type Event struct {
	Id         primitive.ObjectID    `json:"id" bson:"_id"`
	EventType  string    `json:"eventType" bson:"event_type"`
	Status     string    `json:"status" bson:"status"`
	Payload    Info      `json:"payload" bson:"payload"`
	RetryCount int       `json:"retryCount" bson:"retry_count"`
	CreatedAt  time.Time `json:"createdAt" bson:"created_at"`
	UpdatedAt  time.Time `json:"updatedAt" bson:"updated_at"`
}

func NewRegisterEvent(information Info)*Event{
	return &Event{
		Id: primitive.NewObjectID(),
		EventType: "user_registration",
		Payload: information,
		RetryCount: 0,
		Status: "pending",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),

	}
}
