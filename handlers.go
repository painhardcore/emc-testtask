package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	tarantool "github.com/tarantool/go-tarantool"
)

func (a *application) get(c echo.Context) error {
	//get serialnumber from url
	serial := c.Param("serialnumber")
	s, err := strconv.ParseUint(serial, 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	//request to database
	car := new(car)
	err = a.db.GetTyped("car", 0, []interface{}{s}, car)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	//empty car
	if car.Serialnumber == 0 {
		return c.NoContent(http.StatusNoContent)
	}

	return c.JSON(http.StatusOK, car)
}

func (a *application) set(c echo.Context) error {
	//Getting all car fields from json
	car := new(car)
	if err := c.Bind(car); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	if car.Serialnumber == 0 {
		return c.String(http.StatusBadRequest, "SerialNumber can't be empty")
	}
	//Insertion car in database
	resp, err := a.db.Insert("car", car)
	if err != nil {
		if resp.Code == tarantool.ErrTupleFound {
			return c.String(http.StatusBadRequest, "Duplicate")
		}
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "OK")
}
