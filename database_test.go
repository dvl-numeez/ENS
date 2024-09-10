package main

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)



func TestInsert(t *testing.T) {
	store,err:=GetStore(context.Background())
	if err!=nil{
		t.Error(err)
	}
	info:=Info{
		Id: 3556,
		Name: "Numeez",
		Email: "numeez@yahoo.com",
	}
	t.Run("Testing inserting the event in db",func(t *testing.T){
		event:=NewRegisterEvent(info)
		if err:=store.Insert(context.TODO(),event);err!=nil{
			t.Error(err)
		}
	})

}

func TestGetCount(t *testing.T){
	store,err:=GetStore(context.Background())
	if err!=nil{
		t.Error(err)
	}
	stringId:="66dfc3c972f38350735ae5e3"
	id,err:=primitive.ObjectIDFromHex(stringId)
	if err!=nil{
		t.Error(err)
	}
	event:=&Event{
		Id: id,
	}
	count,err:=store.GetCount(context.TODO(),event)
	if err !=nil{
		t.Error(err)
	}
	if count!=0{
		t.Errorf("Expected : %d Got : %d",count,0)
	}

}

func TestGetStatus(t *testing.T){
	store,err:=GetStore(context.Background())
	if err!=nil{
		t.Error(err)
	}
	stringId:="66dfc3c972f38350735ae5e3"
	id,err:=primitive.ObjectIDFromHex(stringId)
	if err!=nil{
		t.Error(err)
	}
	event:=&Event{
		Id: id,
	}
	status,err:=store.GetStatus(context.TODO(),event)
	if err!=nil{
		t.Error(err)
	}
	if (status.Status!="pending" && status.Status!="processed"){
		t.Errorf("Expected : %s Got : %s","prending/processed",status.Status)
	}

}

func TestUpdateStatus(t *testing.T) {
	info:=Info{
		Id: 3556,
		Name: "Numeez",
		Email: "numeez@yahoo.com",
	}
	store,err:=GetStore(context.Background())
	if err!=nil{
		t.Error(err)
	}
	event:=NewRegisterEvent(info)
	err=store.Insert(context.TODO(),event)
	if err!=nil{
		t.Error(err)
	}
	err=store.UpdateStatus(context.TODO(),event)
	if err!=nil{
		t.Error(err)
	}
	status,err:=store.GetStatus(context.TODO(),event)
	if err!=nil{
		t.Error(err)
	}
	if status.Status!="processed"{
		t.Errorf("Expected %s Got : %s","processed",status.Status)
	}
	
}

func TestIncrementCount(t *testing.T) {
	info:=Info{
		Id: 3556,
		Name: "Numeez",
		Email: "numeez@yahoo.com",
	}
	store,err:=GetStore(context.Background())
	if err!=nil{
		t.Error(err)
	}
	event:=NewRegisterEvent(info)
	err=store.Insert(context.TODO(),event)
	if err!=nil{
		t.Error(err)
	}
	err=store.IncrementCount(context.TODO(),event)
	if err!=nil{
		t.Error(err)
	}
	count,err:=store.GetCount(context.TODO(),event)
	if err!=nil{
		t.Error(err)
	}
	if count!=1{
		t.Errorf("Expected : %d Got : %d",1,count)
	}
}