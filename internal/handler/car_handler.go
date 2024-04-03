package handler

import (
	"CatCatalog/internal/model"
	"CatCatalog/internal/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type CarHandler struct {
	service service.CarService
}

func NewCarHandler(service service.CarService) *CarHandler {
	return &CarHandler{service: service}
}

func (h *CarHandler) StartServer(addr string) error {
	r := mux.NewRouter()
	r.HandleFunc("/cars", h.GetAllCars)
	r.HandleFunc("/cars/get/{id}", h.GetCarByID)
	r.HandleFunc("/cars/create", h.CreateCar)
	r.HandleFunc("/cars/delete/{id}", h.DeleteCarByID)
	http.Handle("/", r)

	return http.ListenAndServe(addr, nil)
}

func (h *CarHandler) DeleteCarByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		http.Error(w, "Отсутствует идентификатор автомобиля", http.StatusBadRequest)
		return
	}

	err := h.service.DeleteCarByID(id)
	if err != nil {
		http.Error(w, "Ошибка при удалении автомобиля", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintf(w, "Автомобиль успешно удалён")
	if err != nil {
		return
	}
}

func (h *CarHandler) GetAllCars(w http.ResponseWriter, r *http.Request) {
	cars, err := h.service.GetAllCars()
	if err != nil {
		http.Error(w, "Ошибка при получении информации об автомобилях", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(cars)
	if err != nil {
		return
	}
}

func (h *CarHandler) GetCarByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if id == "" {
		http.Error(w, "Отсутствует идентификатор автомобиля", http.StatusBadRequest)
		return
	}

	car, err := h.service.GetCarByID(id)
	if err != nil {
		http.Error(w, "Ошибка при получении информации об автомобиле", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(car)
	if err != nil {
		return
	}
}

func (h *CarHandler) CreateCar(w http.ResponseWriter, r *http.Request) {
	var car model.Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, "Невозможно прочитать тело запроса", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateCar(car); err != nil {
		http.Error(w, "Ошибка при создании автомобиля", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
