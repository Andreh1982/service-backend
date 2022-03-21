package controllers

import (
	"fmt"
	"net/http"

	"service-backend/database"
	"service-backend/models"

	"github.com/gin-gonic/gin"
)

func (a *APIEnv) GetBuyers(c *gin.Context) {
	CorsSetup(c)
	person, err := database.GetBuyers(a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, person)
}

func (a *APIEnv) GetBuyer(c *gin.Context) {
	id := c.Params.ByName("id")
	person, exists, err := database.GetBuyerByID(id, a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Buyer not found."})
		return
	}
	c.IndentedJSON(http.StatusOK, person)
}

func (a *APIEnv) CreateBuyer(c *gin.Context) {
	person := models.Buyer{}
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

func (a *APIEnv) DeleteBuyer(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := database.GetBuyerByID(id, a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Buyer not found."})
		return
	}
	err = database.DeleteBuyer(id, a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status": "Buyer deleted."})
}

func (a *APIEnv) UpdateBuyer(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := database.GetBuyerByID(id, a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Buyer not exists."})
		return
	}
	updatedBuyer := models.Buyer{}
	err = c.BindJSON(&updatedBuyer)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := database.UpdateBuyer(a.DB, &updatedBuyer); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	a.GetBuyer(c)
}
