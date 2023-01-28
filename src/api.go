package main

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func Register(c echo.Context) error {
	m := &MessageResponse{}
	u := &User{}

	err := c.Bind(u)
	if err != nil {
		log.Printf("la estructura recibida no es válida: %v", err)
		m.Type = "error"
		m.Message = "la estructura no es válida"
		return c.JSON(http.StatusBadRequest, m)
	}

	addUser(u)
	m.Type = "ok"
	m.Message = "registrado correctamente"
	return c.JSON(http.StatusCreated, m)
}

func Login(c echo.Context) error {
	m := &MessageResponse{}
	u := &User{}

	err := c.Bind(u)
	if err != nil {
		log.Printf("la estructura recibida no es válida: %v", err)
		m.Type = "error"
		m.Message = "la estructura no es válida"
		return c.JSON(http.StatusBadRequest, m)
	}

	if !login(u) {
		m.Type = "error"
		m.Message = "usuario o contraseña no válidos"
		return c.JSON(http.StatusUnauthorized, m)
	}

	m.Type = "ok"
	m.Message = "bienvenido"
	m.Data = token
	return c.JSON(http.StatusOK, m)
}
