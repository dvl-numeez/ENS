package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteError(t *testing.T) {
	cases := []string{"This is the new error", "Server down", "Invalid header", "Maximum capacity exceeded"}
	for _, errorMsg := range cases {
		t.Run("Testing the write error message", func(t *testing.T) {
			result := ApiError{}
			output := httptest.NewRecorder()
			WriteError(output, errors.New(errorMsg))
			if err := json.NewDecoder(output.Body).Decode(&result); err != nil {
				t.Error(err)
			}
			if result.Error != errorMsg && output.Code != http.StatusInternalServerError {
				t.Errorf("Expected message : %s Actual message : %s Actual Code : %d Expected Code : %d", errorMsg, result.Error, output.Code, 500)
			}
		})

	}

}

func TestWriteJson(t *testing.T) {
	cases:=[]struct{
		message string
		code int
	}{
		{
			"Hello",
			200,
		},
		{
			"hi",
			300,
		},
		{
			"Bye",
			404,
		},
		{
			"Oh hi Good news",
			202,
		},
	}
	for _,c:=range cases{
		type Message struct {
			Message string
		}
		t.Run("Testing the WriteJson",func(t *testing.T){
			var message Message
			output:=httptest.NewRecorder()
			WriteJson(output,c.code,c.message)
			if err:=json.NewDecoder(output.Body).Decode(&message);err!=nil{
				t.Error(err)
			}
			if message.Message!=c.message && output.Code!=c.code{
				t.Errorf("Expected message : %s Actual message : %s Actual Code : %d Expected Code : %d", c.message,message.Message, output.Code, c.code)
			}

		})
	}
}

func TestHandleRegister(t *testing.T) {
	store,err:=GetStore(context.Background())
	if err!=nil{
		t.Error(err)
	}
	service:=GetProducerService(":3000",store)
	t.Run("When the method is other than POST",func(t *testing.T){
		request:=httptest.NewRequest(http.MethodGet,"/",nil)
		response:=httptest.NewRecorder()
		err:=service.HandleRegister(response,request)
		if err==nil{
			t.Error("Expecting an error but did not get it")
		}
		
	})
	t.Run("Testing the function",func(t *testing.T){
		requestBody:=Info{
			Id: 2892392,
			Name: "Numeez",
			Email: "numeez@gmail.com",
		}
		body,err:=json.Marshal(requestBody)
		if err!=nil{
			t.Error(err)
		}
		request:=httptest.NewRequest(http.MethodPost,"/",bytes.NewBuffer(body))
		response:=httptest.NewRecorder()
		err=service.HandleRegister(response,request)
		if err!=nil{
			t.Error(err)
		}
		var responseMessage ResponseMessage
		json.Unmarshal(response.Body.Bytes(),&responseMessage)
		if responseMessage.Message!="Event created"{
			t.Errorf("Expected : %s Actual : %s","Event created",responseMessage.Message)
		}

	})
}