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
	producerService:=GetProducerService(":3000",store)
	consumerservice:=GetConsumerService(":3001",store)
	wg.Add(2)
	go func(){
		producerService.Run()
	}()

	go func(){
		consumerservice.Run()
	}()

	wg.Wait()


}