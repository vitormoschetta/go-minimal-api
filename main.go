package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var db *sql.DB

var cfg = mysql.Config{
	User:                 "go",
	Passwd:               "go",
	Net:                  "tcp",
	Addr:                 "localhost:3306",
	DBName:               "go",
	AllowNativePasswords: true,
}

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

func main() {
	connectDB()
	createRoutes()
}

func connectDB() {
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected to database")
}

func createRoutes() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
	})

	router.GET("/products", func(c *gin.Context) {
		products, err := FindAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, products)
	})

	router.GET("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		product, err := FindByID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if product.ID == "" {
			c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
			return
		}
		c.JSON(http.StatusOK, product)
	})

	router.POST("/products", func(c *gin.Context) {
		var product Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		product = NewProduct(product.Name, product.Price)
		product, err := Create(product)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, product)
	})

	router.PUT("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		product, err := FindByID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if product.ID == "" {
			c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
			return
		}
		var newProduct Product
		if err := c.ShouldBindJSON(&newProduct); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newProduct.ID = product.ID
		newProduct, err = Update(newProduct)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, newProduct)
	})

	router.DELETE("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		product, err := FindByID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if product.ID == "" {
			c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
			return
		}
		err = Delete(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
	})

	router.Run("localhost:8080")
}

func FindAll() ([]Product, error) {
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func FindByID(id string) (Product, error) {
	var product Product
	err := db.QueryRow("SELECT id, name, price FROM products WHERE id = ?", id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return product, err
	}
	return product, nil
}

func Create(product Product) (Product, error) {
	res, err := db.Exec("INSERT INTO products (id, name, price) VALUES (?, ?, ?)", product.ID, product.Name, product.Price)
	if err != nil {
		return product, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return product, err
	}
	if rowsAffected != 1 {
		return product, fmt.Errorf("expected 1 row affected, got %d", rowsAffected)
	}
	return product, nil
}

func Update(product Product) (Product, error) {
	res, err := db.Exec("UPDATE products SET name = ?, price = ? WHERE id = ?", product.Name, product.Price, product.ID)
	if err != nil {
		return product, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return product, err
	}
	if rowsAffected != 1 {
		return product, fmt.Errorf("expected 1 row affected, got %d", rowsAffected)
	}
	return product, nil
}

func Delete(id string) error {
	res, err := db.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		return fmt.Errorf("expected 1 row affected, got %d", rowsAffected)
	}
	return nil
}
