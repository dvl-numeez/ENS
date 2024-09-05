package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)





type ConsumerService struct{
	address string
	store Store
}

func GetConsumerService(address string ,store Store)*ConsumerService{
	return &ConsumerService{
		address: address,
		store: store,
	}
}

func(service *ConsumerService)Run(){
	// router:=http.NewServeMux()
	err:=service.OpenChannelStream(context.TODO())
	if err!=nil{
		log.Fatal(err)
	}
	// http.ListenAndServe(service.address,router);err!=nil{
	// 	log.Fatal(err)
	// }// fmt.Println("Consumer service is starting on port",service.address)
	// if err:=http.Liste
}


func(service *ConsumerService) OpenChannelStream(ctx context.Context)error{
	fmt.Println("Stream function called")
	store,err:=GetStore(ctx)
	if err!=nil{
		return err
	}
	coll:=store.collection.Collection("Event-Document")
	stream,err:=coll.Watch(ctx,mongo.Pipeline{})
	if err!=nil{
		return err
	}
	defer stream.Close(ctx)
	for stream.Next(ctx){
		var change bson.M 
		var document map[string]interface{}
		if err:=stream.Decode(&change);err!=nil{
			return err
		}
		data:=change["fullDocument"]
		if data!=nil{
		dataInBytes,err:=json.Marshal(data)
		if err!=nil{
			return err
		}
		err = json.Unmarshal(dataInBytes,&document)
		if err!=nil{
			return err
		}

		event,err:=makeEvent(document)
		if err!=nil{
			return err
		}
		err=service.Operate(event)
		if err!=nil{
			log.Fatal(err)
		}
	}
}
	if err:=stream.Err();err!=nil{
		return err
	}
	return nil

}


func(service *ConsumerService) Operate(event *Event)error{
	if event.EventType=="user_registration"{
		err:=emailMockService(event.Payload.Name,event.Payload.Email)
		if err!=nil{
		go service.Retry(context.TODO(),event)
		return nil
			}
		err=service.store.UpdateStatus(context.TODO(),event)
		if err!=nil{
			return err
		}
	}
	return nil
}


func makeEvent(data map[string]interface{})(*Event,error){
	var document map[string]interface{}
	var info map[string]interface{}
	dataInBytes,err:=json.Marshal(data)
		if err!=nil{
			return nil,err
		}
		err = json.Unmarshal(dataInBytes,&document)
		if err!=nil{
			return nil,err
		}
	id:=document["_id"].(string)
	createdAt:=document["created_at"].(string)
	updatedAt:=document["updated_at"].(string)
	status:=document["status"].(string)
	event:=document["event_type"].(string)
	retryCount:=document["retry_count"].(float64)
	infoInBytes,err:=json.Marshal(data["payload"])
		if err!=nil{
			return nil,err
		}
		err = json.Unmarshal(infoInBytes,&info)
		if err!=nil{
			return nil,err
		}
	infoId:=info["user_id"]
	name:=info["name"]
	email:=info["email"]
	dataInformation:=Info{
		Id: int(infoId.(float64)),
		Name: name.(string),
		Email: email.(string),
	}
	idString,err:=primitive.ObjectIDFromHex(id)
	if err!=nil{
		return nil,err
	}

	createdAtTime,err:=time.Parse(time.RFC3339,createdAt)
	if err!=nil{
		return nil,err
	}
	updatedAtTime,err:=time.Parse(time.RFC3339,updatedAt)
	if err!=nil{
		return nil,err
	}

	return &Event{
		Id: idString,
		CreatedAt: createdAtTime,
		UpdatedAt: updatedAtTime,
		Status: status,
		RetryCount: int(retryCount),
		EventType: event,
		Payload: dataInformation,

	},nil
}


func(service *ConsumerService) RetrySend(ctx context.Context , event *Event)error{
	status,err:=service.store.GetStatus(ctx,event)
	if err!=nil{
		return err
	}
	if status.RetryCount==5{
		deadEvent:=NewDeadEvent(event,"Email service is not available",time.Now())
		if err:=service.AddToDeadEventQueue(ctx,deadEvent);err!=nil{
			return err
		}
		return nil

	}
	if status.Status=="pending" && status.RetryCount<=5{
		err:=emailMockService(event.Payload.Name,event.Payload.Email)
		if err!=nil{
			service.store.IncrementCount(ctx,event)
			return err
		}
		err=service.store.UpdateStatus(ctx,event)
		if err!=nil{
			return err
		}
	}
	return nil
}

func(service *ConsumerService)Retry(ctx context.Context , event *Event){
	ticker:=time.NewTicker(1*time.Minute)
	done:=make(chan bool)
	go func(){
		for {
			select{
			case<-done:
				return
			case<-ticker.C:
				service.RetrySend(ctx,event)
					
			}
		}
	}()
	time.Sleep(6*time.Minute)
	ticker.Stop()
	done<-true

}


func(service *ConsumerService)AddToDeadEventQueue(ctx context.Context,event *DeadEvent)error{
	store,err:=GetStore(ctx)
	if err!=nil{
		return err
	}
	coll:=store.collection.Collection("Dead-Event-Queue")
	if _,err:=coll.InsertOne(ctx,event);err!=nil{
		return err
	}
	return nil
}