package main

import (
	"context"
	"log"
)




func main(){
	store,err:=GetStore(context.Background())
	if err!=nil{
		log.Fatal("Error occured : ",err)
	}
	service:=GetProducerService(":3000",store)	
	service.Run()
}