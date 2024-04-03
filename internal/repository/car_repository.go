package repository

import (
	"CatCatalog/internal/model"
	"database/sql"
)

type CarRepository interface {
	GetAllCars() ([]model.Car, error)
	GetCarByID(id string) (model.Car, error)
	DeleteCarByID(id string) error
	UpdateCarByID(id string, updatedCar model.Car) error
	CreateCar(car model.Car) error
}

type carRepository struct {
	db *sql.DB
}

func NewCarRepository(db *sql.DB) CarRepository {
	return &carRepository{db: db}
}

func (c carRepository) GetAllCars() ([]model.Car, error) {
	query := "SELECT id, reg_num, mark, model, year, owner_name, owner_surname, owner_patronymic from cars"

	rows, err := c.db.Query(query)
	if err != nil {
		return []model.Car{}, err
	}

	var cars []model.Car

	for rows.Next() {
		var car model.Car

		err = rows.Scan(&car.Id, &car.RegNum, &car.Mark, &car.Model, &car.Year, &car.Owner.Name, &car.Owner.Surname, &car.Owner.Patronymic)
		if err != nil {
			return []model.Car{}, err
		}
		cars = append(cars, car)
	}

	return cars, nil
}

func (c carRepository) GetCarByID(id string) (model.Car, error) {
	query := "SELECT id, reg_num, mark, model, year, owner_name, owner_surname, owner_patronymic FROM cars WHERE id = $1"

	row := c.db.QueryRow(query, id)

	var car model.Car

	err := row.Scan(&car.Id, &car.RegNum, &car.Mark, &car.Model, &car.Year, &car.Owner.Name, &car.Owner.Surname, &car.Owner.Patronymic)
	if err != nil {
		return model.Car{}, err
	}

	return car, nil
}

func (c carRepository) DeleteCarByID(id string) error {
	//TODO implement me
	panic("implement me")
}

func (c carRepository) UpdateCarByID(id string, updatedCar model.Car) error {
	//TODO implement me
	panic("implement me")
}

func (c carRepository) CreateCar(car model.Car) error {
	_, err := c.db.Exec("INSERT INTO cars (reg_num, mark, model, year, owner_name, owner_surname, owner_patronymic) VALUES ($1, $2, $3, $4, $5, $6, $7)", car.RegNum, car.Mark, car.Model, car.Year, car.Owner.Name, car.Owner.Surname, car.Owner.Patronymic)
	return err
}
