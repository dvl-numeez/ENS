package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type ResponseMessage struct {
	Message string `json:"message"`
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type ApiError struct {
	Error string
}
type ProducerService struct {
	address string
	store   Store
}

func GetProducerService(address string, store Store) *ProducerService {
	return &ProducerService{
		address: address,
		store:   store,
	}
}

func (service *ProducerService) Run() {
	router := http.NewServeMux()
	router.HandleFunc("/register", makeHttpHandler(service.HandleRegister))
	fmt.Println("Server starting on port", service.address)
	if err := http.ListenAndServe(service.address, router); err != nil {
		fmt.Println("Error occurred:", err)
	}
}

func (service *ProducerService)HandleRegister(w http.ResponseWriter, r *http.Request) error {
	if r.Method!="POST"{
		return errors.New("invalid method")
	}
	info := new(Info)
	err := json.NewDecoder(r.Body).Decode(info)
	if err != nil {
		return err
	}
	registerEvent := NewRegisterEvent(*info)
	err = service.store.Insert(r.Context(), registerEvent)
	if err != nil {
		return err
	}
	WriteJson(w, http.StatusCreated, "Event created")
	return nil

}

func makeHttpHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			WriteError(w, err)
		}
	}
}

func WriteError(w http.ResponseWriter, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(ApiError{Error: err.Error()})
}

func WriteJson(w http.ResponseWriter, status int, message string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(ResponseMessage{Message: message})
}
