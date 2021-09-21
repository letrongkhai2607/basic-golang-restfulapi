package controller

import (
	"example.com/m/v2/database"
	"example.com/m/v2/model"

	"github.com/gofiber/fiber/v2"
)

// GetAllProducts query all products
func GetAllProducts(c *fiber.Ctx) error {
	database := database.DB
	var products []model.Product
	database.Find(&products)
	return c.JSON(fiber.Map{"status": "success", "message": "All products", "data": products})
}

// GetProduct query product
func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	database := database.DB
	var product model.Product
	database.Find(&product, id)
	if product.Title == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No product found with ID", "data": nil})

	}
	return c.JSON(fiber.Map{"status": "success", "message": "Product found", "data": product})
}

// CreateProduct new product
func CreateProduct(c *fiber.Ctx) error {
	database := database.DB
	product := new(model.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create product", "data": err})
	}
	database.Create(&product)
	return c.JSON(fiber.Map{"status": "success", "message": "Created product", "data": product})
}

// DeleteProduct delete product
func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	database := database.DB

	var product model.Product
	database.First(&product, id)
	if product.Title == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No product found with ID", "data": nil})

	}
	database.Delete(&product)
	return c.JSON(fiber.Map{"status": "success", "message": "Product successfully deleted", "data": nil})
}