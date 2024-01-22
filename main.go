package main

import (
	"smarty/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	//	router := mux.NewRouter()
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=Password dbname=postgres port=5432 sslmode=disable",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	//	db, err := gorm.Open(postgres.Open("postgres.db"), "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable password=Password")

	if err != nil {

		panic("failed to connect database")

	}

	var drivers = []model.Driver{

		{Name: "Jimmy Johnson", License: "ABC123"},

		{Name: "Howard Hills", License: "XYZ789"},

		{Name: "Craig Colbin", License: "DEF333"},
	}

	var cars = []model.Car{

		{Year: 2000, Make: "Toyota", ModelName: "Tundra", DriverID: 1},

		{Year: 2001, Make: "Honda", ModelName: "Accord", DriverID: 1},

		{Year: 2002, Make: "Nissan", ModelName: "Sentra", DriverID: 2},

		{Year: 2003, Make: "Ford", ModelName: "F-150", DriverID: 3},
	}
	//	db.AutoMigrate(model.Driver{})
	//	db.Select("Name", "License").Create(&drivers)
	//	db.Select("Year", "Make", "ModelName", "DriverID").Create(&cars)
	db.Create(&drivers)
	//	db.AutoMigrate(model.Car{})
	db.Create(&cars)

}
