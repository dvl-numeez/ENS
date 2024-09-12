package main

import (
	"context"
	"errors"
	"reflect"

	// "strings"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestDeadEventQueue(t *testing.T) {
	id := primitive.NewObjectID()
	store, err := GetStore(context.Background())
	if err != nil {
		t.Error(err)
	}
	service := GetConsumerService(":3000", store)
	deadEvent := DeadEvent{
		Id:            id,
		FailedAt:      time.Now(),
		CreatedAt:     time.Now(),
		FailureReason: "Email service failed",
	}
	err = service.AddToDeadEventQueue(context.TODO(), &deadEvent)
	if err != nil {
		t.Error("Dead event did not created it should have been created ideally")
	}
}

func TestMakeEvent(t *testing.T) {
	id := primitive.NewObjectID().Hex()
	created_at := "2006-01-02T15:04:05Z"
	updated_at := "2006-01-02T15:04:05Z"
	retry_count := 0
	event_type := "User Registration"
	status := "pending"
	information := Info{
		Id:    1245786554,
		Name:  "Numeez",
		Email: "numeez@gmail.com",
	}
	payload := map[string]interface{}{
		"user_id": 1245786554,
		"name":    "Numeez",
		"email":   "numeez@gmail.com",
	}
	event := map[string]interface{}{
		"_id":         id,
		"created_at":  created_at,
		"updated_at":  updated_at,
		"retry_count": retry_count,
		"event_type":  event_type,
		"payload":     payload,
		"status":      status,
	}
	result, err := makeEvent(event)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(result.Payload, information) {
		t.Errorf("Expected : %v Actual : %v", result, information)
	}

	t.Run("Giving invalid dates", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		created_at := "2006-01-02T15:04:0"
		updated_at := "2006-01-02T15:04:"
		retry_count := 0
		event_type := "User Registration"
		status := "pending"

		payload := map[string]interface{}{
			"user_id": 1245786554,
			"name":    "Numeez",
			"email":   "numeez@gmail.com",
		}
		event := map[string]interface{}{
			"_id":         id,
			"created_at":  created_at,
			"updated_at":  updated_at,
			"retry_count": retry_count,
			"event_type":  event_type,
			"payload":     payload,
			"status":      status,
		}
		_, err := makeEvent(event)
		if err == nil {
			t.Error("Expected an error but did not get")
		}
	})
}

func TestOperate(t *testing.T) {
	store, err := GetStore(context.Background())
	if err != nil {
		t.Error(err)
	}
	service := GetConsumerService(":3000", store)
	t.Run("Passing different event type", func(t *testing.T) {
		event := Event{
			EventType: "abc",
		}
		err := service.Operate(&event,false)
		outputErr := errors.New("wrong event type")
		if err.Error() != outputErr.Error() {
			t.Errorf("Expected : %s Actual : %s", outputErr.Error(), err.Error())
		}
	})

	t.Run("Testing the update functionality in the function", func(t *testing.T) {
		event:=NewRegisterEvent(Info{
			Id:    287382389,
			Name:  "Numeez",
			Email: "numeez@gmail.com",
		})
		err:=service.store.Insert(context.TODO(),event)
		if err!=nil{
			t.Error(err)
		}
		err=service.Operate(event,true)
		if err!=nil{
			t.Error(err)
		}
		status,err:=service.store.GetStatus(context.TODO(),event)
		if err!=nil{
			t.Error()
		}
		if status.Status!="processed"{
			t.Errorf("Expected %s Actual %s","processed",status.Status)
		}
	})
	t.Run("Testing the actual function",func(t *testing.T){
		event:=NewRegisterEvent(Info{
			Id:    287382389,
			Name:  "Numeez",
			Email: "numeez@gmail.com",
		})
		err:=service.store.Insert(context.TODO(),event)
		if err!=nil{
			t.Error(err)
		}
		err=service.Operate(event,false)
		if err!=nil{
			t.Error(err)
		}
	})

}
