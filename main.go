package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tarantool/go-tarantool"
)
//Our application
type application struct {
	db *tarantool.Connection
}

//ConnectDB make connection with database
func(app *application) ConnectDB() error {
	conn, err := tarantool.Connect("tarantool:3301", tarantool.Opts{User: "guest"})
	if err != nil {
		fmt.Println("Connection refused:", err)
		return err
	}
	app.db = conn
	return nil
}
func main() {
	e := echo.New()
	//make app with database
	app := new(application)
	err:= app.ConnectDB()
	if err != nil {
		return //We don't work without database
	}
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/car/:serialnumber", app.get)
	e.POST("/car", app.set)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))

}
