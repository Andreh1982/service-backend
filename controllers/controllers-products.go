package controllers

import (
	"fmt"
	"net/http"

	"service-backend/database"
	"service-backend/models"

	"github.com/gin-gonic/gin"
)

func (a *APIEnv) GetProducts(c *gin.Context) {
	CorsSetup(c)
	person, err := database.GetProducts(a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, person)
}

func (a *APIEnv) GetProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	person, exists, err := database.GetProductByID(id, a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Product not found."})
		return
	}
	c.IndentedJSON(http.StatusOK, person)
}

func (a *APIEnv) CreateProduct(c *gin.Context) {
	person := models.Product{}
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

func (a *APIEnv) DeleteProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := database.GetProductByID(id, a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Product not found."})
		return
	}
	err = database.DeleteProduct(id, a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status": "Product deleted."})
}

func (a *APIEnv) UpdateProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := database.GetProductByID(id, a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Product not exists."})
		return
	}
	updatedProduct := models.Product{}
	err = c.BindJSON(&updatedProduct)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := database.UpdateProduct(a.DB, &updatedProduct); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	a.GetProduct(c)
}
