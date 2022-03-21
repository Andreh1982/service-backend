package controllers

import (
	"fmt"
	"net/http"

	"service-backend/database"
	"service-backend/models"

	"github.com/gin-gonic/gin"
)

func (a *APIEnv) GetSellers(c *gin.Context) {
	CorsSetup(c)
	person, err := database.GetSellers(a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, person)
}

func (a *APIEnv) GetSeller(c *gin.Context) {
	id := c.Params.ByName("id")
	person, exists, err := database.GetSellerByID(id, a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Seller not found."})
		return
	}
	c.IndentedJSON(http.StatusOK, person)
}

func (a *APIEnv) CreateSeller(c *gin.Context) {
	person := models.Seller{}
	err := c.BindJSON(&person)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := a.DB.Create(&person).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, person)
}

func (a *APIEnv) DeleteSeller(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := database.GetSellerByID(id, a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Seller not found."})
		return
	}
	err = database.DeleteSeller(id, a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status": "Seller deleted."})
}

func (a *APIEnv) UpdateSeller(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := database.GetSellerByID(id, a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Seller not exists."})
		return
	}
	updatedSeller := models.Seller{}
	err = c.BindJSON(&updatedSeller)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := database.UpdateSeller(a.DB, &updatedSeller); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	a.GetSeller(c)
}
