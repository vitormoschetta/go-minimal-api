package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func NewProduct(name string, price float64) Product {
	return Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

var products = []Product{
	NewProduct("Apple", 1.99),
	NewProduct("Orange", 2.99),
	NewProduct("Banana", 3.99),
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
	})

	router.GET("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, products)
	})

	router.GET("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, product := range products {
			if product.ID == id {
				c.JSON(http.StatusOK, product)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
	})

	router.POST("/products", func(c *gin.Context) {
		var product Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		product = NewProduct(product.Name, product.Price)
		products = append(products, product)
		c.JSON(http.StatusCreated, product)
	})

	router.PUT("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		for index, product := range products {
			if product.ID == id {
				var newProduct Product
				if err := c.ShouldBindJSON(&newProduct); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}
				newProduct.ID = product.ID
				products[index] = newProduct
				c.JSON(http.StatusOK, newProduct)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
	})

	router.DELETE("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		for index, product := range products {
			if product.ID == id {
				products = append(products[:index], products[index+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
	})

	router.Run("localhost:8080")
}
