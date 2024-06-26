package service

import (
	"CatCatalog/internal/model"
	"CatCatalog/internal/repository"
	"errors"
)

type CarService interface {
	GetAllCars(offset, limit int) ([]model.Car, error)
	GetCarByID(id string) (model.Car, error)
	DeleteCarByID(id string) error
	UpdateCarByID(id string, updatedCar model.Car) error
	CreateCar(car model.Car) error
}

type carService struct {
	carRepo repository.CarRepository
}

func NewCarService(repo repository.CarRepository) CarService {
	return &carService{carRepo: repo}
}

func (c carService) GetAllCars(offset, limit int) ([]model.Car, error) {
	cars, err := c.carRepo.GetAllCars(offset, limit)
	if err != nil {
		return []model.Car{}, err
	}
	return cars, nil
}

func (c carService) GetCarByID(id string) (model.Car, error) {
	car, err := c.carRepo.GetCarByID(id)
	if err != nil {
		return model.Car{}, err
	}

	return car, nil
}

func (c carService) DeleteCarByID(id string) error {
	err := c.carRepo.DeleteCarByID(id)
	return err
}

func (c carService) UpdateCarByID(id string, updatedCar model.Car) error {
	exists, err := c.carRepo.CheckCarByID(id)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("автомобиль с указанным ID не найден")
	}

	err = c.carRepo.UpdateCarByID(id, updatedCar)
	if err != nil {
		return err
	}

	return nil
}

func (c carService) CreateCar(car model.Car) error {
	if err := c.validateCar(car); err != nil {
		return err
	}

	err := c.carRepo.CreateCar(car)
	if err != nil {
		return err
	}

	return nil
}

func (c carService) validateCar(car model.Car) error {
	if car.RegNum == "" {
		return errors.New("государственный номер автомобиля обязателен")
	}
	return nil
}
