package controllers

import (
	"net/http"
	"strconv"

	"github.com/WillfJunior/api-go/models"
	"github.com/labstack/echo/v4"
)

// GetClients handles GET /clients
func GetClients(c echo.Context) error {
	return c.JSON(http.StatusOK, models.GetClients())
}

// GetClient handles GET /clients/:id
func GetClient(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	client, ok := models.GetClient(id)
	if !ok {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, client)
}

// CreateClient handles POST /clients
func CreateClient(c echo.Context) error {
	client := new(models.Client)
	if err := c.Bind(client); err != nil {
		return err
	}

	models.CreateClient(client)

	return c.JSON(http.StatusCreated, client)
}

// UpdateClient handles PUT /clients/:id
func UpdateClient(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	client, ok := models.GetClient(id)
	if !ok {
		return c.NoContent(http.StatusNotFound)
	}

	if err := c.Bind(client); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, client)
}

// DeleteClient handles DELETE /clients/:id
func DeleteClient(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	models.DeleteClient(id)
	return c.NoContent(http.StatusNoContent)
}
