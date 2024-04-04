package repository

import (
	"CatCatalog/internal/model"
	"database/sql"
	"fmt"
)

type CarRepository interface {
	GetAllCars(offset, limit int) ([]model.Car, error)
	GetCarByID(id string) (model.Car, error)
	DeleteCarByID(id string) error
	CheckCarByID(id string) (bool, error)
	UpdateCarByID(id string, updatedCar model.Car) error
	CreateCar(car model.Car) error
}

type carRepository struct {
	db *sql.DB
}

func NewCarRepository(db *sql.DB) CarRepository {
	return &carRepository{db: db}
}

func (c carRepository) GetAllCars(offset, limit int) ([]model.Car, error) {
	query := "SELECT id, reg_num, mark, model, year, owner_name, owner_surname, owner_patronymic FROM cars"

	// Добавляем операторы OFFSET и LIMIT к запросу
	query += fmt.Sprintf(" OFFSET %d LIMIT %d", offset, limit)

	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}

	var cars []model.Car

	for rows.Next() {
		var car model.Car

		err = rows.Scan(&car.Id, &car.RegNum, &car.Mark, &car.Model, &car.Year, &car.Owner.Name, &car.Owner.Surname, &car.Owner.Patronymic)
		if err != nil {
			return nil, err
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
	var exists bool
	err := c.db.QueryRow("SELECT EXISTS (SELECT 1 FROM cars WHERE id = $1)", id).Scan(&exists)
	if err != nil || !exists {
		return err
	}

	_, err = c.db.Exec("DELETE FROM cars WHERE id = $1", id)
	return err
}

func (c carRepository) CheckCarByID(id string) (bool, error) {
	var exists bool
	err := c.db.QueryRow("SELECT EXISTS (SELECT 1 FROM cars WHERE id = $1)", id).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (c carRepository) UpdateCarByID(id string, updatedCar model.Car) error {
	_, err := c.db.Exec("UPDATE cars SET reg_num = $1, mark = $2, model = $3, year = $4, owner_name = $5, owner_surname = $6, owner_patronymic = $7 WHERE id = $8",
		updatedCar.RegNum, updatedCar.Mark, updatedCar.Model, updatedCar.Year, updatedCar.Owner.Name, updatedCar.Owner.Surname, updatedCar.Owner.Patronymic, id)
	if err != nil {
		return err
	}
	return nil
}

func (c carRepository) CreateCar(car model.Car) error {
	_, err := c.db.Exec("INSERT INTO cars (reg_num, mark, model, year, owner_name, owner_surname, owner_patronymic) VALUES ($1, $2, $3, $4, $5, $6, $7)", car.RegNum, car.Mark, car.Model, car.Year, car.Owner.Name, car.Owner.Surname, car.Owner.Patronymic)
	return err
}
