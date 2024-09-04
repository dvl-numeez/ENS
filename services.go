package main

import (
	"context"
	"log"
	"sync"
)



func Start(){
	var wg sync.WaitGroup
	store,err:=GetStore(context.Background())
	if err!=nil{
		log.Fatal("Error occured : ",err)
	}
	service:=GetProducerService(":3000",store)
	wg.Add(2)
	go func(){
		service.Run()
	}()

	go func(){
		service:=GetConsumerService(":3001",store)
		service.Run()
	}()

	wg.Wait()


}